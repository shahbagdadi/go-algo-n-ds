package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	q := list.New()
	var res []float64
	q.PushBack(root)
	vsum, csum, cnt := 0.0, 0, 1

	for q.Len() > 0 {
		n := q.Front().Value.(*TreeNode)
		q.Remove(q.Front())
		vsum += float64(n.Val)
		csum += 1
		cnt -= 1
		if n.Left != nil {
			q.PushBack(n.Left)
		}
		if n.Right != nil {
			q.PushBack(n.Right)
		}
		if cnt == 0 {
			res = append(res, vsum/float64(csum))
			csum, vsum = 0, 0
			cnt = q.Len()
		}

	}
	return res
}

func main() {
	node15 := TreeNode{15, nil, nil}
	node7 := TreeNode{7, nil, nil}
	node20 := TreeNode{20, &node15, &node7}
	node9 := TreeNode{9, nil, nil}
	root := TreeNode{3, &node9, &node20}
	result := averageOfLevels(&root)
	fmt.Println(result)

}
