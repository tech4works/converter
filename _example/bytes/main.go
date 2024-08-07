package main

import (
	"errors"
	"fmt"
	"github.com/tech4works/converter"
)

func main() {
	// Initialize some variables
	var myString = "Hello World"
	var myEmptyString = ""
	var myInt = 123456
	var myNil interface{} = nil

	// Use CouldBeBytes function
	fmt.Println("CouldBeBytes results:")
	fmt.Println(converter.CouldBeBytes(myString))      // Should return true
	fmt.Println(converter.CouldBeBytes(myEmptyString)) // Should return true
	fmt.Println(converter.CouldBeBytes(myInt))         // Should return true
	fmt.Println(converter.CouldBeBytes(myNil))         // Should return false

	// Use ToBytes function
	fmt.Println("ToBytes results:")
	fmt.Println(converter.ToBytes(myString))      // Should return [72 101 108 108 111 32 87 111 114 108 100]
	fmt.Println(converter.ToBytes(myEmptyString)) // Should return []
	fmt.Println(converter.ToBytes(myInt))         // Should return [49 50 51 52 53 54]

	// Use ToBytesWithErr function
	fmt.Println("ToBytesWithErr results:")
	bs, err := converter.ToBytesWithErr(myString)
	if err != nil {
		fmt.Println(errors.Unwrap(err))
	} else {
		fmt.Println(bs) // Should return [72 101 108 108 111 32 87 111 114 108 100]
	}
	bs, err = converter.ToBytesWithErr(myEmptyString)
	if err != nil {
		fmt.Println(errors.Unwrap(err))
	} else {
		fmt.Println(bs) // Should return []
	}
	bs, err = converter.ToBytesWithErr(myInt)
	if err != nil {
		fmt.Println(errors.Unwrap(err))
	} else {
		fmt.Println(bs) // Should return [49 50 51 52 53 54]
	}
	bs, err = converter.ToBytesWithErr(myNil)
	if err != nil {
		fmt.Println(errors.Unwrap(err)) // Should return "Nil cannot be converted to bytes"
	} else {
		fmt.Println(bs)
	}
}
