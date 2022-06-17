package algorithm

import (
	"fmt"
	"strconv"
	"strings"
)

/**
实现一个算法，确定一个字符串 s 的所有字符是否全都不同。
示例 1：
输入: s = "leetcode"
输出: false
示例 2：
输入: s = "abc"
输出: true
限制：
0 <= len(s) <= 100
如果你不使用额外的数据结构，会很加分。

*/
func IsUnique(astr string) bool {
	//思路1、hashmap，空间复杂度O(n)
	//思路2、位运算符。
	var mark uint32 = 0
	for _, ch := range astr {
		//计算出每个字符距离a的距离
		move_bit := ch - 'a'
		uMoveBit := uint32(move_bit)
		//将1往左侧移动uMoveBit位。假如当前的字符是'c'，则'c'-'a'=2，移动2位。与mark做与运算，如果结果不为0，则证明mark中从右往左数第3位原本就有值，证明不符合条件
		if mark&(1<<uMoveBit) != 0 {
			return false
		} else {
			//如果满足条件，则将mark的当前位置位1
			mark |= (1 << uMoveBit)
		}
	}
	return true
}

/**
给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
示例 1：
输入: s1 = "abc", s2 = "bca"
输出: true
示例 2：
输入: s1 = "abc", s2 = "bad"
输出: false
说明：
0 <= len(s1) <= 100
0 <= len(s2) <= 100
*/
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	//将字符串s1放入map中，然后遍历s2，判断s2的每个字符是否在map中有值，如果有则将其数量减掉，最终判断map中是否都为0
	m := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		m[s1[i]]++
	}
	for i := 0; i < len(s2); i++ {
		if v, ok := m[s2[i]]; !ok || v == 0 {
			return false
		}
		m[s2[i]]--
	}
	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}

/**
URL化。编写一种方法，将字符串中的空格全部替换为%20。假定该字符串尾部有足够的空间存放新增字符，并且知道字符串的“真实”长度。（注：用Java实现的话，请使用字符数组实现，以便直接在数组上操作。）

示例1:
 输入："Mr John Smith    ", 13
 输出："Mr%20John%20Smith"
示例2:
 输入："               ", 5
 输出："%20%20%20%20%20"
提示：
字符串长度在[0, 500000]范围内。

*/
func ReplaceSpaces(S string, length int) string {
	//从后往前替换。  首先遍历S，获取S中空格的数量，计算出转换后的S的长度。然后设定一个指针指向字符串的末尾，依次将字符放入正确的位置
	var blankNum int
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			blankNum++
		}
	}
	if blankNum == 0 {
		return S
	}
	//计算出字符串转换完成之后的长度
	realL := length - blankNum + blankNum*3
	//设定两个指针，分别指向转换前的字符串和转换后的字符串
	p1 := length - 1
	p2 := realL - 1
	tmp := make([]byte, realL)
	for p1 >= 0 {
		if S[p1] == ' ' {
			tmp[p2] = '0'
			tmp[p2-1] = '2'
			tmp[p2-2] = '%'
			p1--
			p2 -= 3
		} else {
			tmp[p2] = S[p1]
			p1--
			p2--
		}
	}
	return string(tmp)
}

/**
给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。
回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。
回文串不一定是字典当中的单词。
示例1：
输入："tactcoa"
输出：true（排列有"tacocat"、"atcocta"，等等）
*/
func CanPermutePalindrome(s string) bool {
	//声明一个map,保存字符和其出现的数量。将字符串中字符的值都放入map中，如果最后map中所有字符数量出现奇数的次数少于等于1次，则证明其可以组合成回文串
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	jishuNum := 0
	for _, v := range m {
		if v%2 == 1 {
			jishuNum++
		}
	}
	if jishuNum <= 1 {
		return true
	} else {
		return false
	}
}

