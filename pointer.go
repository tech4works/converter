package converter

// ToPointer creates a pointer to the given value. This is useful for creating a pointer of a variable without having to
// create a new variable.
//
// Parameters:
//   - T: is a type parameter that can be any type. The type of T will define the type of value this function accepts.
//   - a: is a value of type T for which the pointer will be created.
//
// Returns:
//   - *T: A pointer to a value of type T.
//
// Example:
//
//	intVar := 10
//	strVar := "Hello World"
//	fmt.Println(ToPointer(intVar)) // Outputs: &10
//	fmt.Println(ToPointer(strVar)) // Outputs: &Hello World
func ToPointer[T any](a T) *T {
	return &a
}
