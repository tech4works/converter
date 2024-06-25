package converter

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func CanToUint(a any) bool {
	_, err := ToUintWithErr(a)
	return err == nil
}

func ToUint(a any) uint {
	u, err := ToUintWithErr(a)
	if err != nil {
		panic(err)
	}
	return u
}

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
	case reflect.Interface, reflect.Pointer:
		if reflectValue.IsNil() {
			return 0, errors.New("error convert to uint, it is null")
		}
		return ToUintWithErr(reflectValue.Elem().Interface())
	default:
		return 0, fmt.Errorf("error convert to uint, unsupported type %s", reflectValue.Kind().String())
	}
}

func ToUint8(a any) uint8 {
	u, err := ToUint8WithErr(a)
	if err != nil {
		panic(err)
	}
	return u
}

func ToUint8WithErr(a any) (uint8, error) {
	u, err := ToUintWithErr(a)
	return uint8(u), err
}

func ToUint16(a any) uint16 {
	u, err := ToUint16WithErr(a)
	if err != nil {
		panic(err)
	}
	return u
}

func ToUint16WithErr(a any) (uint16, error) {
	u, err := ToUintWithErr(a)
	return uint16(u), err
}

func ToUint32(a any) uint32 {
	u, err := ToUint32WithErr(a)
	if err != nil {
		panic(err)
	}
	return u
}

func ToUint32WithErr(a any) (uint32, error) {
	u, err := ToUintWithErr(a)
	return uint32(u), err
}

func ToUint64(a any) uint64 {
	u, err := ToUint64WithErr(a)
	if err != nil {
		panic(err)
	}
	return u
}

func ToUint64WithErr(a any) (uint64, error) {
	u, err := ToUintWithErr(a)
	return uint64(u), err
}
