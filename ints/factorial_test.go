package ints

import (
	"fmt"
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		args int
		want int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d!", tt.args), func(t *testing.T) {
			if got := Factorial(tt.args); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
