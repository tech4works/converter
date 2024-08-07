package converter

import (
	"testing"
)

func TestCouldBeComplex(t *testing.T) {
	tests := []struct {
		name string
		arg  interface{}
		want bool
	}{
		{
			name: "InvalidInterfaceType",
			arg:  []int{1, 2, 3},
			want: false,
		},
		{
			name: "ValidInterfaceType",
			arg:  "1+2i",
			want: true,
		},
		{
			name: "ValidNumType",
			arg:  5,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CouldBeComplex(tt.arg); got != tt.want {
				t.Errorf("CouldBeComplex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToComplex64(t *testing.T) {
	tests := []struct {
		name    string
		arg     interface{}
		want    complex64
		wantErr bool
	}{
		{
			name:    "InvalidInterfaceType",
			arg:     []int{1, 2, 3},
			want:    0,
			wantErr: true,
		},
		{
			name:    "ValidInterfaceType",
			arg:     "1+2i",
			want:    1 + 2i,
			wantErr: false,
		},
		{
			name:    "ValidNumType",
			arg:     5,
			want:    5 + 0i,
			wantErr: false,
		},
		{
			name:    "ValidBool",
			arg:     true,
			want:    1 + 0i,
			wantErr: false,
		},
		{
			name:    "ValidBool",
			arg:     false,
			want:    0 + 0i,
			wantErr: false,
		},
		{
			name:    "ValidUint",
			arg:     uint(1),
			want:    1 + 0i,
			wantErr: false,
		},
		{
			name:    "ValidComplex",
			arg:     complex(1, 0),
			want:    1 + 0i,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); (r != nil) != tt.wantErr {
					t.Errorf("ToComplex64() recover = %v, wantErr %v", r, tt.wantErr)
				}
			}()

			if got := ToComplex64(tt.arg); got != tt.want {
				t.Errorf("ToComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToComplex128(t *testing.T) {
	p := 1
	var pn *int

	tests := []struct {
		name    string
		arg     interface{}
		want    complex128
		wantErr bool
	}{
		{
			name:    "InvalidInterfaceType",
			arg:     []int{1, 2, 3},
			want:    0,
			wantErr: true,
		},
		{
			name:    "ValidInterfaceType",
			arg:     "1+2i",
			want:    1 + 2i,
			wantErr: false,
		},
		{
			name:    "ValidNumType",
			arg:     5,
			want:    5 + 0i,
			wantErr: false,
		},
		{
			name:    "ValidFloatType",
			arg:     5.2,
			want:    5.2 + 0i,
			wantErr: false,
		},
		{
			name:    "ValidString",
			arg:     "1+2",
			want:    0,
			wantErr: true,
		},
		{
			name: "ValidPointer",
			arg:  &p,
			want: 1 + 0i,
		},
		{
			name:    "InvalidPointer",
			arg:     pn,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); (r != nil) != tt.wantErr {
					t.Errorf("ToComplex128() recover = %v, wantErr %v", r, tt.wantErr)
				}
			}()

			if got := ToComplex128(tt.arg); got != tt.want {
				t.Errorf("ToComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}
