package converter

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func ToDest(a, dest any) {
	err := ToDestWithErr(a, dest)
	if err != nil {
		panic(err)
	}
}

func ToDestWithErr(a, dest any) error {
	reflectValue := reflect.ValueOf(a)
	reflectDest := reflect.ValueOf(dest)

	if reflectDest.Kind() != reflect.Ptr {
		return errors.New("dest is not a pointer")
	} else if reflectDest.IsNil() {
		return errors.New("dest is nil")
	} else if !reflectValue.IsValid() ||
		(reflectValue.Kind() == reflect.Ptr || reflectValue.Kind() == reflect.Interface) && reflectValue.IsNil() {
		return errors.New("A is nil or invalid")
	}

	switch reflectDest.Elem().Kind() {
	case reflect.Struct, reflect.Map, reflect.Array, reflect.Slice:
		bs, err := ToBytesWithErr(a)
		if err != nil {
			return err
		}
		return json.Unmarshal(bs, dest)
	case reflect.String:
		s, err := ToStringWithErr(a)
		if err != nil {
			return err
		}
		reflectDest.Elem().Set(reflect.ValueOf(s))
	case reflect.Bool:
		b, err := ToBoolWithErr(a)
		if err != nil {
			return err
		}
		reflectDest.Elem().Set(reflect.ValueOf(b))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := ToIntWithErr(a)
		if err != nil {
			return err
		}
		reflectDest.Elem().Set(reflect.ValueOf(i).Convert(reflectDest.Elem().Type()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		ui, err := ToUintWithErr(a)
		if err != nil {
			return err
		}
		reflectDest.Elem().Set(reflect.ValueOf(ui).Convert(reflectDest.Elem().Type()))
	case reflect.Float32, reflect.Float64:
		f, err := ToFloat64WithErr(a)
		if err != nil {
			return err
		}
		reflectDest.Elem().Set(reflect.ValueOf(f).Convert(reflectDest.Elem().Type()))
	case reflect.Complex64, reflect.Complex128:
		c, err := ToComplex128WithErr(a)
		if err != nil {
			return err
		}
		reflectDest.Elem().Set(reflect.ValueOf(c).Convert(reflectDest.Elem().Type()))
	case reflect.Interface:
		reflectDest.Elem().Set(reflectValue)
	default:
		return fmt.Errorf("error convert to dest, unsupported type %s", reflectValue.Kind().String())
	}

	return nil
}
