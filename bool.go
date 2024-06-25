package converter

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func CanToBool(a any) bool {
	_, err := ToBoolWithErr(a)
	return err == nil
}

func ToBool(a any) bool {
	b, err := ToBoolWithErr(a)
	if err != nil {
		panic(err)
	}
	return b
}

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
