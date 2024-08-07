package main

import (
	"fmt"
	"github.com/tech4works/converter"
)

func main() {
	var nilPointer *int

	fmt.Println("CouldBeUint results:")
	fmt.Println(converter.CouldBeUint(nil))        // Should return false.
	fmt.Println(converter.CouldBeUint(3))          // Should return true.
	fmt.Println(converter.CouldBeUint("1234"))     // Should return true.
	fmt.Println(converter.CouldBeUint(nilPointer)) // Should return false.
	fmt.Println(converter.CouldBeUint("test"))     // Should return false.

	fmt.Println("ToUint results:")
	fmt.Println(converter.ToUint(3))      // Should return 3.
	fmt.Println(converter.ToUint("1234")) // Should return 1234.

	fmt.Println("ToUintWithErr results:")
	fmt.Println(converter.ToUintWithErr(3))          // Should return (3, nil).
	fmt.Println(converter.ToUintWithErr("1234"))     // Should return (1234, nil).
	fmt.Println(converter.ToUintWithErr("test"))     // Should return (0, error).
	fmt.Println(converter.ToUintWithErr(nilPointer)) // Should return (0, error).

	fmt.Println("ToUint8 results:")
	fmt.Println(converter.ToUint8(3)) // Should return 3.

	fmt.Println("ToUint8WithErr results:")
	fmt.Println(converter.ToUint8WithErr(3))          // Should return (3, nil).
	fmt.Println(converter.ToUint8WithErr("1234"))     // Should return (194, nil) (overflow).
	fmt.Println(converter.ToUint8WithErr("test"))     // Should return (0, error).
	fmt.Println(converter.ToUint8WithErr(nilPointer)) // Should return (0, error).

	fmt.Println("ToUint16 results:")
	fmt.Println(converter.ToUint16(3))      // Should return 3.
	fmt.Println(converter.ToUint16("1234")) // Should return 1234.

	fmt.Println("ToUint16WithErr results:")
	fmt.Println(converter.ToUint16WithErr(3))          // Should return (3, nil).
	fmt.Println(converter.ToUint16WithErr("1234"))     // Should return (1234, nil).
	fmt.Println(converter.ToUint16WithErr("test"))     // Should return (0, error).
	fmt.Println(converter.ToUint16WithErr(nilPointer)) // Should return (0, error).

	fmt.Println("ToUint32WithErr results:")
	fmt.Println(converter.ToUint32WithErr(3))          // Should return (3, nil).
	fmt.Println(converter.ToUint32WithErr("1234"))     // Should return (1234, nil).
	fmt.Println(converter.ToUint32WithErr("test"))     // Should return (0, error).
	fmt.Println(converter.ToUint32WithErr(nilPointer)) // Should return (0, error).

	fmt.Println("ToUint64WithErr results:")
	fmt.Println(converter.ToUint64WithErr(3))          // Should return (3, nil).
	fmt.Println(converter.ToUint64WithErr("1234"))     // Should return (1234, nil).
	fmt.Println(converter.ToUint64WithErr("test"))     // Should return (0, error).
	fmt.Println(converter.ToUint64WithErr(nilPointer)) // Should return (0, error).
}
