package converter

func ToPointer[T any](a T) *T {
	return &a
}
