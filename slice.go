package converter

func ToSlice[T any](v ...T) []T {
	return v
}

func ToSliceIfNonNil[T any](v ...T) []T {
	var ts []T
	for _, t := range v {
		if t != nil {
			ts = append(ts, t)
		}
	}
	return ts
}