/**
字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
示例 1:
输入:
first = "pale"
second = "ple"
输出: True

示例 2:
输入:
first = "pales"
second = "pal"
输出: False

"tactcoa", "taetcoa"
*/
func OneEditAway(first string, second string) bool {
	if first == second {
		return true
	}

	l1 := len(first)
	l2 := len(second)
	//操作次数
	count := 0
	//双指针
	var p1, p2 = 0, 0
	for p1 < l1 && p2 < l2 {
		//如果当前的值相等，则同时将指针往后移动一位
		if first[p1] == second[p2] {
			p1++
			p2++
		} else {
			//如果不相等，则分几种情况
			count++
			//1、如果p1的下一个值等于p2的当前值，则认定为删除p1的当前值
			if p1 < l1-1 && first[p1+1] == second[p2] {
				p1++
			} else if p2 < l2-1 && first[p1] == second[p2+1] {
				//如果当前p1的值等于p2的下一个值，则认定在p1的当前位置插入p2的当前值，将p2往后移动一位
				p2++
			} else if p1 < l1-1 && p2 < l2-1 && first[p1+1] == second[p2+1] {
				//如果当前p1的下一个值和p2的下一个值相等，则认定为替换当前值，将指针都往后移动一位
				p1++
				p2++
			} else {
				p1++
				p2++
			}
		}
	}
	//判断此时某个指针移动到头之后，另一个指针是否到头，如果没有，则需要将操作次数加上差值
	if p1 == l1 && p2 < l2 {
		count += l2 - p2
	}
	if p2 == l2 && p1 < l1 {
		count += l1 - p1
	}

	return count <= 1
}

/**
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。
若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。
示例1:
 输入："aabcccccaaa"
 输出："a2b1c5a3"
示例2:
 输入："abbccd"
 输出："abbccd"
 解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。
提示：
字符串长度在[0, 50000]范围内。
*/
func CompressString(S string) string {
	if len(S) <= 1 {
		return S
	}
	res := ""
	last := S[0]
	lastCount := 1
	S += "*"
	for i := 1; i < len(S); i++ {
		if S[i] == last {
			lastCount++
		} else {
			res += string(last) + strconv.Itoa(lastCount)
			last = S[i]
			lastCount = 1
		}
	}
	if len(res) >= len(S)-1 {
		return S[:len(S)-1]
	}
	return res
}

/**
给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。
不占用额外内存空间能否做到？

示例 1:
给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],
原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
示例 2:
给定 matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],
原地旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]

*/
func Rotate1(matrix [][]int) {
	//思路1，新数组的第一行即为旧数组的第一列倒序
	/*res := make([][]int, 0)
	for i := 0; i < len(matrix[0]); i++ {
		tmp := make([]int, 0)
		for j := len(matrix) - 1; j >= 0; j-- {
			tmp = append(tmp, matrix[j][i])
		}
		res = append(res, tmp)
	}
	matrix = res*/
	/*
		思路：
		1、先将数组上下翻转
		2、再将数据按照“左上-右下”的斜对角线翻转
	*/
	for i := 0; i < len(matrix)/2; i++ {
		line1 := matrix[i]
		line2 := matrix[len(matrix)-1-i]
		for j := 0; j < len(line1); j++ {
			tmp := line1[j]
			line1[j] = line2[j]
			line2[j] = tmp
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			if i == j {
				continue
			}
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
}

/**
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。
示例 1：
输入：
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出：
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
示例 2：
输入：
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出：
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
*/
func SetZeroes(matrix [][]int) {
	//遍历一遍二维数组，将为0的值的位置记录下来
	m := make([][2]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				m = append(m, [2]int{i, j})
			}
		}
	}
	for _, v := range m {
		//行元素清0
		for i := 0; i < len(matrix[0]); i++ {
			matrix[v[0]][i] = 0
		}
		//列元素清0
		for i := 0; i < len(matrix); i++ {
			matrix[i][v[1]] = 0
		}
	}
	fmt.Println(matrix)
}

/**
字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）。
示例1:
 输入：s1 = "waterbottle", s2 = "erbottlewat"
 输出：True
示例2:
 输入：s1 = "aa", s2 = "aba"
 输出：False
提示：
字符串长度在[0, 100000]范围内。
说明:
你能只调用一次检查子串的方法吗？
*/
func isFlipedString(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Index(s1+s1, s2) != -1
}

/**
编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。

示例1:
 输入：[1, 2, 3, 3, 2, 1]
 输出：[1, 2, 3]
示例2:
 输入：[1, 1, 1, 1, 2]
 输出：[1, 2]
提示：
链表长度在[0, 20000]范围内。
链表元素在[0, 20000]范围内。
进阶：
如果不得使用临时缓冲区，该怎么解决？
*/
func RemoveDuplicateNodes(head *ListNode) *ListNode {
	//借助hashmap
	m := make(map[int]int)
	//遍历链表，将没有出现过的值放入到map中，如果出现了值，则将其删掉
	tmp := head
	root := &ListNode{}
	root.Next = head
	parent := root
	for tmp != nil {
		//如果不在map中，则将其记录到map中，并将指针往后移动一位
		if _, ok := m[tmp.Val]; !ok {
			m[tmp.Val] = 1
			parent = tmp
			tmp = tmp.Next
		} else {
			//如果在map中，则将当前节点删除，
			parent.Next = tmp.Next
			tmp = tmp.Next
		}
	}
	printList(root.Next)
	return root.Next
}

