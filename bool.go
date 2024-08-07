package converter

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// CouldBeBool checks if an arbitrary value can be converted to a boolean value.
//
// This function calls the ToBoolWithErr function to perform the conversion and checks if an error is returned.
// If the error is nil, it means the value can be converted successfully.
//
// Parameters:
// - a: The input value that needs to be checked for conversion to boolean. It can be of any type.
//
// Returns:
// - bool: true if the input value can be converted to boolean, otherwise false.
//
// Examples:
//
//	str := "true"
//	can1 := CouldBeBool(str)  // can1:true
//
//	i := 7
//	can2 := CouldBeBool(i)  // can2:true
//
//	f := 0.0
//	can3 := CouldBeBool(f)  // can3:true
//
//	unsupportedType := []int{1,2,3}
//	can4 := CouldBeBool(unsupportedType)  // can4:false
func CouldBeBool(a any) bool {
	_, err := ToBoolWithErr(a)
	return err == nil
}

// ToBool converts an arbitrary type to a boolean value.
//
// This function uses ToBoolWithErr internally to convert the input value to a boolean.
// If an error is encountered during the conversion, this function will panic with the error.
//
// Parameters:
//   - a: The input value of an arbitrary type to be converted to boolean.
//
// Returns:
//   - bool: The converted boolean value if conversion is successful.
//
// Panic:
//   - This function will panic if the input value cannot be converted to a boolean.
//
// Example:
//
//	str := "true"
//	fmt.Println(ToBool(str)) // true
//
//	i := 7
//	fmt.Println(ToBool(i)) // true
//
//	f := 0.0
//	fmt.Println(ToBool(f)) // false
func ToBool(a any) bool {
	b, err := ToBoolWithErr(a)
	if err != nil {
		panic(err)
	}
	return b
}

// ToBoolWithErr converts an arbitrary type to a boolean value.
//
// This function supports conversion from string, int, float, bool, interface, pointer types to bool type.
// For string, it uses strconv.ParseBool to parse to boolean.
// For numeric type (int, float), it treats each nonzero value as true and zero as false.
// For bool type, it directly returns the original value.
// For interface and pointer type, this function will make a recursive call to get the Elem's bool value if it is not nil.
// For other non-supported types, this function will return an error indicating the type is not supported.
//
// Parameters:
// - a: The input value that needs to be converted to boolean. It can be of any type.
//
// Returns:
// - bool: The converted boolean value.
// - error: An error that will be returned if the input type is not supported for conversion.
//
// Panic:
// - This function does not introduce any intentional panic.
//
// Examples:
//
//	str := "true"
//	b1, err1 := ToBoolWithErr(str)  // b1:true, err1:nil
//
//	i := 7
//	b2, err2 := ToBoolWithErr(i)  // b2:true, err2:nil
//
//	f := 0.0
//	b3, err3 := ToBoolWithErr(f)  // b3:false, err3:nil
//
//	unsupportedType := []int{1,2,3}
//	b4, err4 := ToBoolWithErr(unsupportedType)  // b4:false, err4:error
func ToBoolWithErr(a any) (bool, error) {
	reflectValue := reflect.ValueOf(a)

	switch reflectValue.Kind() {
	case reflect.String:
		return strconv.ParseBool(reflectValue.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if reflectValue.Int() == 0 {
			return false, nil
		}
		return true, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if reflectValue.Uint() == 0 {
			return false, nil
		}
		return true, nil
	case reflect.Float32, reflect.Float64:
		if reflectValue.Float() == 0 {
			return false, nil
		}
		return true, nil
	case reflect.Complex64, reflect.Complex128:
		if reflectValue.Complex() == 0 {
			return false, nil
		}
		return true, nil
	case reflect.Bool:
		return reflectValue.Bool(), nil
	case reflect.Interface, reflect.Pointer:
		if reflectValue.IsNil() {
			return false, errors.New("error convert to bool, it is null")
		}
		return ToBoolWithErr(reflectValue.Elem().Interface())
	default:
		return false, fmt.Errorf("error convert to bool, unsupported type %s", reflectValue.Kind().String())
	}
}
