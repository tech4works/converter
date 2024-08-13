package converter

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// CouldBeUint checks if the value of any type can be converted to a uint without errors.
// It relies on the ToUintWithErr function for the conversion and checks the returned error.
//
// Parameters:
//   - a: The value of any type to be checked for conversion to uint.
//
// Returns:
//   - bool: A boolean indicating whether the value can be converted to uint without errors.
//
// Example:
//
//	var x int = 10
//	var y string = "hello"
//	fmt.Println(CouldBeUint(x)) // true
//	fmt.Println(CouldBeUint(y)) // false
func CouldBeUint(a any) bool {
	_, err := ToUintWithErr(a)
	return err == nil
}

// ToUint converts any given value to an uint type.
// If the conversion is not possible or error is encountered, the function will panic.
// It uses the ToUintWithErr function internally for conversion.
//
// Parameters:
//   - a: Any value that needs to be converted to uint.
//
// Returns:
//   - uint: The resulting uint value after conversion.
//
// Panic:
//   - This function will panic if there is an error during the conversion process.
//
// Example:
//
//	var s string = "100"
//	var i int = 200
//	fmt.Println(ToUint(s))  // 100
//	fmt.Println(ToUint(i))  // 200
//	try {
//		var x = []int{1,2,3}
//		ToUint(x)
//		// This will panic because slices cannot be converted to uint
//	} catch (err) {
//		fmt.Println(err) // error convert to uint, unsupported type slice
//	}
func ToUint(a any) uint {
	u, err := ToUintWithErr(a)
	if err != nil {
		panic(err)
	}
	return u
}

// ToUintWithErr attempts to convert any given value to an uint type.
// It uses reflection to determine the type of the given value.
// The function supports conversion from various types such as String, Integer, Unsigned Integer, Float, Complex,
// Boolean, Interface and Pointer.
// For Integer, Float and Complex types, if the value is negative, an error is returned.
// For Interface and Pointer types, if the value is nil, an error is returned.
//
// The function also ensures type safety by returning an error for any unsupported type, or
// if any error is encountered during the conversion process.
//
// Parameters:
//   - a: The value of any type to be converted to uint.
//
// Returns:
//   - uint: The converted uint value from the given input.
//   - error: An error is returned in case of failure to convert.
//
// Example:
//
//	var x int = -10
//	var y float64 = 10.5
//	fmt.Println(ToUintWithErr(x)) // Returns: (0, error)
//	fmt.Println(ToUintWithErr(y)) // Returns: (10, nil)
func ToUintWithErr(a any) (uint, error) {
	reflectValue := reflect.ValueOf(a)

	switch reflectValue.Kind() {
	case reflect.String:
		i, err := strconv.Atoi(reflectValue.String())
		if err != nil {
			return 0, err
		} else if i < 0 {
			return 0, fmt.Errorf("error: trying to convert a negative int to uint: got %v", i)
		}
		return uint(i), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := reflectValue.Int()
		if i < 0 {
			return 0, fmt.Errorf("error: trying to convert a negative int to uint: got %v", i)
		}
		return uint(i), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return uint(reflectValue.Uint()), nil
	case reflect.Float32, reflect.Float64:
		f := reflectValue.Float()
		if f < 0 {
			return 0, fmt.Errorf("error: trying to convert a negative float to uint: got %v", f)
		}
		return uint(f), nil
	case reflect.Complex64, reflect.Complex128:
		c := real(reflectValue.Complex())
		if c < 0 {
			return 0, fmt.Errorf("error: trying to convert a negative complex number to uint: got %v", c)
		}
		return uint(c), nil
	case reflect.Bool:
		if reflectValue.Bool() {
			return 1, nil
		}
		return 0, nil
	case reflect.Array, reflect.Slice:
		if reflectValue.Type().Elem().Kind() == reflect.Uint8 {
			return ToUintWithErr(string(reflectValue.Bytes()))
		}
		return 0, fmt.Errorf("error convert to uint, unsupported type %s", reflectValue.Kind().String())
	case reflect.Interface, reflect.Pointer:
		if reflectValue.IsNil() {
			return 0, errors.New("error convert to uint, it is null")
		}
		return ToUintWithErr(reflectValue.Elem().Interface())
	default:
		return 0, fmt.Errorf("error convert to uint, unsupported type %s", reflectValue.Kind().String())
	}
}