/**
实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
注意：本题相对原题稍作改动

示例：
输入： 1->2->3->4->5 和 k = 2
输出： 4
说明：
给定的 k 保证是有效的
*/
func KthToLast(head *ListNode, k int) int {
	//快慢指针。
	var p1, p2 = head, head
	for i := 1; i <= k; i++ {
		p2 = p2.Next
	}
	for p2 != nil {
		p2 = p2.Next
		p1 = p1.Next
	}
	return p1.Val
}

/**
实现一种算法，删除单向链表中间的某个节点（即不是第一个或最后一个节点），假定你只能访问该节点。
示例：
输入：单向链表a->b->c->d->e->f中的节点c
结果：不返回任何数据，但该链表变为a->b->d->e->f
*/
func deleteNode(node *ListNode) {
	*node = *node.Next
}

/**
编写程序以 x 为基准分割链表，使得所有小于 x 的节点排在大于或等于 x 的节点之前。如果链表中包含 x，x 只需出现在小于 x 的元素之后(如下所示)。
分割元素 x 只需处于“右半部分”即可，其不需要被置于左右两部分之间。
示例:
输入: head = 3->5->8->5->10->2->1, x = 5
输出: 3->1->2->10->5->5->8

*/
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nilHead := new(ListNode)
	nilHead.Next = head
	pre := nilHead
	var lessHead, lessTail *ListNode
	for pre.Next != nil {
		if pre.Next.Val < x {
			less := pre.Next
			pre.Next = less.Next
			less.Next = nil
			if lessTail == nil {
				lessTail, lessHead = less, less
			} else {
				// 头插法：按照实际运行结果，此处采用的是头插法
				less.Next = lessHead
				lessHead = less
			}
		} else {
			pre = pre.Next
		}
	}

	if lessHead == nil {
		return nilHead.Next
	} else {
		lessTail.Next = nilHead.Next
		return lessHead
	}
}

/**
给定两个用链表表示的整数，每个节点包含一个数位。
这些数位是反向存放的，也就是个位排在链表首部。
编写函数对这两个整数求和，并用链表形式返回结果。

示例：
输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
输出：2 -> 1 -> 9，即912
进阶：假设这些数位是正向存放的，请再做一遍。
示例：
输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
输出：9 -> 1 -> 2，即912

*/
func AddTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	//遍历链表l1和l2，将每个链表的当前值加起来，并加上进位
	tmp1 := l1
	tmp2 := l2
	var res = &ListNode{}
	resTmp := res
	//进位
	var jinwei, cur1, cur2 int
	for tmp1 != nil || tmp2 != nil || jinwei != 0 {
		if tmp1 != nil {
			cur1 = tmp1.Val
			tmp1 = tmp1.Next
		} else {
			cur1 = 0
		}
		if tmp2 != nil {
			cur2 = tmp2.Val
			tmp2 = tmp2.Next
		} else {
			cur2 = 0
		}
		curNode := &ListNode{}
		cur := cur1 + cur2 + jinwei
		if cur >= 10 {
			cur -= 10
			jinwei = 1
		} else {
			jinwei = 0
		}
		curNode.Val = cur
		resTmp.Next = curNode
		resTmp = resTmp.Next
	}
	printList(res.Next)
	return res.Next
}

/**
编写一个函数，检查输入的链表是否是回文的。

示例 1：
输入： 1->2
输出： false
示例 2：
输入： 1->2->2->1
输出： true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/
func IsPalindrome2(head *ListNode) bool {
	//思路1，将链表的值读到一个整形数组中，检测整形数组是否为回文串。  时间复杂度O(n) (遍历链表一遍，检测字符串是否为回文串一遍)
	arr := make([]int, 0)
	tmp := head
	for tmp != nil {
		arr = append(arr, tmp.Val)
		tmp = tmp.Next
	}
	if len(arr) <= 1 {
		return true
	}
	var p1, p2 = 0, len(arr) - 1
	for p1 < p2 {
		if arr[p1] != arr[p2] {
			return false
		}
		p1++
		p2--
	}
	return true

	//如果要做到时间复杂度O(n)空间复杂度O(1)，做到以下步骤，
	// 1、快慢指针，找到链表的中点。如果是偶数节点的链表，则找到前一个中间节点。
	// 2、然后翻转中间节点以后的链表得到l2
	// 3、遍历原始节点和l2，判断值是否相等
}

/**
给定两个（单向）链表，判定它们是否相交并返回交点。请注意相交的定义基于节点的引用，而不是基于节点的值。
换句话说，如果一个链表的第k个节点与另一个链表的第j个节点是同一节点（引用完全相同），则这两个链表相交。

示例 1：
输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Reference of the node with value = 8
输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。

示例 2：
输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Reference of the node with value = 2
输入解释：相交节点的值为 2 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。

示例 3：
输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
输出：null
输入解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
解释：这两个链表不相交，因此返回 null。

注意：
如果两个链表没有交点，返回 null 。
在返回结果后，两个链表仍须保持原有的结构。
可假定整个链表结构中没有循环。
程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。

*/
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	//双指针法。分别用两个指针p1, p2遍历headA和headB，遍历到头之后，p1再从headB的头开始，p2从headA的头开始，直到两个节点相等
	p1 := headA
	p2 := headB
	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = headB
		}

		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = headA
		}
	}
	return p1
}

