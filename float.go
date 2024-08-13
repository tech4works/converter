package converter

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// CouldBeFloat determines whether a given value can be converted to a float64.
// It uses the ToFloat64WithErr function to convert the value and if conversion
// is successful, it returns true. Otherwise, it returns false.
//
// Parameters:
//   - a: Any interface value to be checked whether it could be converted to float64.
//
// Returns:
//   - bool: A boolean value indicating whether the value could be converted to float64.
//
// Example:
//
//	x := "10.2"
//	y := "hello"
//	fmt.Println(CouldBeFloat(x)) // true
//	fmt.Println(CouldBeFloat(y)) // false
func CouldBeFloat(a any) bool {
	_, err := ToFloat64WithErr(a)
	return err == nil
}

// ToFloat32 attempts to convert a given value of any type to a float32 value.
// This function employs the ToFloat32WithErr function to perform the conversion and
// will panic if an error occurs during the conversion process.
//
// Parameters:
//   - a: Any interface value that should be attempted to convert to a float32.
//
// Returns:
//   - float32: The converted float32 value.
//     If the conversion fails, this function will panic.
//
// Panics:
//   - If the underlying function ToFloat32WithErr(a) fails and returns an error, this function will panic.
//
// Example:
//
//	var num int = 10
//	f := ToFloat32(num) // 10.0
//
//	s := "10.5"
//	f = ToFloat32(s) // 10.5
func ToFloat32(a any) float32 {
	f32, err := ToFloat32WithErr(a)
	if err != nil {
		panic(err)
	}
	return f32
}

// ToFloat32WithErr attempts to convert the given value of any type to a float32 value.
// This function employs the ToFloat64WithErr function to convert the value to float64 first then it
// casts the float64 value to float32.
//
// Parameters:
//   - a: Any interface value that should be attempted to convert to a float32.
//
// Returns:
//   - float32: The converted float32 value.
//   - error: An error indicating any issues that occurred during conversion, such as unsupported type
//     or attempting to convert a nil pointer or interface. If conversion is successful, a nil error will be returned.
//
// Example:
//
//	var num int = 10
//	f, err := ToFloat32WithErr(num)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(f) // 10.0
//
//	s := "10.5"
//	f, err = ToFloat32WithErr(s)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(f) // 10.5
func ToFloat32WithErr(a any) (float32, error) {
	f64, err := ToFloat64WithErr(a)
	return float32(f64), err
}

// ToFloat64 attempts to convert a value of any type to a float64 value.
// It employs the ToFloat64WithErr function to accomplish the conversion. If the conversion encounters an error, this
// function will panic.
//
// This function is able to convert varied data types, including Strings, Integers, Unsigned Integers, Floats,
// Complex Numbers, and Booleans. For more details about the conversion process, refer to the GoDoc for ToFloat64WithErr.
//
// Parameters:
//   - a: Any interface value to be converted to a float64 value.
//
// Returns:
//   - float64: The float64 value converted from the given input. If the conversion failed due to an error, this
//     function will panic.
//
// Panic:
//
//	This function can panic if the conversion encounters an error, such as attempting to convert from an unsupported type
//	or a nil pointer or interface.
//
// Example:
//
//	var num int = 10
//	f := ToFloat64(num)
//	fmt.Println(f) // 10.0
//
//	s := "10.5"
//	f = ToFloat64(s)
//	fmt.Println(f) // 10.5
func ToFloat64(a any) float64 {
	f64, err := ToFloat64WithErr(a)
	if err != nil {
		panic(err)
	}
	return f64
}

// ToFloat64WithErr attempts to convert the given value of any type to a float64 value. The function uses reflection
// to determine the type of the given input and applies relevant conversion logic based on the identified type.
// This function can convert a varied list of types, including String, Integers, Unsigned Integers, Floats,
// Complex Numbers, and Booleans.
//
// Note that if the input value can be a pointer or an interface that points to nil, or if the type of data
// isn't supported for conversion, an error will be returned.
//
// Parameters:
//   - a: Any interface value that should be attempted to convert to a float64.
//
// Returns:
//   - float64: The converted float64 value.
//   - error: An error indicating any issues that occurred during conversion, such as unsupported type
//     or attempting to convert a nil pointer or interface. If conversion is successful, a nil error will be returned.
//
// Example:
//
//	var num int = 10
//	f, err := ToFloat64WithErr(num)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(f) // 10.0
//
//	s := "10.5"
//	f, err = ToFloat64WithErr(s)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(f) // 10.5
func ToFloat64WithErr(a any) (float64, error) {
	reflectValue := reflect.ValueOf(a)

	switch reflectValue.Kind() {
	case reflect.String:
		return strconv.ParseFloat(reflectValue.String(), 64)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(reflectValue.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float64(reflectValue.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return reflectValue.Float(), nil
	case reflect.Complex64, reflect.Complex128:
		complexNum := reflectValue.Complex()
		return real(complexNum), nil
	case reflect.Bool:
		if reflectValue.Bool() {
			return 1, nil
		}
		return 0, nil
	case reflect.Array, reflect.Slice:
		if reflectValue.Type().Elem().Kind() == reflect.Uint8 {
			return strconv.ParseFloat(string(reflectValue.Bytes()), 64)
		}
		return 0, fmt.Errorf("error convert to float, unsupported type %s", reflectValue.Kind().String())
	case reflect.Interface, reflect.Pointer:
		if reflectValue.IsNil() {
			return 0, errors.New("error convert to float, it is null")
		}
		return ToFloat64WithErr(reflectValue.Elem().Interface())
	default:
		return 0, fmt.Errorf("error convert to float, unsupported type %s", reflectValue.Kind().String())
	}
}
