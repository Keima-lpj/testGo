package main

import (
	"fmt"
	"testGo/algorithm"
)

func main() {
	//fmt.Println(algorithm.IsUnique("ttnucc"))

	//fmt.Println(algorithm.CheckPermutation("abcd", "bacd"))

	//fmt.Println(algorithm.ReplaceSpaces("Mr John Smith", 13))

	//fmt.Println(algorithm.CanPermutePalindrome("tactcoa"))

	//fmt.Println(algorithm.OneEditAway("a", "b"))

	//fmt.Println(algorithm.CompressString("abbccd"))

	//algorithm.Rotate1([][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}})

	//algorithm.SetZeroes([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}})

	//list := &algorithm.ListNode{Val: 1, Next: &algorithm.ListNode{Val: 2, Next: &algorithm.ListNode{Val: 3, Next: &algorithm.ListNode{Val: 3, Next: &algorithm.ListNode{Val: 2, Next: &algorithm.ListNode{Val: 1, Next: &algorithm.ListNode{Val: 7, Next: &algorithm.ListNode{Val: 8, Next: &algorithm.ListNode{Val: 9, Next: nil}}}}}}}}}
	//algorithm.RemoveDuplicateNodes(list)

	//fmt.Println(algorithm.KthToLast(list, 9))

	/*list1 := &algorithm.ListNode{Val: 1, Next: &algorithm.ListNode{Val: 1, Next: &algorithm.ListNode{Val: 6, Next: nil}}}
	list2 := &algorithm.ListNode{Val: 4, Next: &algorithm.ListNode{Val: 3, Next: &algorithm.ListNode{Val: 8, Next: nil}}}
	algorithm.AddTwoNumbers1(list1, list2)*/

	//list1 := &algorithm.ListNode{Val: 1, Next: &algorithm.ListNode{Val: 2, Next: &algorithm.ListNode{Val: 3, Next: &algorithm.ListNode{Val: 2, Next: &algorithm.ListNode{Val: 1, Next: nil}}}}}
	//list1 := &algorithm.ListNode{Val: -129, Next: &algorithm.ListNode{Val: -129, Next: nil}}
	//fmt.Println(algorithm.IsPalindrome2(list1))

	/*obj := algorithm.ConstructorStackOfPlates(2)
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	fmt.Println(obj.PopAt(0), obj.PopAt(0), obj.PopAt(0))*/

	/*obj := algorithm.SortedStackConstructor()
	obj.Push(42)
	obj.Push(8)
	obj.Push(29)*/

	/*list := &algorithm.TreeNode{Val: 1, Left: &algorithm.TreeNode{Val: 2, Left: &algorithm.TreeNode{Val: 4, Left: &algorithm.TreeNode{Val: 8}}, Right: &algorithm.TreeNode{Val: 5}}, Right: &algorithm.TreeNode{Val: 3, Right: &algorithm.TreeNode{Val: 7}}}
	algorithm.ListOfDepth(list)*/

	/*list := &algorithm.TreeNode{Val: 0}
	fmt.Println(algorithm.IsValidBST1(list))*/
	/*list := &algorithm.TreeNode{Val: 2, Left: &algorithm.TreeNode{Val: 1}, Right: &algorithm.TreeNode{Val: 3}}
	algorithm.InorderSuccessor(list, list.Left)*/

	/*list := &algorithm.TreeNode{Val: 5, Left: &algorithm.TreeNode{Val: 4, Left: &algorithm.TreeNode{Val: 11, Left: &algorithm.TreeNode{Val: 7}, Right: &algorithm.TreeNode{Val: 2}}}, Right: &algorithm.TreeNode{Val: 8, Left: &algorithm.TreeNode{Val: 13}, Right: &algorithm.TreeNode{Val: 4, Left: &algorithm.TreeNode{Val: 5}, Right: &algorithm.TreeNode{Val: 1}}}}
	list = &algorithm.TreeNode{Val: -2, Right: &algorithm.TreeNode{Val: -3}}
	fmt.Println(algorithm.PathSum2(list, -5))*/

	list := &algorithm.TreeNode{Val: 5, Left: &algorithm.TreeNode{Val: 4, Left: &algorithm.TreeNode{Val: 11, Left: &algorithm.TreeNode{Val: 7}, Right: &algorithm.TreeNode{Val: 2}}}, Right: &algorithm.TreeNode{Val: 8, Left: &algorithm.TreeNode{Val: 13}, Right: &algorithm.TreeNode{Val: 4, Left: &algorithm.TreeNode{Val: 5}, Right: &algorithm.TreeNode{Val: 1}}}}
	fmt.Println(algorithm.BSTSequences(list))

}