/**
给定一个有环链表，实现一个算法返回环路的开头节点。
有环链表的定义：在链表中某个节点的next元素指向在它前面出现过的节点，则表明该链表存在环路。

示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：
输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
输入：head = [1], pos = -1
输出：no cycle
解释：链表中没有环。
进阶：
你是否可以不用额外空间解决此题？

*/
func detectCycle(head *ListNode) *ListNode {
	//毫无疑问，快慢指针。 如果快慢指针不能相遇，证明链表上没有环。如果快慢指针可以相遇，则证明链表上有环，且相遇节点必然在环上
	//等到相遇时，将快指针移动到链表的开头，以与慢指针相同的速度移动，则最终两者相遇的点即为环的起点。（可画图用数学公式证明）
	if head == nil {
		return nil
	}
	var fast, slow = head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	//如果两个节点没有相交，证明没有环
	if fast == nil || fast.Next == nil {
		return nil
	}
	//两个节点相交了，则将其中一个指针移动到链表开头位置
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

/**
堆盘子。设想有一堆盘子，堆太高可能会倒下来。因此，在现实生活中，盘子堆到一定高度时，我们就会另外堆一堆盘子。请实现数据结构SetOfStacks，模拟这种行为。S
etOfStacks应该由多个栈组成，并且在前一个栈填满时新建一个栈。此外，SetOfStacks.push()和SetOfStacks.pop()应该与普通栈的操作方法相同（
也就是说，pop()返回的值，应该跟只有一个栈时的情况一样）。 进阶：实现一个popAt(int index)方法，根据指定的子栈，执行pop操作。
当某个栈为空时，应当删除该栈。当栈中没有元素或不存在该栈时，pop，popAt 应返回 -1.

*/
type StackOfPlates struct {
	stack [][]int
	cap   int
}

func ConstructorStackOfPlates(cap int) StackOfPlates {
	return StackOfPlates{stack: make([][]int, 0), cap: cap}
}

func (this *StackOfPlates) Push(val int) {
	//取出stack中最新一个栈，判断其是否超过了容量，如果是，则新建一个栈
	var cur []int
	if len(this.stack) == 0 {
		if this.cap > 0 {
			cur = make([]int, 0)
			cur = append(cur, val)
			this.stack = append(this.stack, cur)
		}
	} else {
		cur = this.stack[len(this.stack)-1]
		if len(cur) == this.cap {
			newS := make([]int, 0)
			newS = append(newS, val)
			this.stack = append(this.stack, newS)
		} else {
			cur = append(cur, val)
			this.stack[len(this.stack)-1] = cur
		}
	}
}

func (this *StackOfPlates) Pop() int {
	//取出stack中最后一个栈的最后一个元素
	if len(this.stack) == 0 {
		return -1
	}
	last := this.stack[len(this.stack)-1]
	if len(last) == 0 {
		return -1
	}
	tmp := last[len(last)-1]
	if len(last) == 1 {
		//如果last中只有一个元素了，则将last删除掉
		this.stack = this.stack[:len(this.stack)-1]
	} else {
		//将元素弹出
		last = last[:len(last)-1]
		//将last放回stack中
		this.stack[len(this.stack)-1] = last
	}
	return tmp
}

func (this *StackOfPlates) PopAt(index int) int {
	if len(this.stack) <= index {
		return -1
	}
	//获取当前的栈
	cur := this.stack[index]
	if len(cur) == 0 {
		return -1
	}
	tmp := cur[len(cur)-1]
	if len(cur) == 1 {
		//直接将当前cur删掉
		if index == len(this.stack)-1 {
			//如果当前元素即为stack的最后一个栈
			this.stack = this.stack[:len(this.stack)-1]
		} else {
			this.stack = append(this.stack[:index], this.stack[index+1:]...)
		}
	} else {
		//弹出最后一个元素
		cur = cur[:len(cur)-1]
		//将last放回stack中
		this.stack[index] = cur
	}
	return tmp
}

/**
实现一个MyQueue类，该类用两个栈来实现一个队列。
示例：
MyQueue queue = new MyQueue();
queue.push(1);
queue.push(2);
queue.peek();  // 返回 1
queue.pop();   // 返回 1
queue.empty(); // 返回 false
*/
type MyQueue struct {
	stack1 []int
	stack2 []int
}

/** Initialize your data structure here. */
func MyQueueConstructor() MyQueue {
	return MyQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	//将元素放入到第一个栈中
	this.stack1 = append(this.stack1, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	//判断第二个栈中是否有元素，如果有，则直接弹出末尾的元素，如果没有，则将栈1的元素移动到栈2中再弹出。如果都没有，则返回-1
	res := -1
	if len(this.stack2) >= 1 {
		res = this.stack2[len(this.stack2)-1]
		this.stack2 = this.stack2[:len(this.stack2)-1]
	} else {
		if len(this.stack1) == 0 {
			return -1
		}
		//将栈1的元素依次取出放入栈2中
		for i := len(this.stack1) - 1; i >= 0; i-- {
			this.stack2 = append(this.stack2, this.stack1[i])
		}
		this.stack1 = make([]int, 0)
		res = this.stack2[len(this.stack2)-1]
		this.stack2 = this.stack2[:len(this.stack2)-1]
	}
	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	//判断第二个栈中是否有元素，如果有，则直接弹出末尾的元素，如果没有，则将栈1的元素移动到栈2中再弹出。如果都没有，则返回-1
	res := -1
	if len(this.stack2) >= 1 {
		res = this.stack2[len(this.stack2)-1]
	} else {
		if len(this.stack1) == 0 {
			return -1
		}
		//将栈1的元素依次取出放入栈2中
		for i := len(this.stack1) - 1; i >= 0; i-- {
			this.stack2 = append(this.stack2, this.stack1[i])
		}
		this.stack1 = make([]int, 0)
		res = this.stack2[len(this.stack2)-1]
	}
	return res
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0 && len(this.stack2) == 0
}

/**
栈排序。 编写程序，对栈进行排序使最小元素位于栈顶。最多只能使用一个其他的临时栈存放数据，但不得将元素复制到别的数据结构（如数组）中。
该栈支持如下操作：push、pop、peek 和 isEmpty。当栈为空时，peek 返回 -1。

示例1:
 输入：
["SortedStack", "push", "push", "peek", "pop", "peek"]
[[], [1], [2], [], [], []]
 输出：
[null,null,null,1,null,2]
示例2:
 输入：
["SortedStack", "pop", "pop", "push", "pop", "isEmpty"]
[[], [], [], [1], [], []]
 输出：
[null,null,null,null,null,true]

这题可以用链表实现也可以用数组实现。链表实现的话，插入时间复杂度是O(n)一次查找位置 O(1)插入  如果是数组，设计到数据的移位，时间复杂度会高一些
*/
type SortedStack struct {
	minStack []int
}

func SortedStackConstructor() SortedStack {
	return SortedStack{
		minStack: make([]int, 0),
	}
}

func (this *SortedStack) Push(val int) {
	//遍历当前栈的值，找到适合当前val的地方，插入
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
	} else {
		for i := len(this.minStack) - 1; i >= 0; i-- {
			//如果val小于当前值，则证明val应该在当前值的上面
			if val <= this.minStack[i] {
				if i == len(this.minStack)-1 {
					this.minStack = append(this.minStack, val)
				} else {
					tmp := make([]int, 0)
					tmp = append(tmp, this.minStack[i+1:]...)
					this.minStack = append(this.minStack[:i+1], val)
					this.minStack = append(this.minStack, tmp...)
				}
				break
			}
			//如果走到最后一位，还是没有val大，则将val放入栈的最下面
			if i == 0 {
				this.minStack = append([]int{val}, this.minStack...)
			}
		}
	}
}

func (this *SortedStack) Pop() {
	if len(this.minStack) >= 1 {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *SortedStack) Peek() int {
	if len(this.minStack) >= 1 {
		return this.minStack[len(this.minStack)-1]
	}
	return -1
}

func (this *SortedStack) IsEmpty() bool {
	return len(this.minStack) == 0
}

/**
动物收容所。有家动物收容所只收容狗与猫，且严格遵守“先进先出”的原则。在收养该收容所的动物时，收养人只能收养所有动物中“最老”（由其进入收容所的时间长短而定）的动物，
或者可以挑选猫或狗（同时必须收养此类动物中“最老”的）。换言之，收养人不能自由挑选想收养的对象。请创建适用于这个系统的数据结构，实现各种操作方法，
比如enqueue、dequeueAny、dequeueDog和dequeueCat。允许使用Java内置的LinkedList数据结构。
enqueue方法有一个animal参数，animal[0]代表动物编号，animal[1]代表动物种类，其中 0 代表猫，1 代表狗。
dequeue*方法返回一个列表[动物编号, 动物种类]，若没有可以收养的动物，则返回[-1,-1]。

示例1:
 输入：
["AnimalShelf", "enqueue", "enqueue", "dequeueCat", "dequeueDog", "dequeueAny"]
[[], [[0, 0]], [[1, 0]], [], [], []]
 输出：
[null,null,null,[0,0],[-1,-1],[1,0]]
示例2:
 输入：
["AnimalShelf", "enqueue", "enqueue", "enqueue", "dequeueDog", "dequeueCat", "dequeueAny"]
[[], [[0, 0]], [[1, 0]], [[2, 1]], [], [], []]
 输出：
[null,null,null,null,[2,1],[0,0],[1,0]]

*/
type AnimalShelf struct {
	cat [][]int
	dog [][]int
}

func AnimalShelfConstructor() AnimalShelf {
	return AnimalShelf{}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	//0代表猫 1代表狗
	if animal[1] == 0 {
		this.cat = append(this.cat, animal)
	} else {
		this.dog = append(this.dog, animal)
	}
}

func (this *AnimalShelf) DequeueAny() []int {
	if len(this.dog) == 0 && len(this.cat) == 0 {
		return []int{-1, -1}
	}
	if len(this.cat) == 0 {
		x := this.dog[0]
		if len(this.dog) == 1 {
			this.dog = this.dog[:0]
		} else {
			this.dog = this.dog[1:]
		}
		return x
	}
	if len(this.dog) == 0 {
		x := this.cat[0]
		if len(this.cat) == 1 {
			this.cat = this.cat[:0]
		} else {
			this.cat = this.cat[1:]
		}
		return x
	}
	if this.cat[0][0] < this.dog[0][0] {
		x := this.cat[0]
		if len(this.cat) == 1 {
			this.cat = this.cat[:0]
		} else {
			this.cat = this.cat[1:]
		}
		return x
	}
	y := this.dog[0]
	if len(this.dog) == 1 {
		this.dog = this.dog[:0]
	} else {
		this.dog = this.dog[1:]
	}
	return y
}

func (this *AnimalShelf) DequeueDog() []int {
	if len(this.dog) == 0 {
		return []int{-1, -1}
	}
	x := this.dog[0]
	if len(this.dog) == 1 {
		this.dog = this.dog[:0]
	} else {
		this.dog = this.dog[1:]
	}
	return x
}

func (this *AnimalShelf) DequeueCat() []int {
	if len(this.cat) == 0 {
		return []int{-1, -1}
	}
	x := this.cat[0]
	if len(this.cat) == 1 {
		this.cat = this.cat[:0]
	} else {
		this.cat = this.cat[1:]
	}
	return x
}

/**
给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。
示例:
给定有序数组: [-10,-3,0,5,9],
一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

          0
         / \
       -3   9
       /   /
     -10  5
*/
func sortedArrayToBST(nums []int) *TreeNode {
	//高度最小的二叉搜索树，则代表最好左右子树的节点数可以一致
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	mid := len(nums) / 2
	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

/**
给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。

示例：
输入：[1,2,3,4,5,null,7,8]
        1
       /  \
      2    3
     / \    \
    4   5    7
   /
  8

输出：[[1],[2,3],[4,5,7],[8]]
*/
func ListOfDepth(tree *TreeNode) []*ListNode {
	//这不就是广度优先遍历二叉树么
	var res []*ListNode
	queue := make([]*TreeNode, 0)
	if tree != nil {
		queue = append(queue, tree)
	}
	for len(queue) > 0 {
		tmp := queue
		queue = make([]*TreeNode, 0)
		//将tmp中所有的节点取出放入到listNode中，并将子节点放入到queue中
		l := &ListNode{}
		ltmp := l
		for k, v := range tmp {
			ltmp.Val = v.Val
			if k < len(tmp)-1 {
				ltmp.Next = &ListNode{}
				ltmp = ltmp.Next
			}
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		res = append(res, l)
	}
	return res
}

/**
实现一个函数，检查二叉树是否平衡。在这个问题中，平衡树的定义如下：任意一个节点，其两棵子树的高度差不超过 1。
示例 1:
给定二叉树 [3,9,20,null,null,15,7]
    3
   / \
  9  20
    /  \
   15   7
返回 true 。
示例 2:
给定二叉树 [1,2,2,3,3,null,null,4,4]
      1
     / \
    2   2
   / \
  3   3
 / \
4   4
返回 false 。

*/
func isBalanced1(root *TreeNode) bool {
	//这里可以优化。如果左右子树已经是不平衡了，就没必要往下走了。（剪枝）
	if root == nil {
		return true
	}
	tmp := getTreeDeep(root.Left) - getTreeDeep(root.Right)
	if tmp > 1 || tmp < -1 {
		return false
	}
	return isBalanced1(root.Left) && isBalanced1(root.Right)
}

func getTreeDeep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return max(getTreeDeep(root.Left), getTreeDeep(root.Right)) + 1
}

/**
实现一个函数，检查一棵二叉树是否为二叉搜索树。
示例 1:
输入:
    2
   / \
  1   3
输出: true
示例 2:
输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。
*/
func isValidBST1(root *TreeNode) bool {
	//前序遍历
	return isValidBSTF1(root, -1<<63, 1<<63-1)
}

func isValidBSTF1(root *TreeNode, min, max int) bool {
	return root == nil || min < root.Val && root.Val < max &&
		isValidBSTF1(root.Left, min, root.Val) &&
		isValidBSTF1(root.Right, root.Val, max)
}

/**
设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。
如果指定节点没有对应的“下一个”节点，则返回null。
示例 1:
输入: root = [2,1,3], p = 1
  2
 / \
1   3
输出: 2
示例 2:
输入: root = [5,3,6,2,4,null,null,1], p = 6
      5
     / \
    3   6
   / \
  2   4
 /
1
输出: null
*/

func InorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	//题目要求返回中继节点的后续，那么只需要反过来，获取前序遍历的上一个节点即可
	var res = make([]*TreeNode, 0)
	inorderSuccessorF(root, p, &res)
	if len(res) == 0 {
		return nil
	} else {
		return res[len(res)-1]
	}
}

func inorderSuccessorF(root, p *TreeNode, res *[]*TreeNode) bool {
	//由于本题要找中序遍历的下一个节点，那么只需要反过来，以右，中，左的顺序遍历二叉树即可
	if root == nil {
		return false
	}
	flag := inorderSuccessorF(root.Right, p, res)
	if flag == true {
		return true
	}
	if root == p {
		return true
	}
	*res = append(*res, root)
	flag = inorderSuccessorF(root.Left, p, res)
	return flag
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	//二叉搜索树的中继节点的下一个值，则证明要找值比此节点大的第一个值
	res := root
	tmp := root
	for tmp != nil {
		//如果当前节点的值小于等于p节点的值，则p节点一定在当前节点的右子树中
		if tmp.Val <= p.Val {
			tmp = tmp.Right
		} else {
			res = tmp
			tmp = tmp.Left
		}
	}
	if res.Val <= p.Val {
		return nil
	} else {
		return res
	}
}

/**
设计并实现一个算法，找出二叉树中某两个节点的第一个共同祖先。不得将其他的节点存储在另外的数据结构中。注意：这不一定是二叉搜索树。
例如，给定如下二叉树: root = [3,5,1,6,2,0,8,null,null,7,4]

    3
   / \
  5   1
 / \ / \
6  2 0  8
  / \
 7   4
示例 1:
输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
示例 2:
输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。

*/
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	//如果root和p,q某个节点是同一个，那pq的公共祖先一定是root
	if root == p || root == q || root == nil {
		return root
	}
	//判断左子树是否为pq的公共节点
	left := lowestCommonAncestor(root.Left, p, q)
	//判断右子树是否为pq的公共节点
	right := lowestCommonAncestor(root.Right, p, q)
	//如果left和right都不为nil，则返回当前节点
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	} else {
		return right
	}
}

