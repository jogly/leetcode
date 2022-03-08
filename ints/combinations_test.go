package ints

import (
	"fmt"
	"testing"
)

func TestNCR(t *testing.T) {
	type args struct {
		n int
		r int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{0, 0}, 1},
		{args{1, 0}, 1},
		{args{1, 1}, 1},
		{args{2, 0}, 1},
		{args{2, 1}, 2},
		{args{2, 2}, 1},
		{args{3, 0}, 1},
		{args{3, 1}, 3},
		{args{3, 2}, 3},
		{args{3, 3}, 1},
		{args{4, 0}, 1},
		{args{4, 1}, 4},
		{args{4, 2}, 6},
		{args{4, 3}, 4},
		{args{4, 4}, 1},
		{args{5, 0}, 1},
		{args{5, 1}, 5},
		{args{5, 2}, 10},
		{args{5, 3}, 10},
		{args{5, 4}, 5},
		{args{5, 5}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%dC%d", tt.args.n, tt.args.r), func(t *testing.T) {
			if got := NCR(tt.args.n, tt.args.r); got != tt.want {
				t.Errorf("NCR() = %v, want %v", got, tt.want)
			}
		})
	}
}
