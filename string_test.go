package converter

import (
	"testing"
)

var testData = []struct {
	name        string
	input       any
	expect      string
	expectError bool
}{
	{name: "String", input: "test", expect: "test", expectError: false},
	{name: "Integer", input: 64, expect: "64", expectError: false},
	{name: "Unsigned integer", input: uint(64), expect: "64", expectError: false},
	{name: "Float", input: 64.64, expect: "64.64", expectError: false},
	{name: "Complex number", input: complex(64, 64), expect: "(64+64i)", expectError: false},
	{name: "Boolean", input: true, expect: "true", expectError: false},
	{name: "Byte Slice", input: []byte("test"), expect: "test", expectError: false},
	{name: "Map", input: map[string]interface{}{"key": "value"}, expect: `{"key":"value"}`, expectError: false},
	{name: "Nil pointer", input: (*ConvertTest)(nil), expect: "", expectError: true},
}

type ConvertTest struct {
	Value string
}

func TestToStringWithErr(t *testing.T) {
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			ret, err := ToStringWithErr(data.input)
			if err != nil && !data.expectError {
				t.Errorf("error occurred: %v", err)
			} else if err == nil && data.expectError {
				t.Error("expected error but none occurred")
			} else if ret != data.expect {
				t.Errorf("expected %q but got %q", data.expect, ret)
			}
		})
	}
}

func TestCouldBeString(t *testing.T) {
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			ret := CouldBeString(data.input)
			if ret && data.expectError {
				t.Error("expected failure but conversion was successful")
			} else if !ret && !data.expectError {
				t.Error("expected success but conversion failed")
			}
		})
	}
}

func TestToString(t *testing.T) {
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !data.expectError {
					t.Errorf("unexpected panic: %v", r)
				}
			}()

			ret := ToString(data.input)
			if ret != data.expect {
				t.Errorf("expected %q but got %q", data.expect, ret)
			}
		})
	}
}