/**
检查子树。你有两棵非常大的二叉树：T1，有几万个节点；T2，有几万个节点。设计一个算法，判断 T2 是否为 T1 的子树。
如果 T1 有这么一个节点 n，其子树与 T2 一模一样，则 T2 为 T1 的子树，也就是说，从节点 n 处把树砍断，得到的树与 T2 完全相同。

示例1:
 输入：t1 = [1, 2, 3], t2 = [2]
 输出：true
示例2:
 输入：t1 = [1, null, 2, 4], t2 = [3, 2]
 输出：false
提示：
树的节点数目范围为[0, 20000]。

*/
func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	//定义一个方法，判断两个树是否为同一个树
	var isSame func(p, q *TreeNode) bool
	isSame = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		return p.Val == q.Val && isSame(p.Left, q.Left) && isSame(p.Right, q.Right)
	}
	if t1 == nil {
		return t2 == nil
	}
	return isSame(t1, t2) || checkSubTree(t1.Left, t2) || checkSubTree(t1.Right, t2)
}

/**
给定一棵二叉树，其中每个节点都含有一个整数数值(该值或正或负)。设计一个算法，打印节点数值总和等于某个给定值的所有路径的数量。
注意，路径不一定非得从二叉树的根节点或叶节点开始或结束，但是其方向必须向下(只能从父节点指向子节点方向)。
示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:
3
解释：和为 22 的路径有：[5,4,11,2], [5,8,4,5], [4,11,7]
提示：
节点总数 <= 10000

*/
var pathSumRes2 [][]int

