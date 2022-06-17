package main

/*
红黑树是一种含有红黑节点并能自平衡的二叉查找树，它必须满足如下性质
1、每个节点要么是黑色的，要么是红色的
2、根节点是黑色
3、每个叶子节点（NIL）是黑色
4、每个红色节点的两个子节点一定都是黑色
5、任意一节点到每个叶子节点的路径都包含相同数量的黑节点

从性质5又可以推出：
5.1、如果一个节点存在黑子节点，那么该节点肯定有两个子节点

*/

//首先定义需要的类型
const (
	RED   = true
	BLACK = false
)

type Item interface {
	Less(than Item) bool
}

//节点
type Node struct {
	Parent *Node //父节点
	Left   *Node //左子节点
	Right  *Node //右子节点
	Color  bool  //颜色
	Item         //存储的Item
}

//红黑树
type RBTree struct {
	NIL   *Node
	Root  *Node
	count uint64
}

//新建一个红黑树
func New() *RBTree {
	node := Node{nil, nil, nil, BLACK, nil}
	return &RBTree{NIL: &node, Root: &node, count: 0}
}

//节点左旋
func (rbt *RBTree) LeftRotate(no *Node) {
	//如果要左旋的节点的右子节点为空，则不做任何处理
	if no.Right == nil {
		return
	}

	//          |                                  |
	//          X                                  Y
	//         / \         left rotate            / \
	//        α   Y       ------------->         X   γ
	//           / \                            / \
	//          β   γ                          α   β

	//因为是左旋，相当于当前节点的右子节点需要“租借”子节点给左侧，所以先将右子节点取出
	rchild := no.Right
	//由于是左旋，所以将右子节点Y的左子节点β赋予当前节点X作为右子节点
	no.Right = rchild.Left
	//如果右子节点Y的左子节点β不为空，则将β的父节点设置为当前节点
	if rchild.Left != nil {
		rchild.Left.Parent = no
	}

	//将右子节点的父节点切换为当前节点的父节点（完成旋转）
	rchild.Parent = no.Parent

	//判断，如果当前节点的父节点为nil，则证明当前节点为根节点
	if no.Parent == nil {
		rbt.Root = rchild
	} else if no == no.Parent.Left {
		//如果当前节点是其父节点的左子节点，则将父节点的左子节点X替换为当前节点的右子节点Y
		no.Parent.Left = rchild
	} else {
		//如果当前节点是其父节点的右子节点，则同样操作
		no.Parent.Right = rchild
	}

	//将当前节点X设置为右子节点Y的左子节点，如右图所示。
	rchild.Left = no
	no.Parent = rchild
}

func (rbt *RBTree) RightRotate(no *Node) {
	//如果当前节点的左节点为空，则直接返回
	if no.Left == nil {
		return
	}

	//          |                                  |
	//          X                                  Y
	//         / \         right rotate           / \
	//        Y   γ      ------------->          α   X
	//       / \                                    / \
	//      α   β                                  β   γ

	//由于是右旋，当前节点的左子节点需要“租借”子节点给右侧，所以先将左子节点取出
	lchild := no.Left
	//将左子节点的右子节点分给当前节点
	no.Left = lchild.Right
	//如果左子节点的右子节点不为空，则将其父节点设置为当前节点
	if lchild.Right != nil {
		lchild.Right.Parent = no
	}

	//将左子节点的父节点切换为当前节点的父节点（转换的过程）
	lchild.Parent = no.Parent

	//判断，如果当前节点为父节点，则
	if no.Parent == nil {
		rbt.Root = lchild
	} else if no == no.Parent.Left {
		//如果当前节点是其父节点的左节点，则
		no.Parent.Left = lchild
	} else {
		no.Parent.Right = lchild
	}

	//将当前节点设置为其左子节点的右子节点
	lchild.Right = no
	no.Parent = lchild

}
