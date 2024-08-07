package converter

import (
	"testing"
)

func TestToPointer(t *testing.T) {
	type testCase struct {
		name string
		val  interface{}
	}

	testCases := []testCase{
		{
			name: "WithInteger",
			val:  5,
		},
		{
			name: "WithString",
			val:  "5",
		},
		{
			name: "WithFloat",
			val:  5.5,
		},
		{
			name: "WithBool",
			val:  true,
		},
		{
			name: "WithNil",
			val:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := ToPointer(tc.val)
			want := &tc.val
			if *got != *want {
				t.Errorf("ToPointer(%v) = %v; want %v", tc.val, *got, *want)
			}
		})
	}
}
