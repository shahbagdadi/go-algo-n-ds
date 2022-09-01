package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	cnt := 0
	traverse(root, math.MinInt, &cnt)
	return cnt
}

func traverse(node *TreeNode, maxv int, cnt *int) {
	if node == nil {
		return
	}
	if node.Val >= maxv {
		*cnt += 1
	}
	traverse(node.Left, max(maxv, node.Val), cnt)
	traverse(node.Right, max(maxv, node.Val), cnt)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	node3 := TreeNode{3, nil, nil}
	node1 := TreeNode{1, &node3, nil}
	node11 := TreeNode{1, nil, nil}
	node5 := TreeNode{5, nil, nil}
	node4 := TreeNode{4, &node11, &node5}
	root := TreeNode{3, &node1, &node4}
	fmt.Println(goodNodes(&root))

}
