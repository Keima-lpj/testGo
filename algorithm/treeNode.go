package algorithm

import "fmt"

//二叉树结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历
func PrintFront(root *TreeNode) {
	if root != nil {
		fmt.Println(root.Val)
		fmt.Println("---------------------")
		PrintFront(root.Left)
		PrintFront(root.Right)
	}
}

//中遍历
func PrintMid(root *TreeNode) {
	if root != nil {
		PrintMid(root.Left)
		fmt.Println(root.Val)
		fmt.Println("---------------------")
		PrintMid(root.Right)
	}
}

//后序遍历
func PrintBack(root *TreeNode) {
	if root != nil {
		PrintBack(root.Left)
		PrintBack(root.Right)
		fmt.Println(root.Val)
		fmt.Println("---------------------")
	}
}
