package converter

import (
	"testing"
)

type intCase struct {
	name    string
	input   any
	want    int
	wantErr bool
}

func TestCouldBeInt(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  bool
	}{
		{name: "strNum", input: "123", want: true},
		{name: "strNotNum", input: "notNum", want: false},
		{name: "int", input: 8, want: true},
		{name: "float", input: 7.0, want: true},
		{name: "bool", input: true, want: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CouldBeInt(tc.input)
			if got != tc.want {
				t.Errorf("Got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestToIntWithErr(t *testing.T) {
	for _, tc := range initIntTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ToIntWithErr(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("Error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("Got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestToInt8WithErr(t *testing.T) {
	for _, tc := range initIntTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ToInt8WithErr(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("Error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if int(got) != tc.want {
				t.Errorf("Got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestToInt16WithErr(t *testing.T) {
	for _, tc := range initIntTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ToInt16WithErr(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("Error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if int(got) != tc.want {
				t.Errorf("Got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestToInt32WithErr(t *testing.T) {
	for _, tc := range initIntTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ToInt32WithErr(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("Error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if int(got) != tc.want {
				t.Errorf("Got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestToInt64WithErr(t *testing.T) {
	for _, tc := range initIntTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ToInt64WithErr(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("Error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if int(got) != tc.want {
				t.Errorf("Got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	for _, tc := range initIntTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}
			got := ToInt(tc.input)
			if got != tc.want {
				t.Errorf("Got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestToInt8(t *testing.T) {
	for _, tc := range initIntTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}
			got := ToInt8(tc.input)
			if int(got) != tc.want {
				t.Errorf("Got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestToInt16(t *testing.T) {
	for _, tc := range initIntTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}
			got := ToInt16(tc.input)
			if int(got) != tc.want {
				t.Errorf("Got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestToInt32(t *testing.T) {
	for _, tc := range initIntTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}
			got := ToInt32(tc.input)
			if int(got) != tc.want {
				t.Errorf("Got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	for _, tc := range initIntTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}
			got := ToInt64(tc.input)
			if int(got) != tc.want {
				t.Errorf("Got %v; want %v", got, tc.want)
			}
		})
	}
}

func initIntTestCases() []intCase {
	var pn *int
	p := 1
	return []intCase{
		{name: "strNum", input: "123", want: 123, wantErr: false},
		{name: "strNotNum", input: "notNum", want: 0, wantErr: true},
		{name: "int", input: 8, want: 8, wantErr: false},
		{name: "float", input: 7.0, want: 7, wantErr: false},
		{name: "uint", input: uint(31), want: 31, wantErr: false},
		{name: "complex", input: complex(12, 1), want: 12, wantErr: false},
		{name: "pointer", input: &p, want: 1, wantErr: false},
		{name: "pointerNil", input: pn, wantErr: true},
		{name: "boolTrue", input: true, want: 1, wantErr: false},
		{name: "boolFalse", input: false, want: 0, wantErr: false},
	}
}
