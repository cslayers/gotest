package main

import (
	"testing"
)

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case", args{1, 0}, 1},
		{"case", args{0, 1}, 1},
		{"case", args{1, 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_substract(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case", args{1, 0}, 1},
		{"case", args{0, 1}, -1},
		{"case", args{1, -2}, 3},
		{"case", args{1, 2}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := substract(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("substract() = %v, want %v", got, tt.want)
			}
		})
	}
}