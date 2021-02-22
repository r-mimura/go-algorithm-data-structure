package mylib

import (
	"testing"
)

func Test_min(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "normal case", args: args{nums: []int{1, 1e9}}, want: 1},
		{name: "multiple inputs case", args: args{nums: []int{1, 1e9, 1e5, -1e5, -1e9, 0}}, want: -1e9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.nums...); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "normal case", args: args{nums: []int{1, 1e9}}, want: 1e9},
		{name: "multiple inputs case", args: args{nums: []int{1, 1e18, 1e5, -1e5, -1e9, 0}}, want: 1e18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.nums...); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gcd(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "normal case", args: args{m: 51, n: 15}, want: 3},
		{name: "normal case2", args: args{m: 15, n: 51}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}
