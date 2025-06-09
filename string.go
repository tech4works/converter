package converter

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// TODO: Complementar exemplos

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
	if a == nil {
		return "", errors.New("error convert to string, it is null")
	}

	reflectValue := reflect.ValueOf(a)

	if reflectValue.Kind() == reflect.Ptr || reflectValue.Kind() == reflect.Interface {
		if reflectValue.IsNil() {
			return "", errors.New("error convert to string, it is null")
		}
		return ToStringWithErr(reflectValue.Elem().Interface())
	}

	if stringer, ok := a.(fmt.Stringer); ok {
		return stringer.String(), nil
	} else if err, ok := a.(error); ok {
		return err.Error(), nil
	}

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
	case reflect.Slice:
		if reflectValue.Type().Elem().Kind() == reflect.Uint8 {
			return string(reflectValue.Bytes()), nil
		}
		marshal, _ := json.Marshal(reflectValue.Interface())
		return string(marshal), nil
	case reflect.Array:
		if reflectValue.Type().Elem().Kind() == reflect.Uint8 {
			bytes := make([]byte, reflectValue.Len())
			for i := 0; i < reflectValue.Len(); i++ {
				bytes[i] = byte(reflectValue.Index(i).Uint())
			}
			return string(bytes), nil
		}
		marshal, _ := json.Marshal(reflectValue.Interface())
		return string(marshal), nil
	case reflect.Map, reflect.Struct:
		marshal, _ := json.Marshal(reflectValue.Interface())
		return string(marshal), nil
	default:
		return "", fmt.Errorf("error convert to string, unsupported type %s", reflectValue.Kind().String())
	}
}

// ToCompactString converts a given value to a string representation in a compact form.
// It uses the ToCompactStringWithErr function to convert the value into compact string .
// If it encounters an error during this process, it will panic.
//
// The compact form of a string is obtained by removing extra white spaces.
//
// Please note that this conversion process utilizes reflection to inspect the type of the value, apply appropriate
// conversion methods and catch any potential errors.
// If any unsupported types are encountered during this conversion process, this function will panic.
//
// Parameters:
//   - a: The value of any type to be converted to a string in a compact form.
//
// Returns:
//   - string: The compact string representation of the provided value.
//
// Panics:
//   - If it encounters an error during this process, it will panic.
//
// Example:
//
//	 var x map[string]string = map[string]string{
//			"a": "apple",
//			"b": "banana",
//		}
//	 s := ToCompactString(x)     // "{"a":"apple","b":"banana"}"
//	 fmt.Println(s)
//
//	 y := "Hello,         World!   How are you?     "
//	 s = ToCompactString(x)       // "Hello, World! How are you?"
//	 fmt.Println(s)
func ToCompactString(a any) string {
	str, err := ToCompactStringWithErr(a)
	if err != nil {
		panic(err)
	}
	return str
}

// ToCompactStringWithErr converts a given value to a string
// representation in a compact form, where all white spaces are replaced by single spaces.
// It leverages the ToStringWithErr function for conversion,
// which caters to various types handled via reflection.
// Should the conversion encounter an error, it returns the error.
//
// The compact form of a string is one where multiple consecutive white spaces
// get replaced by a single space. The function utilizes regular expressions
// to achieve this.
//
// Parameters:
//   - a: The value of any type to be converted to a compact string.
//
// Returns:
//   - string: The compact string representation of the provided value.
//   - error: An error is returned in case of failure to convert.
//
// Panics:
//   - While the function itself doesn't panic, be aware that the underlying
//     ToStringWithErr function might panic for unsupported types.
//
// Example:
//
//	var x map[string]string = map[string]string{
//		"first": "Hello,     World!   ",
//		"second": "How     are you?  ",
//	}
//	s, err := ToCompactStringWithErr(x)
//	if err != nil {
//		fmt.Println(err)
//	}
//	// "{"first":"Hello, World!","second":"How are you?"}", nil
//	fmt.Println(s, err)
//
//	y := "Hello,       World!    How   are    you?       "
//	s, err = ToCompactStringWithErr(y)
//	// "Hello, World! How are you?", nil
//	fmt.Println(s, err)
func ToCompactStringWithErr(a any) (string, error) {
	s, err := ToStringWithErr(a)
	if err != nil {
		return "", err
	}

	regex := regexp.MustCompile(`\s+`)

	s = strings.TrimSpace(s)
	s = regex.ReplaceAllString(s, " ")

	return s, nil
}
