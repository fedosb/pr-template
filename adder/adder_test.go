package adder

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		x []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no values",
			args: args{x: []int{}},
			want: 0,
		},
		{
			name: "one value",
			args: args{x: []int{1}},
			want: 1,
		},
		{
			name: "two values",
			args: args{x: []int{1, 2}},
			want: 3,
		},
		{
			name: "three values",
			args: args{x: []int{1, 2, 3}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.x...); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
