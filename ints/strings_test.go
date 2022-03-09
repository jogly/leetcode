package ints

import (
	"fmt"
	"math"
	"testing"
)

func TestChars(t *testing.T) {
	tests := []struct {
		args int
		want int
	}{
		{0, 1},
		{1, 1},
		{2, 1},
		{9, 1},
		{10, 2},
		{99, 2},
		{100, 3},
		{9999, 4},
		{math.MaxInt64, 19},
		{-1, 2},
		{-10, 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.args), func(t *testing.T) {
			if got := Chars(tt.args); got != tt.want {
				t.Errorf("Chars() = %v, want %v (%d)", got, tt.want, tt.args)
			}
		})
	}
}
