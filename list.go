package main

import "fmt"

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

//打印链表的方法
func printList1(res *ListNode1) {
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

func main() {

	list := &ListNode1{Val: 1, Next: &ListNode1{Val: 2, Next: &ListNode1{Val: 3, Next: &ListNode1{Val: 4, Next: &ListNode1{Val: 5, Next: &ListNode1{Val: 6, Next: &ListNode1{Val: 7, Next: nil}}}}}}}
	//获取第1个节点和第三个节点
	list1Parent, list1 := getListIndex1(list, 1)
	fmt.Println("list1:")
	printList1(list1)
	fmt.Println("list1Parent:")
	printList1(list1Parent)

	list3Parent, list3 := getListIndex1(list, 3)
	fmt.Println("list3:")
	printList1(list3)
	fmt.Println("list3Parent:")
	printList1(list3Parent)
}

//这个方法用于获取链表的第index个节点和其父节点
//注意这里虽然是传递的指针，但是由于在操作list的时候前面没有加*,所以其实也只是操作了值拷贝，不是直接对引用进行的操作
func getListIndex1(list *ListNode1, index int) (*ListNode1, *ListNode1) {
	//*list = *list.Next   这一句才会真正的影响list的值
	x := 1
	parent := &ListNode1{Val: 0, Next: list}
	for list != nil {
		if x == index {
			return parent, list
		} else {
			x++
			parent = list
			list = list.Next
		}
	}
	return nil, nil
}
