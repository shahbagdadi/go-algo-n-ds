package main

import (
	"container/list"
)

func pacificAtlantic(heights [][]int) [][]int {
	pacific, atlantic := map[[2]int]struct{}{}, map[[2]int]struct{}{}
	R, C := len(heights), len(heights[0])
	for i := 0; i < R; i++ {
		pacific[[2]int{i, 0}] = struct{}{}
		atlantic[[2]int{i, C - 1}] = struct{}{}
	}
	for j := 0; j < C; j++ {
		pacific[[2]int{0, j}] = struct{}{}
		atlantic[[2]int{R - 1, j}] = struct{}{}
	}
	bfs(heights, pacific, R, C)
	bfs(heights, atlantic, R, C)
	result := [][]int{}
	for k := range atlantic {
		if _, ok := pacific[k]; ok {
			result = append(result, []int{k[0], k[1]})
		}

	}
	return result
}

func bfs(heights [][]int, ocean map[[2]int]struct{}, R int, C int) {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	q := list.New()
	for k := range ocean {
		data := make([]int, len(k))
		copy(data, k[:])
		q.PushFront(data)
	}
	for q.Len() > 0 {
		e := q.Front()
		q.Remove(e)
		n := e.Value.([]int)
		for _, dir := range dirs {
			x := n[0] + dir[0]
			y := n[1] + dir[1]
			_, exists := ocean[[2]int{x, y}]
			if 0 <= x && x < R && 0 <= y && y < C && heights[x][y] >= heights[n[0]][n[1]] && !exists {
				q.PushFront([]int{x, y})
				ocean[[2]int{x, y}] = struct{}{}
			}
		}
		n = nil
	}
}