func PathSum2(root *TreeNode, sum int) int {
	//这里为了防止leetcode报错
	pathSumRes2 = make([][]int, 0)
	//回溯解法
	pathSumF1(root, []int{}, sum, 0)
	return len(pathSumRes2)
}

func pathSumF1(root *TreeNode, cur []int, sum, lastSum int) {
	if root == nil {
		return
	}
	if root.Val+lastSum == sum {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		tmp = append(tmp, root.Val)
		pathSumRes2 = append(pathSumRes2, tmp)
		return
	}
	if root.Val+lastSum > sum && sum > 0 && lastSum > 0 {
		if len(cur) > 0 {
			//将第一个值弹出
			head := cur[0]
			cur = cur[1:]
			pathSumF1(root, cur, sum, lastSum-head)
		} else {
			//如果当前cur没有值，但是当前的val还是大于sum的话，则继续下个节点
			pathSumF1(root.Left, cur, sum, lastSum)
			pathSumF1(root.Right, cur, sum, lastSum)
		}
		return
	}
	cur = append(cur, root.Val)
	lastSum += root.Val
	pathSumF1(root.Left, cur, sum, lastSum)
	pathSumF1(root.Right, cur, sum, lastSum)
}

/**
从左向右遍历一个数组，通过不断将其中的元素插入树中可以逐步地生成一棵二叉搜索树。给定一个由不同节点组成的二叉搜索树，输出所有可能生成此树的数组。
示例：
给定如下二叉树
        2
       / \
      1   3
返回：
[
   [2,1,3],
   [2,3,1]
]

*/
// BSTSequences 二叉搜索树序列（04.09）
func BSTSequences(root *TreeNode) [][]int {
	// 空树直接返回空切片
	if root == nil {
		// 注意这里返回长度为 1 的切片，即 [[]]，返回 [] 会判错
		return make([][]int, 1)
	}
	// q 模拟队列存储所有可能的下一节点
	var q []*TreeNode
	// path 记录路径，为了方便调用 findPath 函数这里直接存进 root.Val
	var path = []int{root.Val}
	// res 存储找到的 path
	var res [][]int

	findPath(root, q, path, &res)

	return res
}

