package utils

import (
	"math/big"
	"testing"
)

func TestParseEther(t *testing.T) {

	tests := []struct {
		name   string
		amount interface{}
		want   string
	}{
		{
			name:   "type_float64",
			amount: 123456789.1234,
			want:   "123456789123400000000000000",
		},
		{
			name:   "type_string",
			amount: "123456789.1234",
			want:   "123456789123400000000000000",
		},
		{
			name:   "type_int64",
			amount: 123,
			want:   "123000000000000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseEther(tt.amount); got.String() != tt.want {
				t.Errorf("FormatEther() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatEther(t *testing.T) {

	bigAmount := new(big.Int)
	bigAmount, _ = bigAmount.SetString("123456789123400000000000000", 10)

	tests := []struct {
		name   string
		amount interface{}
		want   string
	}{
		{
			name:   "type_string",
			amount: "123456789123400000000000000",
			want:   "123456789.1234",
		},
		{
			name:   "type_big_int",
			amount: bigAmount,
			want:   "123456789.1234",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatEther(tt.amount); got.String() != tt.want {
				t.Errorf("FormatEther() = %v, want %v", got, tt.want)
			}
		})
	}
}
