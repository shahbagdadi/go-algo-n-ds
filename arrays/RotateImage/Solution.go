package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	last := len(matrix) - 1
	for i := 0; i < len(matrix)/2; i++ {
		matrix[i], matrix[last-i] = matrix[last-i], matrix[i]
	}
	R, C := len(matrix), len(matrix[0])
	for i := 0; i < R; i++ {
		for j := i; j < C; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	fmt.Print(matrix)
}


func main() {
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}