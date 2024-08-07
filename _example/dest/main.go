package main

import (
	"fmt"
	"github.com/tech4works/converter"
)

type Student struct {
	Name    string
	Age     int
	Color   string `json:"Color"`
	IsEmpty bool   `json:"empty"`
}

func main() {
	// ToDest example with Struct
	mapValue := map[string]interface{}{
		"name":  "John",
		"age":   32,
		"Color": "Blue",
	}

	var student Student
	converter.ToDest(mapValue, &student)

	fmt.Printf("ToDest Struct: %+v\n", student)

	// ToDestWithErr example with Struct
	mapValueErr := map[string]interface{}{
		"name":  "Adam",
		"Color": "Green",
	}

	var studentErr Student
	err := converter.ToDestWithErr(mapValueErr, &studentErr)

	if err != nil {
		fmt.Printf("ToDestWithErr Struct Failure: %s\n", err)
		return
	}

	fmt.Printf("ToDestWithErr Struct: %+v\n", studentErr)

	// ToDest example with Integer
	var integerValue int
	converter.ToDest("123", &integerValue)

	fmt.Printf("ToDest Integer: %d\n", integerValue)

	// ToDestWithErr example with Integer
	var integerValueErr int
	err = converter.ToDestWithErr("ABC", &integerValueErr)

	if err != nil {
		fmt.Printf("ToDestWithErr Integer Failure: %s\n", err)
		return
	}

	fmt.Printf("ToDestWithErr Integer: %d\n", integerValueErr)

	// ToDest example with Boolean
	var boolValue bool
	converter.ToDest("false", &boolValue)

	fmt.Printf("ToDest Boolean: %t\n", boolValue)

	// ToDestWithErr example with Boolean
	var boolValueErr bool
	err = converter.ToDestWithErr("notbool", &boolValueErr)

	if err != nil {
		fmt.Printf("ToDestWithErr Boolean Failure: %s\n", err)
		return
	}

	fmt.Printf("ToDestWithErr Boolean: %t\n", boolValueErr)
}