// ToUint8 attempts to convert any given value to an uint8 type.
// It uses the ToUint8WithErr function to attempt the conversion,
// and panics if any error is encountered during the conversion process.
//
// Parameters:
//   - a: Any value to be converted to uint8.
//
// Returns:
//   - uint8: The converted uint8 value from the given input.
//
// Panics:
//   - An error is returned from the ToUint8WithErr function during the conversion process.
//
// Example:
//
//	var x int = -10
//	var y float64 = 10.5
//	fmt.Println(ToUint8(x)) // Panics with error
//	fmt.Println(ToUint8(y)) // Returns: 10
func ToUint8(a any) uint8 {
	u, err := ToUint8WithErr(a)
	if err != nil {
		panic(err)
	}
	return u
}

// ToUint8WithErr attempts to convert a given value to uint8 type.
// It uses the function ToUintWithErr for the conversion process.
// If there are any errors encountered during the conversion, they are returned as the second value.
//
// Parameters:
//   - a: The value to be converted to uint8 type.
//
// Returns:
//   - uint8: The converted uint8 value from the input, if the conversion was successful.
//   - error: An error encountered during the conversion, if any.
//
// Example:
//
//	var x int = 10
//	var y string = "not a number"
//	fmt.Println(ToUint8WithErr(x)) // Returns: (10, nil)
//	fmt.Println(ToUint8WithErr(y)) // Returns: (0, error)
func ToUint8WithErr(a any) (uint8, error) {
	u, err := ToUintWithErr(a)
	return uint8(u), err
}

// ToUint16 attempts to convert any given value to an uint16 type.
// It uses the ToUint16WithErr function to handle the conversion.
// If an error is encountered during the conversion process, the function will panic.
//
// Parameters:
//   - a: The value of any type to be converted to uint16.
//
// Returns:
//   - uint16: The converted uint16 value from the given input.
//
// Panics:
//   - If the ToUint16WithErr function returns an error,
//     ToUint16 will panic with the error as its argument.
//
// Example:
//
//	var x int = -10
//	var y float64 = 10.5
//	fmt.Println(ToUint16(x)) // Panics with error
//	fmt.Println(ToUint16(y)) // Returns: 10
func ToUint16(a any) uint16 {
	u, err := ToUint16WithErr(a)
	if err != nil {
		panic(err)
	}
	return u
}

// ToUint16WithErr attempts to convert any given value to an uint16 type.
// This function utilizes the ToUintWithErr function, which supports conversion from
// various types such as String, Integer, Unsigned Integer, Float, Complex, Boolean, Interface, and Pointer.
//
// Parameters:
//   - a: The value of any type to be converted to uint16.
//
// Returns:
//   - uint16: The converted uint16 value from the given input.
//   - error: An error is returned in case of failure to convert.
//
// Example:
//
//	var x int = -10
//	var y float64 = 10.5
//	fmt.Println(ToUint16WithErr(x)) // Returns: (0, error)
//	fmt.Println(ToUint16WithErr(y)) // Returns: (10, nil)
func ToUint16WithErr(a any) (uint16, error) {
	u, err := ToUintWithErr(a)
	return uint16(u), err
}

