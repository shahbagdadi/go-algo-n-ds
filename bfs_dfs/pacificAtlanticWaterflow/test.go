package main

import (
	"container/list"
	"fmt"
)

func main() {

	q := list.New()

	q.PushBack([]int{10, 9})
	q.PushBack([]int{9, 8})
	q.PushBack([]int{8, 7})

	for i := q.Front(); i != nil; i = i.Next() {
		fmt.Printf("value = %v", i.Value.([]int))
	}
}
