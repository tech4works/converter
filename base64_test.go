package converter

import (
	"testing"
)

type base64Case struct {
	name    string
	arg     any
	want    string
	wantErr bool
}

func TestCouldBeBase64(t *testing.T) {
	testCases := []struct {
		name string
		arg  any
		want bool
	}{
		{"Valid Base64 String", "not base64, too!!", true},
		{"Invalid Base64 String", func() {}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CouldBeBase64(tc.arg)
			if got != tc.want {
				t.Errorf("CouldBeBase64() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestToBase64(t *testing.T) {
	for _, tc := range initToBase64TestsCase() {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}
			got := ToBase64(tc.arg)
			if got != tc.want {
				t.Errorf("ToBase64() = %s, want %s", got, tc.want)
			}
		})
	}
}

func TestFromBase64(t *testing.T) {
	for _, tc := range initFromBase64TestsCase() {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}
			got := FromBase64(tc.arg)
			if string(got) != string(tc.want) {
				t.Errorf("FromBase64() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFromBase64ToString(t *testing.T) {
	for _, tc := range initFromBase64TestsCase() {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}
			got := FromBase64ToString(tc.arg)
			if got != tc.want {
				t.Errorf("FromBase64ToString() = %s, want %s", got, tc.want)
			}
		})
	}
}

func initToBase64TestsCase() []base64Case {
	return []base64Case{
		{"String to Base64", "golang", "Z29sYW5n", false},
		{"Error", func() {}, "Z29sYW5n", true},
	}
}

func initFromBase64TestsCase() []base64Case {
	return []base64Case{
		{"Base64 to String", "Z29sYW5n", "golang", false},
		{"Error", func() {}, "Z29sYW5n", true},
	}
}
