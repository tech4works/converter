package converter

import (
	"testing"
)

type boolCase struct {
	name string
	arg  any
	want bool
	err  bool
}

func TestCouldBeBool(t *testing.T) {
	tests := []struct {
		name string
		arg  any
		want bool
	}{
		{
			name: "BoolTrue",
			arg:  true,
			want: true,
		},
		{
			name: "BoolFalse",
			arg:  false,
			want: true,
		},
		{
			name: "StringTrue",
			arg:  "true",
			want: true,
		},
		{
			name: "StringFalse",
			arg:  "false",
			want: true,
		},
		{
			name: "InvalidString",
			arg:  "ttrrue",
			want: false,
		},
		{
			name: "Int",
			arg:  1,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CouldBeBool(tt.arg); got != tt.want {
				t.Errorf("CouldBeBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBool(t *testing.T) {
	tests := []boolCase{
		{
			name: "BoolTrue",
			arg:  true,
			want: true,
		},
		{
			name: "BoolFalse",
			arg:  false,
			want: false,
		},
		{
			name: "StringTrue",
			arg:  "true",
			want: true,
		},
		{
			name: "StringFalse",
			arg:  "false",
			want: false,
		},
		{
			name: "InvalidString",
			arg:  "ttrrue",
			err:  true,
		},
		{
			name: "Int",
			arg:  1,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			got := ToBool(tt.arg)
			if got != tt.want {
				t.Errorf("ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBoolWithErr(t *testing.T) {
	b := true
	var p *int

	tests := []boolCase{
		{
			name: "Test boolean true",
			arg:  true,
			want: true,
			err:  false,
		},
		{
			name: "Test boolean false",
			arg:  false,
			want: false,
			err:  false,
		},
		{
			name: "Test positive integer",
			arg:  1989,
			want: true,
			err:  false,
		},
		{
			name: "Test zero integer",
			arg:  0,
			want: false,
			err:  false,
		},
		{
			name: "Test negative integer",
			arg:  -1,
			want: true,
			err:  false,
		},
		{
			name: "Test positive float",
			arg:  3.1415,
			want: true,
			err:  false,
		},
		{
			name: "Test zero float",
			arg:  0.0,
			want: false,
			err:  false,
		},
		{
			name: "Test negative float",
			arg:  -2.7182,
			want: true,
			err:  false,
		},
		{
			name: "Test string true",
			arg:  "true",
			want: true,
			err:  false,
		},
		{
			name: "Test string false",
			arg:  "false",
			want: false,
			err:  false,
		},
		{
			name: "Test string number",
			arg:  "123",
			want: false,
			err:  true,
		},
		{
			name: "Test null interface",
			arg:  nil,
			want: false,
			err:  true,
		},
		{
			name: "Test unsupported type",
			arg:  []int{1, 2, 3},
			want: false,
			err:  true,
		},
		{
			name: "Test Uint",
			arg:  uint(1),
			want: true,
		},
		{
			name: "Test Uint",
			arg:  uint(0),
			want: false,
		},
		{
			name: "Test Complex",
			arg:  complex(1, 2),
			want: true,
		},
		{
			name: "Test Complex",
			arg:  complex(0, 0),
			want: false,
		},
		{
			name: "Test Pointer",
			arg:  &b,
			want: true,
		},
		{
			name: "Test Pointer",
			arg:  p,
			err:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBoolWithErr(tt.arg)
			if (err != nil) != tt.err {
				t.Errorf("ToBoolWithErr() error = %v, wantErr %v", err, tt.err)
			} else if got != tt.want {
				t.Errorf("ToBoolWithErr() = %v, want %v", got, tt.want)
			}
		})
	}
}
