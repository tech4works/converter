package converter

import (
	"errors"
	"fmt"
	"reflect"
)

// Lener is implemented by types that can report their length.
//
// This matches the convention used across the Go ecosystem (e.g. sort.Interface).
type Lener interface {
	Len() int
}

// Sizer is implemented by types that can report their size.
//
// Commonly used to represent "size in bytes" in some libraries/types.
type Sizer interface {
	Size() int
}

// Lengther is a custom convention for types that expose "Length".
type Lengther interface {
	Length() int
}

// ToLength returns the length or integer representation of the given value.
//
// It is a convenience wrapper around ToLengthWithErr and will panic
// if the provided value is not supported or if an error occurs.
//
// Supported behaviors:
//   - string, array, slice, map: returns len()
//   - struct: returns number of fields
//   - numeric types: returns value converted to int
//   - complex: returns real part converted to int
//   - pointer/interface: resolves underlying value recursively
//
// Use this function only when you are certain the input type is valid.
func ToLength(a any) int {
	length, err := ToLengthWithErr(a)
	if err != nil {
		panic(err)
	}
	return length
}

// ToLengthWithErr returns the length or integer representation of the given value.
//
// Behavior by type:
//
//   - reflect.String      → len(string)
//   - reflect.Array       → len(array)
//   - reflect.Slice       → len(slice)
//   - reflect.Map         → len(map)
//   - reflect.Struct      → number of struct fields
//   - reflect.Int*        → integer value
//   - reflect.Uint*       → integer value
//   - reflect.Float*      → truncated integer value
//   - reflect.Complex*    → real part truncated to int
//   - reflect.Pointer     → resolves element recursively
//   - reflect.Interface   → resolves underlying value recursively
//
// Returns an error when:
//   - the value is nil (pointer/interface)
//   - the type is not supported
//
// Note: float and complex values are truncated (not rounded).
func ToLengthWithErr(a any) (int, error) {
	if a == nil {
		return 0, errors.New("error getting length, it is null")
	}

	reflectValue := reflect.ValueOf(a)
	reflectType := reflectValue.Type()

	if reflectValue.Kind() == reflect.Ptr || reflectValue.Kind() == reflect.Interface {
		if reflectValue.IsNil() {
			return 0, errors.New("error getting length, it is null")
		} else if result, ok, err := resolveLengthImplementsIfPresent(reflectType, reflectValue); err != nil || ok {
			return result, err
		}
		return ToLengthWithErr(reflectValue.Elem().Interface())
	} else if result, ok, err := resolveLengthImplementsIfPresent(reflectType, reflectValue); err != nil || ok {
		return result, err
	}

	switch reflectValue.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		return reflectValue.Len(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(reflectValue.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return int(reflectValue.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return int(reflectValue.Float()), nil
	case reflect.Struct:
		return reflectValue.NumField(), nil
	case reflect.Complex64, reflect.Complex128:
		return int(real(reflectValue.Complex())), nil
	case reflect.Interface, reflect.Pointer:
		if reflectValue.IsNil() {
			return 0, errors.New("error getting the interface/pointer size, it is null")
		} else {
			return ToLengthWithErr(reflectValue.Elem().Interface())
		}
	default:
		return 0, errors.New(fmt.Sprintf("error getting %s size, type not supported", reflectValue.Kind().String()))
	}
}

func implementsLener(t reflect.Type) bool {
	if t == nil {
		return false
	}
	return t.Implements(reflect.TypeOf((*Lener)(nil)).Elem())
}

func implementsSizer(t reflect.Type) bool {
	if t == nil {
		return false
	}
	return t.Implements(reflect.TypeOf((*Sizer)(nil)).Elem())
}

func implementsLengther(t reflect.Type) bool {
	if t == nil {
		return false
	}
	return t.Implements(reflect.TypeOf((*Lengther)(nil)).Elem())
}

func resolveLengthImplementsIfPresent(t reflect.Type, v reflect.Value) (int, bool, error) {
	// Order matters: Len() is the most idiomatic and semantic "length".
	if implementsLener(t) {
		n := v.Interface().(Lener).Len()
		if n < 0 {
			return 0, true, errors.New("error getting length, Len() returned negative value")
		}
		return n, true, nil
	}

	// Size() is often used as "bytes size" (optional but useful).
	if implementsSizer(t) {
		n := v.Interface().(Sizer).Size()
		if n < 0 {
			return 0, true, errors.New("error getting length, Size() returned negative value")
		}
		return n, true, nil
	}

	// Length() is custom convention (optional).
	if implementsLengther(t) {
		n := v.Interface().(Lengther).Length()
		if n < 0 {
			return 0, true, errors.New("error getting length, Length() returned negative value")
		}
		return n, true, nil
	}

	return 0, false, nil
}
