package main

import (
	"fmt"
	"github.com/tech4works/converter"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	var sampleNumber = 10
	var sampleChannel = make(chan int)
	var sampleString = "Hello, World!"
	var sampleStudent = Student{Name: "John", Age: 20}
	var sampleInterface interface{} = 15

	// Test CouldBeString function
	fmt.Println("CouldBeString Results:")
	fmt.Println(converter.CouldBeString(sampleNumber))  // Should return true
	fmt.Println(converter.CouldBeString(sampleChannel)) // Should return false

	// Test ToString function
	fmt.Println("ToString Results:")
	fmt.Println(converter.ToString(sampleNumber)) // Should return "10"
	fmt.Println(converter.ToString(sampleString)) // Should return "Hello, World!"

	// Test ToStringWithErr function
	fmt.Println("ToStringWithErr Results:")
	str, err := converter.ToStringWithErr(sampleNumber)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str) // Should print "10"
	}

	str, err = converter.ToStringWithErr(sampleString)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str) // Should print "Hello, World!"
	}

	str, err = converter.ToStringWithErr(sampleStudent)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str) // Should print `{"Name":"John","Age":20}`
	}

	str, err = converter.ToStringWithErr(sampleInterface)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str) // Should print "15"
	}
}
