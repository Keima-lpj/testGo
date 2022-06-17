package main

import (
	"fmt"
	"testGo/queue"
)

type treeNode struct {
	value       int
	left, right *treeNode
}

type myTreeNode struct {
	node *treeNode
}

//工厂函数，用于新建struct
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

//这里的node是值传递，里面做的操作不能变更node的值
func (node treeNode) setValue(value int) {
	node.value = value
}

//这里的node接受者由于是一个指针，里面的操作会变更这个指针指向的结构体的值
func (node *treeNode) setValue2(value int) {
	node.value = value
}

func main() {
	//新建结构体的几种方法
	var tree1 treeNode
	fmt.Println(tree1) //{0 <nil> <nil>}
	tree2 := treeNode{value: 3}
	fmt.Println(tree2) //{3 <nil> <nil>}
	tree2.left = &treeNode{}
	tree2.right = &treeNode{value: 5, left: nil, right: nil}
	tree2.right.left = new(treeNode)
	tree2.left.right = createNode(2)

	tree2.left.right.setValue(4)
	//这里打印出来任然是2，证明上面的方法确实是值传递，会将tree2的值copy一份传进去
	fmt.Println(tree2.left.right.value)

	//如果这个方法的接受者是指针，则寓言内部会将tree2.left.right的指针传递进去
	tree2.left.right.setValue2(4)
	//这里打印出来是4，证明接受者是指针的话，是引用传递
	fmt.Println(tree2.left.right.value)

	fmt.Println(tree2) //{3 0xc04204c440 0xc04204c460}

	//这里对tree1取地址，即便取了地址，也一样可以正常用里面的函数
	ptree1 := &tree1
	ptree1.setValue2(20)
	fmt.Println(ptree1)

	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	q2 := queue.Queue2{[]int{1, 2, 3}}
	q2.Push([]int{4, 5, 6})
	fmt.Println(q2)
	fmt.Println(q2.Pop())
}
