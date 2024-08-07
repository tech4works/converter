package converter

import "encoding/base64"

// TODO: Adicionar exemplos

// CouldBeBase64 checks if a given value can be decoded from a base64 string without causing an error.
// It utilizes the FromBase64WithErr function to attempt a conversion from base64.
// If decoding is successful, the function returns true, otherwise it returns false.
//
// Parameters:
//   - a: The value of any type that is assumed to be a base64 string and attempted to decode.
//
// Returns:
//   - bool: A boolean value indicating whether the provided value could be decoded from a base64 string.
//
// Example:
//
//	var x string = "SGVsbG8gd29ybGQ=" // Hello world in base64
//	var y string = "NotBase64"
//	fmt.Println(CouldBeBase64(x)) // true
//	fmt.Println(CouldBeBase64(y)) // false
func CouldBeBase64(a any) bool {
	_, err := ToBase64WithErr(a)
	return err == nil
}

// ToBase64 attempts to convert a given value to a base64 string representation. The function leverages
// ToBase64WithErr to handle the conversion. If the conversion fails, the function panics.
//
// The function handles various types which are supported by the ToBase64WithErr method and
// underlying methods ToBytesWithErr and ToStringWithErr.
//
// Parameters:
//   - a: The value of any type to be converted to a base64 string.
//
// Returns:
//   - string: The base64 string representation of the provided value.
//
// Panics:
//   - If the conversion fails due to an error, the function panics.
//
// Example:
//
//	str := "Hello, World!"
//	encoded := ToBase64(str)
//	fmt.Println(encoded) // SGVsbG8sIFdvcmxkIQ==
//
//	num := 1234
//	encoded = ToBase64(num)
//	fmt.Println(encoded) // MTIzNA==
func ToBase64(a any) string {
	str, err := ToBase64WithErr(a)
	if err != nil {
		panic(err)
	}
	return str
}

// ToBase64WithErr converts a given value to a base64 string representation. The function leverages
// the ToBytesWithErr function to convert the value to a byte slice format first, which is then
// encoded into a base64 string.
//
// The function handles various types which are supported by the implemented ToBytesWithErr function
// and can be used to encode data of those types into a base64 string.
//
// Parameters:
//   - a: The value of any type to be converted to a base64 string.
//
// Returns:
//   - string: The base64 encoded string representation of the provided value.
//   - error: An error is returned in case of failure to convert.
//
// If the conversion process fails at any point, the function returns an error.
// If the data is successfully converted and encoded, it returns the base64 string and nil for the error.
//
// Example:
//
//	num := 12345
//	str, err := ToBase64WithErr(num)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(str) // MTIzNDU=
//
// Example with a string:
//
//	text := "Hello world!"
//	str, err := ToBase64WithErr(text)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(str) // SGVsbG8gd29ybGQh
func ToBase64WithErr(a any) (string, error) {
	bs, err := ToBytesWithErr(a)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bs), nil
}

// FromBase64 decodes a base64-encoded string. The function first transforms any given value to a string representation,
// then attempts to decode this string from base64 to a byte array. If decoding fails, it will trigger a panic.
//
// Parameters:
//   - a: The interface value to be decoded. This value will be converted internally to a string before decoding.
//
// Returns:
//   - []byte: A byte array resulting from decoding the base64-encoded value.
//
// Panic:
//   - The function will panic if base64 decoding fails, typically due to invalid base64 input.
//
// Example:
//
//	y := "SGVsbG8gd29ybGQ="
//	decoded := FromBase64(y)
//	fmt.Println(string(decoded)) // "Hello world"
//
//	z := 10
//	// The following will cause a runtime panic
//	FromBase64(z)
func FromBase64(a any) []byte {
	bs, err := FromBase64WithErr(a)
	if err != nil {
		panic(err)
	}
	return bs
}

// FromBase64WithErr attempts to decode a base64-encoded string from any possible types of inputs.
// The function first converts the received value to a string using ToStringWithErr function.
// If the conversion is unsuccessful, an error will be returned.
// Afterwards, it attempts to decode the provided string using the base64.StdEncoding.DecodeString method.
// If the decoding is also unsuccessful, an error will be returned.
// If both operations are successful, the decoded byte array and nil error are returned.
//
// Parameters:
//   - a: Any interface value to be decoded from base64. This value is internally converted to a string before decoding.
//
// Returns:
//   - []byte: The byte array resulting from decoding the base64-encoded string. This value could be nil if any error
//     occurred.
//   - error: An error object indicating any error occurred during the work of the function. It could be nil if no
//     errors occurred.
//
// Potential errors:
//   - Failed to convert input to string
//   - Input can't be decoded from base64
//
// Example:
//
//	var x string = "SGVsbG8gd29ybGQ=" // "Hello World" encoded in base64
//	var y int = 10
//	fmt.Println(FromBase64WithErr(x)) // []byte ("Hello World"), nil
//	fmt.Println(FromBase64WithErr(y)) // nil, error ("Failed to convert")
func FromBase64WithErr(a any) ([]byte, error) {
	s, err := ToStringWithErr(a)
	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.DecodeString(s)
}

// FromBase64ToString decodes a base64-encoded value to a string. Internally, it utilizes
// the FromBase64ToStringWithErr function, which attempts the decoding process.
// If the process is successful, it returns the decoded string. If it's not, it will panic and
// stop the operation.
//
// Parameters:
//   - a: Any interface value to be decoded from base64. This value is internally converted to a string
//     before the decoding process.
//
// Returns:
//   - string: The string resulting from decoding the base64-encoded value.
//
// Panics:
//   - When the decoding operation encounters an error, the function will stop the operation
//     and cause the system to panic.
//
// Example:
//
//	var x string = "SGVsbG8gd29ybGQ=" // "Hello World" encoded in base64
//	var y int = 10
//
//
//	fmt.Println(FromBase64ToString(x)) // "Hello World"
//
//	// This will panic because int values are not supported
//	fmt.Println(FromBase64ToString(y)) // panic: error("Unsupported type for base64 decoding")
func FromBase64ToString(a any) string {
	s, err := FromBase64ToStringWithErr(a)
	if err != nil {
		panic(err)
	}
	return s
}

// FromBase64ToStringWithErr decodes a base64-encoded value into a string.
//
// Parameters:
//   - a: Any interface value that is base64-encoded. The function first attempts to convert this value into a string.
//
// Returns:
//   - string: The decoded string representation of the base64-encoded input.
//   - error: An error that occurred during the conversion or decoding process. If no issues occurred, it returns nil.
//
// Potential errors:
//   - Failed to convert any value to string type
//   - Base64 decoding error
//
// Note: This function relies on FromBase64WithErr and ToStringWithErr to perform the decoding and conversion.
//
// Example:
//
//	var b64str string = "SGVsbG8gd29ybGQ=" // "Hello World" in base64
//	var b64int int = 10 // Invalid base64 value
//
//	// This call is successful, it decodes the string and returns "Hello World" with nil error.
//	result, err := FromBase64ToStringWithErr(b64str)
//	if err != nil {
//		log.Printf("An error occurred: %v", err)
//	} else {
//		fmt.Println(result)  // "Hello World"
//	}
//
//	// This call fails and returns an empty string with error because an int value is invalid for base64 decoding.
//	_, err = FromBase64ToStringWithErr(b64int)
//	if err != nil {
//		log.Printf("An error occurred: %v", err)  // error("Unsupported type for base64 decoding")
//	}
func FromBase64ToStringWithErr(a any) (string, error) {
	bs, err := FromBase64WithErr(a)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
