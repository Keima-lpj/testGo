package main

type TreeNodeB struct {
	Value int        //当前节点的值
	Left  *TreeNodeB //左子节点
	Right *TreeNodeB //右子节点
}

type BinarySearchTree struct {
	Root *TreeNodeB
}

func (tree BinarySearchTree) Insert(v int) {
	tree.Root.Insert(v)
}

//向二叉树中插入数据
func (node *TreeNodeB) Insert(v int) {
	//v小于当前节点的值，往左边插
	if v < node.Value {
		if node.Left != nil {
			node.Left.Insert(v)
		} else {
			node.Left = &TreeNodeB{v, nil, nil}
		}
	} else {
		//大于当前节点的值，往右边插
		if node.Right != nil {
			node.Right.Insert(v)
		} else {
			node.Right = &TreeNodeB{v, nil, nil}
		}
	}
}

//遍历二叉树  树的遍历有前序，后序，中序等等。这里以中序为例。
func (tree BinarySearchTree) InOrder() []int {
	var res []int
	tree.Root.InOrder(&res)
	return res
}

func (node *TreeNodeB) InOrder(result *[]int) {
	if node.Left != nil {
		node.Left.InOrder(result)
	}
	*result = append(*result, node.Value)
	if node.Right != nil {
		node.Right.InOrder(result)
	}
}

//最大最小值
func (tree BinarySearchTree) FindMin() int {
	node := tree.Root
	for {
		if node.Left != nil {
			node = node.Left
		} else {
			return node.Value
		}
	}
}

func (tree BinarySearchTree) FindMax() int {
	node := tree.Root
	for {
		if node.Right != nil {
			node = node.Right
		} else {
			return node.Value
		}
	}
}

//查找某个值
func (tree BinarySearchTree) Contains(v int) bool {
	node := tree.Root
	for {
		if node == nil {
			return false
		} else if node.Value == v {
			return true
		} else if node.Value > v {
			node = node.Left
		} else {
			node = node.Right
		}
	}
}
