package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	traversal(root)
}

func traversal(root *TreeNode) {
	if root == nil {
		return
	}
	// fmt.Println("enter: ", root.Val)
	traversal(root.Left)
	fmt.Println("middle: ", root.Val)
	traversal(root.Right)
	// fmt.Println("leave: ", root.Val)
}
