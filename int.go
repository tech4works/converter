package converter

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// CouldBeInt attempts to convert a given value to an integer representation.
// It uses the ToIntWithErr function to perform the conversion and returns a boolean indicating the success or failure
// of the conversion.
//
// Parameters:
//   - a: Any interface value to be attempted for conversion into an integer.
//
// Returns:
//   - bool: A boolean value indicating whether the conversion to an integer was successful or not.
//
// Example:
//
//	i := "1234"
//	fmt.Println(CouldBeInt(i)) // true
//
//	j := "Hello"
//	fmt.Println(CouldBeInt(j)) // false
func CouldBeInt(a any) bool {
	_, err := ToIntWithErr(a)
	return err == nil
}

// ToInt converts a value of any type to an integer form.
//
// This function internally uses ToIntWithErr function to perform the conversion. If there are any errors encountered during this conversion process, namely if the input parameter 'a' is of unsupported type, the function will cause a panic.
//
// Parameters:
//   - a: The value to be converted into an integer. This can be of any type.
//
// Returns:
//   - int: The integer representation of the provided value. If an error occurs during conversion, this function will panic.
//
// Panics:
//   - If the underlying function ToIntWithErr returns an error during the conversion process.
//
// Example:
//
//		str := "100"
//		i := ToInt(str)
//		fmt.Println(i)  // 100
//
//		arr := [3]int{1,2,3}
//		i = ToInt(arr)
//	 Caught panic:  error convert to int, unsupported type array
func ToInt(a any) int {
	i, err := ToIntWithErr(a)
	if err != nil {
		panic(err)
	}
	return i
}

