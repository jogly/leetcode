package ints

import (
	"bytes"
	"strings"
	"testing"
)

func TestFprintCarets(t *testing.T) {
	type args struct {
		arr    []int
		carets []int
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "empty",
			args: args{
				arr:    []int{},
				carets: []int{},
			},
			wantW: "",
		},
		{
			name: "one",
			args: args{
				arr:    []int{1},
				carets: []int{0},
			},
			wantW: "1\n^",
		},
		{
			name: "two",
			args: args{
				arr:    []int{1, 2},
				carets: []int{1},
			},
			wantW: "1 2\n  ^",
		},
		{
			name: "mixed space",
			args: args{
				arr:    []int{1, 20, 300},
				carets: []int{1},
			},
			wantW: strings.Trim(`
  1  20 300
      ^    
`, "\n"),
		},
		{
			name: "multiple carets",
			args: args{
				arr:    []int{-100, 20, 30},
				carets: []int{0, 1},
			},
			wantW: strings.Trim(`
-100   20   30
   ^    ^     
`, "\n"),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			FprintCarets(w, tt.args.arr, tt.args.carets...)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Got:\n%s\nWant:\n%s", gotW, tt.wantW)
			}
		})
	}
}
