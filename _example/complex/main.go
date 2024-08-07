package main

import (
	"fmt"
	"github.com/tech4works/converter"
)

func main() {
	x := "5.2"
	y := 10
	z := 1 + 2i
	a := "not a numeric string"

	fmt.Println(converter.CouldBeComplex(x)) // true, as "5.2" can be converted to complex128
	fmt.Println(converter.CouldBeComplex(y)) // true, as 10 can be converted to complex128
	fmt.Println(converter.CouldBeComplex(z)) // true, as it's already complex128
	fmt.Println(converter.CouldBeComplex(a)) // false, "not a numeric string" cannot be converted to complex128

	str := "5+3i"
	complexNum := converter.ToComplex64(str)
	fmt.Printf("The complex number is: %v\n", complexNum) // The complex number is: (5+3i)

	num := 10
	complexNum = converter.ToComplex64(num)
	fmt.Printf("The complex number is: %v\n", complexNum) // The complex number is: (10+0i)

	str = "5+3i"
	c128 := converter.ToComplex128(str)
	fmt.Printf("Converted value is: %v\n", c128) // Output: Converted value is: (5+3i)

	num = 10
	c128 = converter.ToComplex128(num)
	fmt.Printf("Converted value is: %v\n", c128) // Output: Converted value is: (10+0i)
}
