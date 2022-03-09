package ints

import "testing"

func TestMin(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal",
			args: args{0, 0},
			want: 0,
		},
		{
			name: "a < b",
			args: args{0, 1},
			want: 0,
		},
		{
			name: "a > b",
			args: args{1, 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal",
			args: args{0, 0},
			want: 0,
		},
		{
			name: "a < b",
			args: args{0, 1},
			want: 1,
		},
		{
			name: "a > b",
			args: args{1, 0},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
