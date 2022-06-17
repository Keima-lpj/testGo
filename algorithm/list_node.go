package algorithm

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//打印链表的方法
func printList(res *ListNode) {
	a := res
	fmt.Print("res:")
	for {
		if a == nil {
			break
		}
		fmt.Print(a.Val)
		a = a.Next
	}
	fmt.Print("\n")
}
