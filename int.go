package converter

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func CanToInt(a any) bool {
	_, err := ToIntWithErr(a)
	return err == nil
}

func ToInt(a any) int {
	i, err := ToIntWithErr(a)
	if err != nil {
		panic(err)
	}
	return i
}

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

func ToInt8(a any) int8 {
	i, err := ToInt8WithErr(a)
	if err != nil {
		panic(err)
	}
	return i
}

func ToInt8WithErr(a any) (int8, error) {
	i, err := ToIntWithErr(a)
	return int8(i), err
}

func ToInt16(a any) int16 {
	i, err := ToInt16WithErr(a)
	if err != nil {
		panic(err)
	}
	return i
}

func ToInt16WithErr(a any) (int16, error) {
	i, err := ToIntWithErr(a)
	return int16(i), err
}

func ToInt32(a any) int32 {
	i, err := ToInt32WithErr(a)
	if err != nil {
		panic(err)
	}
	return i
}

func ToInt32WithErr(a any) (int32, error) {
	i, err := ToIntWithErr(a)
	return int32(i), err
}

func ToInt64(a any) int64 {
	i, err := ToInt64WithErr(a)
	if err != nil {
		panic(err)
	}
	return i
}

func ToInt64WithErr(a any) (int64, error) {
	i, err := ToIntWithErr(a)
	return int64(i), err
}
