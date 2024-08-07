package converter

// CouldBeBytes checks whether a given value could be converted to bytes.
// It attempts to convert a value to bytes using the ToBytesWithErr function,
// checking if an error is nil (indicating a successful conversion), and returns this result.
//
// Parameters:
//   - a: A value of any type that is to be checked for possible byte conversion.
//
// Returns:
//   - bool: A boolean value indicating whether the given value can be converted to bytes.
//
// Example:
//
//	name := "John Doe"
//	num := 1234
//	fmt.Println(CouldBeBytes(name)) // true
//	fmt.Println(CouldBeBytes(num))  // true
func CouldBeBytes(a any) bool {
	_, err := ToBytesWithErr(a)
	return err == nil
}

// ToBytes converts any given value to bytes. It leverages the ToBytesWithErr() function underneath
// and if any error occurs during the byte conversion, the function panics.
//
// Parameters:
//   - a: Any interface{} type value which needs to be converted into bytes.
//
// Returns:
//   - []byte: A byte slice representing the input value.
//
// Panic:
//   - If any error occurs during the byte conversion, the function will panic.
//
// Example:
//
//	var x int = 10
//	fmt.Println(ToBytes(x)) // [49 48]
//	y := "Hello"
//	fmt.Println(ToBytes(y)) // [72 101 108 108 111]
//
// Please note that the output of `ToBytes` will be a slice representing each character of the input data in ASCII.
func ToBytes(a any) []byte {
	bs, err := ToBytesWithErr(a)
	if err != nil {
		panic(err)
	}
	return bs
}

// ToBytesWithErr attempts to convert a given value to a byte slice representation based on its type.
// The function leverages ToStringWithErr function to convert the value to a string format first, which is
// then cast into a byte slice. If the conversion fails at any point, the function returns an error.
//
// The function handles various types which are supported by the ToStringWithErr method:
//   - String, Integer, Unsigned Integer, Float
//   - Complex, Boolean, Array, Map, Struct, Nil, etc.
//
// Parameters:
//   - a: The value of any type to be converted to a byte slice.
//
// Returns:
//   - []byte: The byte slice representation of the provided value.
//   - error: An error is returned in case of failure to convert.
//
// Example:
//
//	num := 1234
//	result, err := ToBytesWithErr(num)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(result) // [49 50 51 52]
//
// The output for the `ToBytesWithErr` function is a byte slice representing each character of the string converted
// data in ASCII.
func ToBytesWithErr(a any) ([]byte, error) {
	s, err := ToStringWithErr(a)
	return []byte(s), err
}
