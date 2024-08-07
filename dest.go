package converter

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// ToDest converts value 'a' into destination 'dest', returning an error if conversion is not successful.
// If an error arises during the conversion, it panics.
//
// It leverages the use of `ToDestWithErr()` function for performing the conversion and error handling.
//
// Parameters:
//   - a: Value of any type which is to be converted into destination 'dest'.
//   - dest: Destination variable where the converted value will be populated.
//
// Example:
//
//	var dest int
//	a := "10"
//	ToDest(a, &dest) // dest is now 10
//
//	var strDest string
//	ToDest(a, &strDest) // returns panic due to error in converting a from string to integer
func ToDest(a, dest any) {
	err := ToDestWithErr(a, dest)
	if err != nil {
		panic(err)
	}
}

// ToDestWithErr attempts to convert a given value to a specified destination type. The function uses reflection to
// determine the types of the given value and the destination. It checks if the destination is a pointer and whether
// it's nil. It also checks if the given value is valid and not nil.
// The function handles conversions to various types such as Struct, Map, Array, Slice, Boolean, Integer, Unsigned
// Integer, Float, Complex, and String by leveraging auxiliary functions, namely ToBytesWithErr, ToStringWithErr,
// ToBoolWithErr, ToIntWithErr, ToUintWithErr, and ToFloat64WithErr.
// If the given value cannot be converted to the destination type, the function returns an error.
//
// Parameters:
//   - a: The value of any type to be converted to the destination type.
//   - dest: The destination variable where the converted value will be placed. It needs to be a pointer.
//
// Returns:
//   - error: An error is returned in case of failure to convert.
//
// Example:
//
//	var dest string
//	a := 10
//	err := ToDestWithErr(a, &dest)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(dest) // "10"
//
// For the above example, the function converts the integer 10 to a string and assigns it to the `dest` variable.
func ToDestWithErr(a, dest any) error {
	reflectValue := reflect.ValueOf(a)
	reflectDest := reflect.ValueOf(dest)

	if reflectDest.Kind() != reflect.Ptr {
		return errors.New("dest is not a pointer")
	} else if reflectDest.IsNil() {
		return errors.New("dest is nil")
	} else if !reflectValue.IsValid() ||
		(reflectValue.Kind() == reflect.Ptr || reflectValue.Kind() == reflect.Interface) && reflectValue.IsNil() {
		return errors.New("A is nil or invalid")
	}

	switch reflectDest.Elem().Kind() {
	case reflect.Struct, reflect.Map, reflect.Array, reflect.Slice:
		bs, err := ToBytesWithErr(a)
		if err != nil {
			return err
		}
		return json.Unmarshal(bs, dest)
	case reflect.String:
		s, err := ToStringWithErr(a)
		if err != nil {
			return err
		}
		reflectDest.Elem().Set(reflect.ValueOf(s))
	case reflect.Bool:
		b, err := ToBoolWithErr(a)
		if err != nil {
			return err
		}
		reflectDest.Elem().Set(reflect.ValueOf(b))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := ToIntWithErr(a)
		if err != nil {
			return err
		}
		reflectDest.Elem().Set(reflect.ValueOf(i).Convert(reflectDest.Elem().Type()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		ui, err := ToUintWithErr(a)
		if err != nil {
			return err
		}
		reflectDest.Elem().Set(reflect.ValueOf(ui).Convert(reflectDest.Elem().Type()))
	case reflect.Float32, reflect.Float64:
		f, err := ToFloat64WithErr(a)
		if err != nil {
			return err
		}
		reflectDest.Elem().Set(reflect.ValueOf(f).Convert(reflectDest.Elem().Type()))
	case reflect.Complex64, reflect.Complex128:
		c, err := ToComplex128WithErr(a)
		if err != nil {
			return err
		}
		reflectDest.Elem().Set(reflect.ValueOf(c).Convert(reflectDest.Elem().Type()))
	case reflect.Interface:
		reflectDest.Elem().Set(reflectValue)
	default:
		return fmt.Errorf("error convert to dest, unsupported type %s", reflectValue.Kind().String())
	}

	return nil
}