func findPath(root *TreeNode, q []*TreeNode, path []int, res *[][]int) {
	// 有左子节点时添加到队列里
	if root.Left != nil {
		q = append(q, root.Left)
	}
	// 有右子节点时添加到队列里
	if root.Right != nil {
		q = append(q, root.Right)
	}
	// 队列为空时代表找到了一个 path，将 path 添加到 res 中
	if len(q) == 0 {
		*res = append(*res, path)
		return
	}
	// 取出队列中的一个节点作为 nextRoot，同步更新 nextQ 和 nextPath
	for i, nextRoot := range q {
		// 这里不能直接 nextQ := append(q[:i], q[i+1:]...)
		// append() 函数是把后面的元素依次追加到前面的切片中，再用来初始化 nextQ
		// 这种写法会对 q 产生操作，从而影响循环
		tmpQ := make([]*TreeNode, len(q))
		copy(tmpQ, q)
		// nextQ 为 q 取出 nextRoot 所剩
		nextQ := append(tmpQ[:i], tmpQ[i+1:]...)
		// 这里也不能直接 nextPath := append(path, nextRoot.Val)
		// 理由大致同上：在某个节点同时有左右子节点时，会有两个不同的 nextRoot 共用一个 path
		// 这时为第一个 nextRoot 修改 path 会对第二个产生影响
		tmpPath := make([]int, len(path))
		copy(tmpPath, path)
		// nextPath 为 path 追加 nextRoot.Val
		nextPath := append(tmpPath, nextRoot.Val)
		// 递归，直到队列为空
		findPath(nextRoot, nextQ, nextPath, res)
	}
}
