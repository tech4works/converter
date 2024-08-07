package converter

import (
	"testing"
)

type uintCase struct {
	name    string
	arg     any
	want    uint
	wantErr bool
}

func TestCouldBeUint(t *testing.T) {
	cases := []struct {
		name string
		arg  any
		want bool
	}{
		{"positiveInt", 1, true},
		{"negativeInt", -1, false},
		{"positiveFloat", 1.2, true},
		{"negativeFloat", -1.2, false},
		{"nonNumString", "hello", false},
		{"numString", "1", true},
		{"boolFalse", false, true},
		{"boolTrue", true, true},
		{"nilPointer", nil, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := CouldBeUint(c.arg)
			if got != c.want {
				t.Errorf("CouldBeUint(%v) = %v, want %v", c.arg, got, c.want)
			}
		})
	}
}

func TestToUintWithErr(t *testing.T) {
	for _, c := range initUintTestCases() {
		t.Run(c.name, func(t *testing.T) {
			got, err := ToUintWithErr(c.arg)
			if (err != nil) != c.wantErr {
				t.Errorf("Error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if got != c.want {
				t.Errorf("Got %v; want %v", got, c.want)
			}
		})
	}
}

func TestToUint8WithErr(t *testing.T) {
	for _, c := range initUintTestCases() {
		t.Run(c.name, func(t *testing.T) {
			got, err := ToUint8WithErr(c.arg)
			if (err != nil) != c.wantErr {
				t.Errorf("Error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if uint(got) != c.want {
				t.Errorf("Got %v; want %v", got, c.want)
			}
		})
	}
}

func TestToUint16WithErr(t *testing.T) {
	for _, c := range initUintTestCases() {
		t.Run(c.name, func(t *testing.T) {
			got, err := ToUint16WithErr(c.arg)
			if (err != nil) != c.wantErr {
				t.Errorf("Error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if uint(got) != c.want {
				t.Errorf("Got %v; want %v", got, c.want)
			}
		})
	}
}

func TestToUint32WithErr(t *testing.T) {
	for _, c := range initUintTestCases() {
		t.Run(c.name, func(t *testing.T) {
			got, err := ToUint32WithErr(c.arg)
			if (err != nil) != c.wantErr {
				t.Errorf("Error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if uint(got) != c.want {
				t.Errorf("Got %v; want %v", got, c.want)
			}
		})
	}
}

func TestToUint64WithErr(t *testing.T) {
	for _, c := range initUintTestCases() {
		t.Run(c.name, func(t *testing.T) {
			got, err := ToUint64WithErr(c.arg)
			if (err != nil) != c.wantErr {
				t.Errorf("Error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if uint(got) != c.want {
				t.Errorf("Got %v; want %v", got, c.want)
			}
		})
	}
}

func TestToUint(t *testing.T) {
	for _, c := range initUintTestCases() {
		t.Run(c.name, func(t *testing.T) {
			if c.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}
			got := ToUint(c.arg)
			if got != c.want {
				t.Errorf("Got %v; want %v", got, c.want)
			}
		})
	}
}

func TestToUint8(t *testing.T) {
	for _, c := range initUintTestCases() {
		t.Run(c.name, func(t *testing.T) {
			if c.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}
			got := ToUint8(c.arg)
			if uint(got) != c.want {
				t.Errorf("Got %v; want %v", got, c.want)
			}
		})
	}
}

func TestToUint16(t *testing.T) {
	for _, c := range initUintTestCases() {
		t.Run(c.name, func(t *testing.T) {
			if c.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}
			got := ToUint16(c.arg)
			if uint(got) != c.want {
				t.Errorf("Got %v; want %v", got, c.want)
			}
		})
	}
}

func TestToUint32(t *testing.T) {
	for _, c := range initUintTestCases() {
		t.Run(c.name, func(t *testing.T) {
			if c.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}
			got := ToUint32(c.arg)
			if uint(got) != c.want {
				t.Errorf("Got %v; want %v", got, c.want)
			}
		})
	}
}

func TestToUint64(t *testing.T) {
	for _, c := range initUintTestCases() {
		t.Run(c.name, func(t *testing.T) {
			if c.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}
			got := ToUint64(c.arg)
			if uint(got) != c.want {
				t.Errorf("Got %v; want %v", got, c.want)
			}
		})
	}
}

func initUintTestCases() []uintCase {
	var pn *int
	p := 1
	return []uintCase{
		{"uint", uint(1), 1, false},
		{"complex", complex(2, 1), 2, false},
		{"complex", complex(-2, 1), 0, true},
		{"positiveInt", 1, 1, false},
		{"negativeInt", -1, 0, true},
		{"nonNumString", "hello", 0, true},
		{"numString", "1", 1, false},
		{"negativeString", "-1", 0, true},
		{"boolFalse", false, 0, false},
		{"boolTrue", true, 1, false},
		{"boolTrue", true, 1, false},
		{"pointer", &p, 1, false},
		{"nilPointer", pn, 0, true},
	}
}
