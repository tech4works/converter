package converter

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func CanToFloat(a any) bool {
	_, err := ToFloat64WithErr(a)
	return err == nil
}

func ToFloat32(a any) float32 {
	f32, err := ToFloat32WithErr(a)
	if err != nil {
		panic(err)
	}
	return f32
}

func ToFloat32WithErr(a any) (float32, error) {
	f64, err := ToFloat64WithErr(a)
	return float32(f64), err
}

func ToFloat64(a any) float64 {
	f64, err := ToFloat64WithErr(a)
	if err != nil {
		panic(err)
	}
	return f64
}

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
	case reflect.Interface, reflect.Pointer:
		if reflectValue.IsNil() {
			return 0, errors.New("error convert to float64, it is null")
		}
		return ToFloat64WithErr(reflectValue.Elem().Interface())
	default:
		return 0, fmt.Errorf("error convert to float64, unsupported type %s", reflectValue.Kind().String())
	}
}
