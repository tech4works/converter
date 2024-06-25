package converter

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func CanToString(a any) bool {
	_, err := ToStringWithErr(a)
	return err == nil
}

func ToString(a any) string {
	str, err := ToStringWithErr(a)
	if err != nil {
		panic(err)
	}
	return str
}

func ToStringWithErr(a any) (string, error) {
	reflectValue := reflect.ValueOf(a)

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
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Struct:
		if (reflectValue.Kind() == reflect.Array || reflectValue.Kind() == reflect.Slice) &&
			reflectValue.Type().Elem().Kind() == reflect.Uint8 {
			return string(reflectValue.Bytes()), nil
		}
		marshal, _ := json.Marshal(reflectValue.Interface())
		return string(marshal), nil
	case reflect.Ptr, reflect.Interface:
		if reflectValue.IsNil() {
			return "", errors.New("error convert to string, it is null")
		}
		return ToStringWithErr(reflectValue.Elem().Interface())
	default:
		return "", fmt.Errorf("error convert to string, unsupported type %s", reflectValue.Kind().String())
	}
}