// ToUint32 attempts to convert any given value to an uint32 type.
// It uses the "ToUint32WithErr" function, which converts various types such as String, Integer, Unsigned Integer,
// Float, Complex, Boolean, Interface, and Pointer to uint32 and returns an error value if the conversion fails.
// If the function encounters an error during the conversion, the ToUint32 function will panic and display the error.
//
// Parameters:
//   - a: The value of any type to be converted to uint32
//
// Returns:
//   - uint32: The converted uint32 value from the given input, or panic if an error occurred during the conversion.
//
// Panic:
//   - If the input value cannot be converted to a uint32, the function will panic and display the error returned by "ToUint32WithErr".
//
// Example:
//
//	var x int = -10
//	var y float64 = 10.5
//	fmt.Println(ToUint32(x)) // will panic with an error message: "error: trying to convert a negative int to uint32"
//	fmt.Println(ToUint32(y)) // Returns: 10
func ToUint32(a any) uint32 {
	u, err := ToUint32WithErr(a)
	if err != nil {
		panic(err)
	}
	return u
}

// ToUint32WithErr attempts to convert any given value to a uint32 type.
// The function supports conversion from various types such as String, Integer, Unsigned Integer, Float,
// Complex, Boolean, Interface, and Pointer.
// For Interface and Pointer types, if the value is nil, an error is returned.
// The function also ensures type safety by returning an error for any unsupported type or if any error
// is encountered during the conversion process.
//
// Parameters:
//   - a: The value of any type to be converted to uint32.
//
// Returns:
//   - uint32: The converted uint32 value from the given input.
//   - error: An error is returned in case of failure to convert.
//
// Example:
//
//	var x int = -10
//	var y float64 = 10.5
//	fmt.Println(ToUint32WithErr(x)) // Returns: (0, error)
//	fmt.Println(ToUint32WithErr(y)) // Returns: (10, nil)
func ToUint32WithErr(a any) (uint32, error) {
	u, err := ToUintWithErr(a)
	return uint32(u), err
}

// ToUint64 receives a value of any type and attempts to convert it to an uint64 type.
// It uses the function ToUint64WithErr for this conversion, and if any error occurs during the conversion, it panics
// and stops the execution.
//
// Parameters:
//   - a: This is any value that needs to be converted to an uint64.
//
// Returns:
//   - uint64: This is the converted uint64 value from the given input.
//
// Panics:
//   - The function will panic if there is an error during the conversion.
//
// Example:
//
//	var negativeInt int = -10
//	var floatNum float64 = 10.5
//	fmt.Println(ToUint64(negativeInt)) // Panics
//	fmt.Println(ToUint64(floatNum)) // Prints: 10
func ToUint64(a any) uint64 {
	u, err := ToUint64WithErr(a)
	if err != nil {
		panic(err)
	}
	return u
}

// ToUint64WithErr attempts to convert any given value to an uint64 type.
// It utilizes the function ToUintWithErr for conversion purposes which supports conversion
// from different types such as String, Integer, Unsigned Integer, Float, Complex, Boolean, Interface, and Pointer.
//
// Parameters:
//   - a: The value of any type to be converted to uint64.
//
// Returns:
//   - uint64: The converted uint64 value from the given input.
//   - error: An error is returned in case of failure to convert with an appropriate message.
//
// Possible errors:
//   - The function might return an error if the conversion from the given type to uint64 is not possible.
//   - For Integer, Float, and Complex types, if the value is negative, an error is returned.
//   - For Interface and Pointer types, if the value is nil, an error is returned.
//   - It also ensures type safety by returning an error for any unsupported type.
//
// Example:
//
//	var negativeInt int = -10
//	var floatNum float64 = 10.5
//	result, err := ToUint64WithErr(negativeInt) // Returns: (0,error)
//	if err != nil {
//		fmt.Println(err) // Prints: "error: trying to convert a negative int to uint: got -10"
//	}
//	result, err = ToUint64WithErr(floatNum) // Returns: (10,nil)
//
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(result) // Prints: 10
//	}
func ToUint64WithErr(a any) (uint64, error) {
	u, err := ToUintWithErr(a)
	return uint64(u), err
}
