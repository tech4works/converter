package converter

import (
	"testing"
)

type destCase struct {
	name  string
	a     any
	dest  any
	err   bool
	panic bool
}

type testStruct struct {
	Name string `json:"name,omitempty"`
	Sub  struct {
		Integer int `json:"integer,omitempty"`
	} `json:"sub"`
	Document string `json:"document,omitempty"`
	Status   string `json:"status,omitempty"`
}

func TestToDestWithErr(t *testing.T) {
	testCases := []destCase{
		{
			name: "String",
			a:    "value param",
			dest: ToPointer(any(2)),
			err:  false,
		},
		{
			name: "Int",
			a:    42,
			dest: ToPointer(int64(0)),
			err:  false,
		},
		{
			name: "Uint",
			a:    23,
			dest: ToPointer(uint(0)),
			err:  false,
		},
		{
			name: "Float",
			a:    42.23123,
			dest: ToPointer(0.0),
			err:  false,
		},
		{
			name: "Struct",
			a:    "{\"name\": \"John Dude\", \"sub\": {\"integer\": 23}, \"document\": \"021.093.123-23\", \"status\": \"ACTIVE\"}",
			dest: ToPointer(testStruct{}),
			err:  false,
		},
		{
			name: "Struct From Bytes",
			a:    []byte("{\"name\": \"John Dude\", \"sub\": {\"integer\": 23}, \"document\": \"021.093.123-23\", \"status\": \"ACTIVE\"}"),
			dest: ToPointer(testStruct{}),
			err:  false,
		},
		{
			name: "Interface",
			a:    []byte("{\"name\": \"John Dude\", \"sub\": {\"integer\": 23}, \"document\": \"021.093.123-23\", \"status\": \"ACTIVE\"}"),
			dest: ToPointer(any(0)),
			err:  false,
		},
		{
			name: "Map",
			a:    "{\"name\": \"John Dude\", \"sub\": {\"integer\": 23}, \"document\": \"021.093.123-23\", \"status\": \"ACTIVE\"}",
			dest: ToPointer(map[string]any{}),
			err:  false,
		},
		{
			name: "Map From Bytes",
			a:    []byte("{\"name\": \"John Dude\", \"sub\": {\"integer\": 23}, \"document\": \"021.093.123-23\", \"status\": \"ACTIVE\"}"),
			dest: ToPointer(map[string]any{}),
			err:  false,
		},
		{
			name: "Slice of Struct",
			a:    "[{\"name\": \"John Dude\", \"sub\": {\"integer\": 23}, \"document\": \"021.093.123-23\", \"status\": \"ACTIVE\"}]",
			dest: ToPointer([]testStruct{}),
			err:  false,
		},
		{
			name: "Slice of Map",
			a:    "[{\"name\": \"John Dude\", \"sub\": {\"integer\": 23}, \"document\": \"021.093.123-23\", \"status\": \"ACTIVE\"}]",
			dest: ToPointer([]map[string]any{}),
			err:  false,
		},
		{
			name: "Func",
			a:    func() {},
			dest: ToPointer(func() {}),
			err:  true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if err := ToDestWithErr(tt.a, tt.dest); (err != nil) && !tt.err {
				t.Errorf("ToDestWithErr() = %v, want = %v", err, tt.err)
				return
			}
			toString, err := ToStringWithErr(tt.dest)
			if err == nil {
				t.Logf("ToDestWithErr() = %v", toString)
			}
		})
	}
}
