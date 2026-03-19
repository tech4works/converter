package converter

import "reflect"

func ToSlice[T any](v ...T) []T {
	return v
}

func ToSliceIfNonNil[T any](v ...T) []T {
	var ts []T
	for _, t := range v {
		if isNonNil(t) {
			ts = append(ts, t)
		}
	}
	return ts
}

func isNonNil(i any) bool {
	if i == nil {
		return true
	}
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Func, reflect.Chan:
		return !val.IsNil()
	default:
		return true
	}
}
