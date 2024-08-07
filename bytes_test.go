package converter

import (
	"errors"
	"testing"
)

func TestCouldBeBytes(t *testing.T) {
	tests := []struct {
		name string
		arg  any
		want bool
	}{
		{"ByteString", "Hello", true},
		{"EmptyString", "", true},
		{"Number", 123, true},
		{"Nil", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CouldBeBytes(tt.arg); got != tt.want {
				t.Errorf("CouldBeBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBytes(t *testing.T) {
	tests := []struct {
		name        string
		arg         any
		want        []byte
		expectPanic bool
	}{
		{"ByteString", "Hello", []byte("Hello"), false},
		{"EmptyString", "", []byte(""), false},
		{"Number", 123, nil, true},
		{"Nil", nil, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !tt.expectPanic {
					t.Errorf("ToBytes panic = %v, want panic = %v", r, tt.expectPanic)
				}
			}()

			if got := ToBytes(tt.arg); !tt.expectPanic && string(got) != string(tt.want) {
				t.Errorf("ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBytesWithErr(t *testing.T) {
	tests := []struct {
		name    string
		arg     any
		want    []byte
		wantErr error
	}{
		{"ByteString", "Hello", []byte("Hello"), nil},
		{"EmptyString", "", []byte(""), nil},
		{"Number", 123, []byte("123"), nil},
		{"Nil", nil, nil, errors.New("Nil cannot be converted to string")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBytesWithErr(tt.arg)

			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("ToBytesWithErr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if string(got) != string(tt.want) {
				t.Errorf("ToBytesWithErr() = %v, want %v", got, tt.want)
			}
		})
	}
}
