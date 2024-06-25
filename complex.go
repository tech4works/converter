package converter

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func CanToComplex(a any) bool {
	_, err := ToComplex128WithErr(a)
	return err == nil
}

func ToComplex64(a any) complex64 {
	c64, err := ToComplex64WithErr(a)
	if err != nil {
		panic(err)
	}
	return c64
}

func ToComplex64WithErr(a any) (complex64, error) {
	c128, err := ToComplex128WithErr(a)
	return complex64(c128), err
}

func ToComplex128(a any) complex128 {
	c128, err := ToComplex128WithErr(a)
	if err != nil {
		panic(err)
	}
	return c128
}

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
