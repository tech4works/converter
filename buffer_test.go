package converter

import (
	"bytes"
	"errors"
	"testing"
)

func TestToBuffer(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name      string
		args      args
		want      *bytes.Buffer
		wantPanic bool
	}{
		{
			name: "ValidInputCase",
			args: args{
				a: "test string",
			},
			want:      bytes.NewBufferString("test string"),
			wantPanic: false,
		},
		{
			name: "InvalidNilInputCase",
			args: args{
				a: nil,
			},
			want:      nil,
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); (r != nil) != tt.wantPanic {
					t.Errorf("ToBuffer() recover() = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := ToBuffer(tt.args.a); !bytes.Equal(got.Bytes(), tt.want.Bytes()) {
				t.Errorf("ToBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBufferWithErr(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name    string
		args    args
		want    *bytes.Buffer
		wantErr error
	}{
		{
			name: "ValidInputCase",
			args: args{
				a: "Another test string",
			},
			want:    bytes.NewBufferString("Another test string"),
			wantErr: nil,
		},
		{
			name: "InvalidNilInputCase",
			args: args{
				a: nil,
			},
			want:    nil,
			wantErr: errors.New("Invalid input: nil"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBufferWithErr(tt.args.a)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("ToBufferWithErr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && tt.want != nil && !bytes.Equal(got.Bytes(), tt.want.Bytes()) {
				t.Errorf("ToBufferWithErr() = %v, want %v", got, tt.want)
			}
		})
	}
}
