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

func TestToDest(t *testing.T) {
	for _, tt := range initDestTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			ToDest(tt.a, tt.dest)
		})
	}
}

func TestToDestWithErr(t *testing.T) {
	for _, tt := range initDestTestCases() {
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

func initDestTestCases() []destCase {
	return []destCase{
		{
			name: "String",
			a:    "value param",
			dest: ToPointer(any(2)),
			err:  false,
		},
		{
			name: "String",
			a:    "value param",
			dest: ToPointer(""),
			err:  false,
		},
		{
			name: "Bool",
			a:    "true",
			dest: ToPointer(false),
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
			name: "Complex",
			a:    42.23123,
			dest: ToPointer(complex(0.0, 0)),
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
		{
			name: "Nil Dest",
			a:    "{}",
			dest: (*map[string]any)(nil),
			err:  true,
		},
		{
			name: "Invalid Dest",
			a:    "{}",
			dest: map[string]any{},
			err:  true,
		},
		{
			name: "Invalid Value",
			a:    nil,
			dest: ToPointer(map[string]any{}),
			err:  true,
		},
		{
			name: "Invalid Value Func",
			a:    func() {},
			dest: ToPointer(map[string]any{}),
			err:  true,
		},
		{
			name: "Invalid Value Func",
			a:    func() {},
			dest: ToPointer(""),
			err:  true,
		},
		{
			name: "Invalid Value Bool",
			a:    func() {},
			dest: ToPointer(false),
			err:  true,
		},
		{
			name: "Invalid Value Int",
			a:    func() {},
			dest: ToPointer(1),
			err:  true,
		},
		{
			name: "Invalid Value Uint",
			a:    func() {},
			dest: ToPointer(uint(1)),
			err:  true,
		},
		{
			name: "Invalid Value Float",
			a:    func() {},
			dest: ToPointer(1.0),
			err:  true,
		},
		{
			name: "Invalid Value Complex",
			a:    func() {},
			dest: ToPointer(complex(1.0, 1)),
			err:  true,
		},
	}
}
