package binarytree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintPreOrder(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)

	PrintPreOrder(root.Left)
	PrintPreOrder(root.Right)
}