// ToIntWithErr attempts to convert a given value to an integer representation and returns an error if unsuccessful.
//
// The function uses reflect.ValueOf() method to get the reflection Value of the input parameter 'a' and switch-case
// statement to handle different kinds of value types: String, Int, Uint, Float, Complex, Bool, Interface and Pointer.
//
// In the case of Interface and Pointer, if the value is nil, it returns an error, otherwise,
// the function recursively calls itself with the contained value in Interface and Pointer.
//
// If the kind does not match any cases, a default case returns an error indicating an unsupported type.
//
// Parameters:
//   - a: The value of any type to be converted to an integer.
//
// Returns:
//   - int: The integer representation of the provided value.
//   - error: An error indicating either the value is nil or unsupported type for conversion.
//
// Example:
//
//	text := "123"
//	i, err := ToIntWithErr(text)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(i) // 123
//
//	for j := 0; j < 5; j++ {
//		i, err := ToIntWithErr(j)
//		if err != nil {
//		fmt.Println(err)
//		}
//		fmt.Println(i) // prints 0, 1, 2, ..., 4
//	}
func ToIntWithErr(a any) (int, error) {
	reflectValue := reflect.ValueOf(a)

	switch reflectValue.Kind() {
	case reflect.String:
		return strconv.Atoi(reflectValue.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(reflectValue.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return int(reflectValue.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return int(reflectValue.Float()), nil
	case reflect.Complex64, reflect.Complex128:
		return int(real(reflectValue.Complex())), nil
	case reflect.Bool:
		if reflectValue.Bool() {
			return 1, nil
		}
		return 0, nil
	case reflect.Interface, reflect.Pointer:
		if reflectValue.IsNil() {
			return 0, errors.New("error convert to int, it is null")
		}
		return ToIntWithErr(reflectValue.Elem().Interface())
	default:
		return 0, fmt.Errorf("error convert to int, unsupported type %s", reflectValue.Kind().String())
	}
}

// ToInt8 attempts to convert a given value to an int8 representation.
//
// The function leverages the ToInt8WithErr function to perform the conversion and
// captures any potential error. If there is an error during the conversion, the function
// will panic.
//
// Parameters:
//   - a: The value of any type to be converted to int8.
//
// Returns:
//   - int8: The int8 representation of the provided value.
//
// Panics:
//   - If the conversion to int8 is unsuccessful, the function will panic.
//
// Example:
//
//		i := ToInt8("123")
//		fmt.Println(i) // 123
//
//		var a interface{}
//		a = "45"
//		j := ToInt8(a)
//		fmt.Println(j) // 45
//
//	 Please note that for the second example we are using an interface{} value.
//	 In this case we still can convert it because the underlying value of the
//	 pointer is a string which can be converted to int8.
func ToInt8(a any) int8 {
	i, err := ToInt8WithErr(a)
	if err != nil {
		panic(err)
	}
	return i
}

// ToInt8WithErr tries to convert a value of any type to int8.
//
// It leverages the function ToIntWithErr for this task, which
// attempts to convert the value to an integer type. The result is
// then converted to an int8 type. If the original conversion
// fails, the error is passed through.
//
// Parameters:
//   - a: The value of any type that is to be converted to int8.
//
// Returns:
//   - int8: The int8 representation of the provided value.
//   - error: An error indicating that the conversion operation was
//     unsuccessful. Error can occur if the type of 'a' does not
//     support conversion to int type or if 'a' is nil.
//
// Example:
//
//	strVal := "128"
//	newVal, err := ToInt8WithErr(strVal)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(newVal) // 128
//	}
//
//	ptrVal := new(int)
//	*ptrVal = 32
//	newVal, err = ToInt8WithErr(ptrVal)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(newVal) // 32
//	}
func ToInt8WithErr(a any) (int8, error) {
	i, err := ToIntWithErr(a)
	return int8(i), err
}

// ToInt16 converts a given value to an int16 type. It uses the ToInt16WithErr function
// to attempt to convert the value and panics if the conversion is unsuccessful.
//
// Parameters:
//   - a: Any value to be converted to an int16.
//
// Returns:
//   - int16: The int16 representation of the provided value.
//
// Panics:
//   - If the provided value is of an unsupported type or can't be converted, the function will panic.
//
// Example:
//
//	var a int32 = 10
//	var b string = "20"
//	fmt.Println(ToInt16(a)) // 10
//	fmt.Println(ToInt16(b)) // 20
//
//	c := false
//	fmt.Println(ToInt16(c)) // Causes a panic
func ToInt16(a any) int16 {
	i, err := ToInt16WithErr(a)
	if err != nil {
		panic(err)
	}
	return i
}

// ToInt16WithErr attempts to convert a given value to a int16 representation and returns an error if unsuccessful.
//
// The function utilizes the ToIntWithErr function to convert the given value
// to an int type first, and then convert the result to an int16 type.
//
// If the given value is of a type that cannot be converted to an integer,
// the function will return an error.
//
// Parameters:
//   - a: The value of any type to be converted to an int16.
//
// Returns:
//   - int16: The int16 representation of the provided value.
//   - error: An error indicating either the value is of an unsupported type for conversion.
//
// Example:
//
//	i, err := ToInt16WithErr("123")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(i) // prints 123
//
//	i, err := ToInt16WithErr(true)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(i) // Erroneous case, will print an error
//	}
func ToInt16WithErr(a any) (int16, error) {
	i, err := ToIntWithErr(a)
	return int16(i), err
}

// ToInt32 converts a given value to an integer of type int32.
// If the conversion fails, this function will panic.
//
// Parameters:
//   - a: The value of any type to be converted to an int32.
//
// Returns:
//   - int32: The int32 representation of the provided value. This function will panic if the value cannot be converted.
//
// Panic:
//   - This function will panic if it cannot convert the provided value to an int32.
//
// Example:
//
//	text := "123"
//	i := ToInt32(text)
//	fmt.Println(i) // Output: 123
//
//	n := 456
//	m := ToInt32(n)
//	fmt.Println(m) // Output: 456
//
//	// The following will panic as string "hello" is not a valid integer
//	panicValue := "hello"
//	fmt.Println(ToInt32(panicValue))
func ToInt32(a any) int32 {
	i, err := ToInt32WithErr(a)
	if err != nil {
		panic(err)
	}
	return i
}

// ToInt32WithErr attempts to convert a given value into an int32 representation and returns an error if unsuccessful.
//
// The function makes use of the ToIntWithErr function to try and convert the input parameter 'a' to an integer.
// The result returned from ToIntWithErr is then typecast to an int32.
// If the call to ToIntWithErr results in an error, the error is passed along.
//
// Parameters:
//   - a: The value of any type to be converted to an int32.
//
// Returns:
//   - int32: The int32 representation of the provided value.
//   - error: An error indicating a failure in the conversion process.
//
// Example:
//
//		text := "123"
//		i, err := ToInt32WithErr(text)
//		if err != nil {
//			fmt.Println(err)
//		}
//		fmt.Println(i) // Output: 123
//
//		for j := 0; j < 5; j++ {
//			i, err := ToInt32WithErr(j)
//			if err != nil {
//			fmt.Println(err)
//			}
//			fmt.Println(i) // prints 0, 1, 2, ..., 4
//		}
//
//		// The following will yield an error as string "hello" is not a valid integer
//		errorVal := "hello"
//		val, err := ToInt32WithErr(errorVal)
//		if err != nil {
//	     fmt.Println(err)
//	 }
func ToInt32WithErr(a any) (int32, error) {
	i, err := ToIntWithErr(a)
	return int32(i), err
}

// ToInt64 attempts to convert a given value to a 64-bit integer representation.
// It relies on the helper function ToInt64WithErr for its operation, using the latter's output to provide its own.
// In the event that ToInt64WithErr returns an error, ToInt64 will panic and present that error. Otherwise,
// it returns the 64-bit integer value obtained from ToInt64WithErr.
//
// Parameters:
//   - a: The value of any type to be processed and converted into a 64-bit integer.
//
// Returns:
//   - int64: The 64-bit integer representation of the provided value.
//
// Panic:
//   - This function will panic if there is an error while converting the input to a 64-bit integer representation.
//
// Example:
//
//		var a int32 = 1024
//		var b string = "2048"
//		fmt.Println(ToInt64(a)) // Output: 1024
//		fmt.Println(ToInt64(b)) // Output: 2048
//	    fmt.Println(ToInt64(nil)) // Output: panic
func ToInt64(a any) int64 {
	i, err := ToInt64WithErr(a)
	if err != nil {
		panic(err)
	}
	return i
}

// ToInt64WithErr attempts to convert a given value to a 64-bit integer. This function wraps the ToIntWithErr function
// and is similar to it in terms of the conversion process, however, it casts the returned integer from ToIntWithErr to
// int64 before returning.
//
// Parameters:
//   - a: The value of any type to be processed and converted into a 64-bit integer.
//
// Returns:
//   - int64: The 64-bit integer representation of the 'a' value if conversion is successful.
//   - error: An error indicating either an issue in the conversion process, the value is nil, or unsupported type for conversion.
//
// Example:
//
//	var a int32 = 999
//	var b string = "999"
//	i, err := ToInt64WithErr(a)
//	if err != nil {
//	    fmt.Println(err)
//	}
//	fmt.Println(i) // Output: 999
//
//	i, err = ToInt64WithErr(b)
//	if err != nil {
//	    fmt.Println(err)
//	}
//	fmt.Println(i) // Output: 999
//	_, _ = ToInt64WithErr(nil) // This will raise an error because 'nil' cannot be converted to int64.
func ToInt64WithErr(a any) (int64, error) {
	i, err := ToIntWithErr(a)
	return int64(i), err
}
