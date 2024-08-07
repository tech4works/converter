package main

import (
	"fmt"
	"github.com/tech4works/converter"
)

func main() {
	fmt.Println("ToBoolWithErr results:")
	val, err := converter.ToBoolWithErr("true")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val) // Should return true.
	}

	val, err = converter.ToBoolWithErr("false")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val) // Should return false.
	}
	_, err = converter.ToBoolWithErr("invalid")
	if err != nil {
		fmt.Println(err) // Should return error.
	}

	val, err = converter.ToBoolWithErr(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val) // Should return true.
	}

	val, err = converter.ToBoolWithErr(0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val) // Should return false.
	}

	fmt.Println()
	fmt.Println("ToBool results:")
	fmt.Println(converter.ToBool("true")) // Should return true.
	fmt.Println(converter.ToBool(false))  // Should return false.
	fmt.Println(converter.ToBool(1))      // Should return true.
	fmt.Println(converter.ToBool(0))      // Should return false.

	fmt.Println()
	fmt.Println("CouldBeBool results:")
	fmt.Println(converter.CouldBeBool("true"))    // Should return true.
	fmt.Println(converter.CouldBeBool("false"))   // Should return true.
	fmt.Println(converter.CouldBeBool("invalid")) // Should return false.
	fmt.Println(converter.CouldBeBool(1))         // Should return true.
	fmt.Println(converter.CouldBeBool(0))         // Should return true.
}
