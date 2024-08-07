package converter

import "bytes"

// TODO: Adicionar exemplos

// ToBuffer attempts to convert a given value into a buffer. It internally uses ToBufferWithErr function that
// converts the given value to a byte slice first.
// If the conversion fails, it panics and stops the execution.
//
// Parameters:
//   - a: The value of any type to be converted to a buffer.
//
// Returns:
//   - *bytes.Buffer: The buffer representation of the passed value.
//
// Example:
//
//	myValue := "This is test"
//	buffer := ToBuffer(myValue)
//	fmt.Println(buffer) // This is test
//
//	myValue := 1234
//	buffer := ToBuffer(myValue)
//	fmt.Println(buffer) // 1234
//
// Panics:
//   - If the conversion to a buffer fails at any point, it will panic.
func ToBuffer(a any) *bytes.Buffer {
	buf, err := ToBufferWithErr(a)
	if err != nil {
		panic(err)
	}
	return buf
}

// ToBufferWithErr attempts to convert a given value to a bytes.Buffer
// It first tries to convert the supplied argument to a byte slice by
// calling the function ToBytesWithErr. If a conversion error occurs,
// it returns nil along with the error.
//
// Parameters:
//   - a: The value of any type to be converted to a bytes.Buffer
//
// Returns:
//   - *bytes.Buffer: The bytes.Buffer representation of the provided value.
//   - error: An error is returned in case of failure to convert.
//
// Example:
//
//	myValue := "This is a test"
//	buffer, err := ToBufferWithErr(myValue)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(buffer) // This is a test
//
//	num := 1234
//	buffer, err := ToBufferWithErr(num)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(buffer) // 1234
func ToBufferWithErr(a any) (*bytes.Buffer, error) {
	bs, err := ToBytesWithErr(a)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(bs), nil
}
