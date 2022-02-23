package ints

import (
	"reflect"
	"testing"
)

func TestDedupeInPlace(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty",
			args: args{
				a: []int{},
			},
			want: []int{},
		},
		{
			name: "one",
			args: args{
				a: []int{0},
			},
			want: []int{0},
		},
		{
			name: "duplicates",
			args: args{
				a: []int{0, 1, 1, 2},
			},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DedupeInPlace(tt.args.a)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DedupeInPlace() = %v, want %v", got, tt.want)
			}
			if len(tt.args.a) > 0 && &got[0] != &tt.args.a[0] {
				t.Error("DedupeInPlace() modified underlying memory", &got, &tt.args.a)
			}
		})
	}
}
