package converter

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// CouldBeComplex determines if the provided parameter can be converted to a complex128 value without producing an error.
// It does so by calling the ToComplex128WithErr function and checking whether the error returned by this function is nil.
//
// Parameters:
//   - a: An interface value to be checked for complex128 convertibility.
//
// Returns:
//   - bool: A boolean value indicating whether the value could be converted to a complex128 without an error.
//
// Example:
//
//	var x string = "5.2"
//	y := 10
//	z := complex128(1+2i)
//	a := "not a numeric string"
//	fmt.Println(CouldBeComplex(x)) // true, as "5.2" can be converted to complex128
//	fmt.Println(CouldBeComplex(y)) // true, as 10 can be converted to complex128
//	fmt.Println(CouldBeComplex(z)) // true, as it's already complex128
//	fmt.Println(CouldBeComplex(a)) // false, "not a numeric string" cannot be converted to complex128
func CouldBeComplex(a any) bool {
	_, err := ToComplex128WithErr(a)
	return err == nil
}

// ToComplex64 tries to convert a given value to a complex64 value based on its type.
// This function leverages ToComplex64WithErr which converts the value to complex64,
// and if there is an error, this function will panic.
//
// Parameters:
//   - a: The value of any type to be converted to a complex value.
//
// Returns:
//   - complex64: The complex64 value representation of the provided value.
//
// Panic:
//   - This function will panic if there is a failure to convert the provided value to complex64.
//
// Example:
//
//	var str string = "5+3i"
//	complexNum := ToComplex64(str)
//	fmt.Printf("The complex number is: %v\n", complexNum) // The complex number is: (5+3i)
//
//	num := 10
//	complexNum = ToComplex64(num)
//	fmt.Printf("The complex number is: %v\n", complexNum) // The complex number is: (10+0i)
//
//	try {
//	    z := nil
//	    complexNum = ToComplex64(z) // This will cause a panic
//	} catch(err) {
//	    fmt.Println(err.Error())
//	}
func ToComplex64(a any) complex64 {
	c64, err := ToComplex64WithErr(a)
	if err != nil {
		panic(err)
	}
	return c64
}

// ToComplex64WithErr tries to convert a given value to a complex64 value based on its type.
// This function leverages ToComplex128WithErr which converts the value to complex128, and
// then this function reduces the precision of that result to complex64 representation.
// If there is an error, it will be returned along with the zero-valued complex64.
//
// Parameters:
//   - a: The value of any type to be converted to a complex value.
//
// Returns:
//   - complex64: The complex64 value representation of the provided value.
//   - error: An error that will be returned in case of failure to convert.
//
// Example:
//
//	var str string = "5+3i"
//	complexNum, err := ToComplex64WithErr(str)
//	if nil != err {
//	    fmt.Printf("Could not convert value: %s\n", err.Error())
//	} else {
//	    fmt.Printf("The complex number is: %v\n", complexNum) // The complex number is: (5+3i)
//	}
//
//	num := 10
//	complexNum, err = ToComplex64WithErr(num)
//	if nil != err {
//	    fmt.Printf("Could not convert value: %s\n", err.Error())
//	} else {
//	    fmt.Printf("The complex number is: %v\n", complexNum) // The complex number is: (10+0i)
//	}
func ToComplex64WithErr(a any) (complex64, error) {
	c128, err := ToComplex128WithErr(a)
	return complex64(c128), err
}

// ToComplex128 attempts to convert a given value to a complex128 value based on its type.
// It uses the ToComplex128WithErr function to perform the conversion.
// If there is an error during the conversion, it will panic and output the respective error.
//
// Parameters:
//   - a: The value of any type to be converted to a complex128 value.
//
// Returns:
//   - complex128: The complex128 value representation of the provided value.
//
// Panic:
//   - If there is an error during the conversion, the function will panic.
//
// Example:
//
//	var str string = "5+3i"
//	c128 := ToComplex128(str)
//	fmt.Printf("Converted value is: %v\n", c128) // Output: Converted value is: (5+3i)
//
//	num := 10
//	c128 = ToComplex128(num)
//	fmt.Printf("Converted value is: %v\n", c128) // Output: Converted value is: (10+0i)
func ToComplex128(a any) complex128 {
	c128, err := ToComplex128WithErr(a)
	if err != nil {
		panic(err)
	}
	return c128
}

// ToComplex128WithErr tries to convert a given value to a complex128 value based on its type.
// The function uses the reflect package to determine the kind of the given value,
// and then uses type assertion to convert the value to complex128
//
// The function handles different kinds of values including:
//   - String: If the string can be parsed to a complex number, it's done or else an error is returned
//   - Integer: The integer value is converted to complex128 with imaginary part as 0
//   - Unsigned Integer: The Unsigned integer value is converted to complex128 with imaginary part as 0
//   - Float : The Float value is converted to complex128 with imaginary part as 0
//   - Complex64 and Complex128: The Complex value is returned as is
//   - Boolean: If true, complex128 represents 1, if false, represents 0
//   - Interface or Pointer: If the underlying value is nil, an error is returned otherwise the function is recursively
//     called on the Elem() of the value
//
// For all other types, an error is returned.
//
// Parameters:
//   - a: The value of any type to be converted to a complex value .
//
// Returns:
//   - complex128: The complex value representation of the provided value.
//   - error: An error is returned in case of failure to convert.
//
// Example:
//
//	var str string = "5+3i"
//	complexNum, err := ToComplex128WithErr(str)
//	if nil != err {
//	    fmt.Printf("Could not convert value: %s\n", err.Error())
//	} else {
//	    fmt.Printf("The complex number is: %v\n", complexNum) // The complex number is: (5+3i)
//	}
//
//	num := 10
//	complexNum, err = ToComplex128WithErr(num)
//	if nil != err {
//	    fmt.Printf("Could not convert value: %s\n", err.Error())
//	} else {
//	    fmt.Printf("The complex number is: %v\n", complexNum) // The complex number is: (10+0i)
//	}
func ToComplex128WithErr(a any) (complex128, error) {
	reflectValue := reflect.ValueOf(a)

	switch reflectValue.Kind() {
	case reflect.String:
		c, err := strconv.ParseComplex(reflectValue.String(), 128)
		if err != nil {
			return 0, err
		}
		return c, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return complex(float64(reflectValue.Int()), 0), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return complex(float64(reflectValue.Uint()), 0), nil
	case reflect.Float32, reflect.Float64:
		return complex(reflectValue.Float(), 0), nil
	case reflect.Complex64, reflect.Complex128:
		return reflectValue.Complex(), nil
	case reflect.Bool:
		if reflectValue.Bool() {
			return complex(float64(1), 0), nil
		}
		return complex(float64(0), 0), nil
	case reflect.Interface, reflect.Pointer:
		if reflectValue.IsNil() {
			return 0, errors.New("error convert to complex, it is null")
		}
		return ToComplex128WithErr(reflectValue.Elem().Interface())
	default:
		return 0, fmt.Errorf("error convert to complex, unsupported type %s", reflectValue.Kind().String())
	}
}
