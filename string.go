package converter

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// CouldBeString inspects if a given value could be converted to a string representation with an error.
// The function uses the ToStringWithErr function to attempt the conversion and
// returns a boolean indicating if the operation is successful or not.
//
// Parameters:
//   - a: Any interface value to be checked for string convertibility.
//
// Returns:
//   - bool: A boolean value indicating whether the value could be successfully converted to a string.
//
// Example:
//
//	var x int = 10
//	y := make(chan int)
//	fmt.Println(CouldBeString(x)) // true
//	fmt.Println(CouldBeString(y)) // false
func CouldBeString(a any) bool {
	_, err := ToStringWithErr(a)
	return err == nil
}

// ToString attempts to convert any given input value to a string representation.
// It uses the ToStringWithErr function for this purpose. If successful, the function returns the string
// representation. However, if the conversion encounters an error, the function panics.
//
// It is important to note that while this function provides a convenient way to convert a value
// to a string without having to handle errors explicitly, it may lead to runtime panics if used
// with unsupported types or nil values. Therefore, use it judiciously or consider using ToStringWithErr
// function directly if error handling is preferred.
//
// Parameters:
//   - a: The value of any type that needs to be converted to a string.
//
// Returns:
//   - string: The string representation of the input value.
//
// Panics:
//   - When the conversion process encounters an error (e.g., when attempting to convert a nil value or unsupported type).
//
// Example:
//
//	var a int = 10
//	b := "World"
//	fmt.Println(ToString(a)) // "10"
//	fmt.Println(ToString(b)) // "World"
//
//	c := func() {}
//	fmt.Println(ToString(c)) // Following line causes a runtime panic as the function is unable to handle the conversion
func ToString(a any) string {
	str, err := ToStringWithErr(a)
	if err != nil {
		panic(err)
	}
	return str
}

// ToStringWithErr converts a given value to a string representation based on its type.
// The function uses reflection to determine the type of the value and applies the appropriate
// conversion method.
// If the conversion fails, the function returns an error.
//
// The function handles the following types:
//   - String: Returns the string as is.
//   - Integers (of various sizes): Converts the integer to a string.
//   - Unsigned Integers (of various sizes): Converts the unsigned integer to a string.
//   - Floats (32 and 64 bits): Converts the float to a string.
//   - Complex numbers (64 and 128 bits): Converts the complex number to a string.
//   - Boolean: Converts the boolean to a string.
//   - Arrays and Slices: If an element type is uint8, converts the byte slice to a string. For other types, marshals
//     the value to JSON.
//   - Maps and Structs: Marshals the value to JSON.
//   - Pointers and Interfaces: If the value is nil, returns an error. Otherwise, attempts to convert the element value
//     to a string.
//   - Other types: Returns an error indicating an unsupported type.
//
// Parameters:
//   - a: The value of any type to be converted to a string.
//
// Returns:
//   - string: The string representation of the provided value.
//   - error: An error is returned in case of failure to convert.
//
// Example:
//
//	var x int = 10
//	y := "Hello"
//	fmt.Println(ToStringWithErr(x)) // "10", nil
//	fmt.Println(ToStringWithErr(y)) // "Hello", nil
func ToStringWithErr(a any) (string, error) {
	reflectValue := reflect.ValueOf(a)

	switch reflectValue.Kind() {
	case reflect.String:
		return reflectValue.String(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(reflectValue.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(reflectValue.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(reflectValue.Float(), 'g', -1, 64), nil
	case reflect.Complex64, reflect.Complex128:
		return strconv.FormatComplex(reflectValue.Complex(), 'g', -1, 64), nil
	case reflect.Bool:
		return strconv.FormatBool(reflectValue.Bool()), nil
	case reflect.Array, reflect.Slice:
		if reflectValue.Type().Elem().Kind() == reflect.Uint8 {
			return string(reflectValue.Bytes()), nil
		}
		marshal, _ := json.Marshal(reflectValue.Interface())
		return string(marshal), nil
	case reflect.Map, reflect.Struct:
		marshal, _ := json.Marshal(reflectValue.Interface())
		return string(marshal), nil
	case reflect.Ptr, reflect.Interface:
		if reflectValue.IsNil() {
			return "", errors.New("error convert to string, it is null")
		}
		return ToStringWithErr(reflectValue.Elem().Interface())
	default:
		return "", fmt.Errorf("error convert to string, unsupported type %s", reflectValue.Kind().String())
	}
}
