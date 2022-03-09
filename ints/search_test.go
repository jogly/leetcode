package ints

import "testing"

func TestSearchAny(t *testing.T) {
	type args struct {
		a []int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			args: args{
				a: []int{},
				x: 1,
			},
			want: -1,
		},
		{
			name: "one",
			args: args{
				a: []int{1},
				x: 1,
			},
			want: 0,
		},
		{
			name: "any",
			args: args{
				a: []int{1, 2, 3, 4, 6, 6, 6, 7, 8, 9, 10},
				x: 6,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchAny(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("SearchAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchFirst(t *testing.T) {
	type args struct {
		a []int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			args: args{
				a: []int{},
				x: 1,
			},
			want: -1,
		},
		{
			name: "one",
			args: args{
				a: []int{1},
				x: 1,
			},
			want: 0,
		},
		{
			name: "any",
			args: args{
				a: []int{1, 2, 3, 4, 6, 6, 6, 7, 8, 9, 10},
				x: 6,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchFirst(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("SearchFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchLast(t *testing.T) {
	type args struct {
		a []int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			args: args{
				a: []int{},
				x: 1,
			},
			want: -1,
		},
		{
			name: "one",
			args: args{
				a: []int{1},
				x: 1,
			},
			want: 0,
		},
		{
			name: "any",
			args: args{
				a: []int{1, 2, 3, 4, 6, 6, 6, 7, 8, 9, 10},
				x: 6,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchLast(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("SearchLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchClosest(t *testing.T) {
	type args struct {
		a []int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			args: args{
				a: []int{},
				x: 1,
			},
			want: -1,
		},
		{
			name: "one",
			args: args{
				a: []int{1},
				x: 1,
			},
			want: 0,
		},
		{
			name: "any",
			args: args{
				a: []int{1, 2, 3, 4, 6, 6, 6, 7, 8, 9, 10},
				x: 5,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchClosest(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("SearchClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
