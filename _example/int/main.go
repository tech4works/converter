package main

import (
	"fmt"
	"github.com/tech4works/converter"
)

func main() {
	var str = "1234"
	var num = 1234
	var nilPointer *int

	fmt.Println("CouldBeInt result:")
	fmt.Println(converter.CouldBeInt(str))        // Should return true.
	fmt.Println(converter.CouldBeInt(num))        // Should return true.
	fmt.Println(converter.CouldBeInt(nilPointer)) // Should return false.

	fmt.Println("ToIntWithErr result:")
	i, err := converter.ToIntWithErr(str)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(i)
	}

	fmt.Println("ToInt8WithErr result:")
	i8, err := converter.ToInt8WithErr(num)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(i8)
	}

	fmt.Println("ToInt16WithErr result:")
	i16, err := converter.ToInt16WithErr(num)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(i16)
	}

	fmt.Println("ToInt result:")
	fmt.Println(converter.ToInt(str))
	fmt.Println("ToInt8 result:")
	fmt.Println(converter.ToInt8(num))
	fmt.Println("ToInt16 result:")
	fmt.Println(converter.ToInt16(num))
}
