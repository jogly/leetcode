package main

import (
	"testing"
)

func Test_isNumber(t *testing.T) {
	tests := []string{
		"-0.1", "2", "0089", "+3.14", "4.", "-.9",
		"2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789",
	}
	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			if ok := isNumber(tt); !ok {
				t.Errorf("isNumber() = %v, want %v", ok, true)
			}
		})
	}
}
