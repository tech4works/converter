package converter

import (
	"fmt"
	"reflect"
	"time"
)

func ToTimeWithErr(a any) (time.Time, error) {
	reflectValue := reflect.ValueOf(a)
	switch reflectValue.Kind() {
	case reflect.String:
		layouts := []string{time.Layout, time.ANSIC, time.UnixDate, time.RubyDate, time.RFC822, time.RFC822Z,
			time.RFC850, time.RFC1123, time.RFC1123Z, time.RFC3339, time.RFC3339Nano, time.Kitchen, time.Stamp,
			time.DateTime, time.DateOnly, time.TimeOnly}
		for _, layout := range layouts {
			if t, err := time.Parse(layout, reflectValue.String()); err == nil {
				return t, nil
			}
		}
		return time.Time{}, fmt.Errorf("cannot convert string to time.Time: Unknown format \"%s\"",
			reflectValue.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return time.UnixMilli(reflectValue.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return time.UnixMilli(int64(reflectValue.Uint())), nil
	case reflect.Float32, reflect.Float64:
		return time.UnixMilli(int64(reflectValue.Float())), nil
	default:
		if reflectValue.Type() == reflect.TypeOf(time.Time{}) {
			return reflectValue.Interface().(time.Time), nil
		}
		return time.Time{}, fmt.Errorf("cannot convert to time.Time from type: %s", reflectValue.Kind().String())
	}
}

func ToTime(a any) time.Time {
	t, err := ToTimeWithErr(a)
	if err != nil {
		panic(err)
	}
	return t
}

func ToDate(a any) time.Time {
	d, err := ToDateWithErr(a)
	if err != nil {
		panic(err)
	}
	return d
}

func ToDateWithErr(a any) (time.Time, error) {
	t, err := ToTimeWithErr(a)
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()), nil
}
