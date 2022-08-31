package main

import (
	"fmt"
	"testing"
)

func Test_rotate(t *testing.T) {

	type args struct {
		matrix [][]int
	}
	ip1 := args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}
	ip2 := args{[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}}

	tests := []struct {
		name string
		args args
	}{
		{"1", ip1},
		{"2", ip2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.matrix)
			fmt.Println()
		})
	}
}
