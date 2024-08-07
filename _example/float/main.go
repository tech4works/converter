package main

import (
	"fmt"
	"github.com/tech4works/converter"
)

func main() {
	// Variables for demonstration
	var integerValue int = 10
	var floatValue float64 = 20.5
	var complexValue complex128 = 1 + 2i
	var stringValue string = "30.5"
	var boolValue bool = true
	var nilPointer *int = nil

	// Testing CouldBeFloat function
	fmt.Println("CouldBeFloat results:")
	fmt.Println(converter.CouldBeFloat(integerValue)) // Should return true.
	fmt.Println(converter.CouldBeFloat(floatValue))   // Should return true.
	fmt.Println(converter.CouldBeFloat(complexValue)) // Should return true.
	fmt.Println(converter.CouldBeFloat(stringValue))  // Should return true.
	fmt.Println(converter.CouldBeFloat(boolValue))    // Should return true.
	fmt.Println(converter.CouldBeFloat(nilPointer))   // Should return false.

	// Testing ToFloat32 Function
	fmt.Println("ToFloat32 results:")
	fmt.Println(converter.ToFloat32(integerValue)) // Should return 10.0.
	fmt.Println(converter.ToFloat32(stringValue))  // Should return 30.5.
	// Note: This will panic if an unsupported type or nil is passed.

	// Testing ToFloat32WithErr Function
	fmt.Println("ToFloat32WithErr results:")
	val, err := converter.ToFloat32WithErr(integerValue)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val) // Should return 10.0.
	}
	// Note: This will return an error if an unsupported type or nil is passed.

	// Testing ToFloat64 Function
	fmt.Println("ToFloat64 results:")
	fmt.Println(converter.ToFloat64(integerValue)) // Should return 10.0.
	fmt.Println(converter.ToFloat64(stringValue))  // Should return 30.5.
	// Note: This will panic if an unsupported type or nil is passed.

	// Testing ToFloat64WithErr Function
	fmt.Println("ToFloat64WithErr results:")
	val64, err := converter.ToFloat64WithErr(integerValue)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val64) // Should return 10.0.
	}
	// Note: This will return an error if an unsupported type or nil is passed.
}
