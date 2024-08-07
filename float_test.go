package converter

import (
	"testing"
)

type floatCase struct {
	name    string
	args    any
	want    float64
	wantErr bool
}

func TestCouldBeFloat(t *testing.T) {
	tests := []struct {
		name string
		args interface{}
		want bool
	}{
		{"int", 42, true},
		{"float32", float32(42.0), true},
		{"float64", 42.0, true},
		{"string", "42", true},
		{"bool", true, true},
		{"complex128", complex128(42 + 42i), true},
		{"unsupported type", struct{}{}, false},
		{"nil", nil, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := CouldBeFloat(tc.args); got != tc.want {
				t.Errorf("CouldBeFloat() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestToFloat32(t *testing.T) {
	for _, tc := range initFloatTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			if got := ToFloat32(tc.args); float64(got) != tc.want {
				t.Errorf("ToFloat32() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestToFloat32WithErr(t *testing.T) {
	for _, tc := range initFloatTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ToFloat32WithErr(tc.args)

			if float64(got) != tc.want {
				t.Errorf("ToFloat32WithErr() = %v, want %v", got, tc.want)
			}

			if (err != nil) != tc.wantErr {
				t.Errorf("ToFloat32WithErr() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestToFloat64(t *testing.T) {
	for _, tc := range initFloatTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			if got := ToFloat64(tc.args); got != tc.want {
				t.Errorf("ToFloat32() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestToFloat64WithErr(t *testing.T) {
	for _, tc := range initFloatTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ToFloat64WithErr(tc.args)

			if got != tc.want {
				t.Errorf("ToFloat64WithErr() = %v, want %v", got, tc.want)
			}

			if (err != nil) != tc.wantErr {
				t.Errorf("ToFloat64WithErr() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func initFloatTestCases() []floatCase {
	p := 23
	var pn *int

	return []floatCase{
		{"int", 42, 42.0, false},
		{"uint", uint(42), 42.0, false},
		{"float32", float32(42.0), 42.0, false},
		{"float64", 42.0, 42.0, false},
		{"string", "42", 42.0, false},
		{"bool", true, 1.0, false},
		{"bool", false, 0.0, false},
		{"complex128", complex128(42 + 42i), 42.0, false},
		{"unsupported type", struct{}{}, 0, true},
		{"nil", nil, 0, true},
		{"nil pointer", pn, 0, true},
		{"pointer", &p, 23.0, false},
	}
}
