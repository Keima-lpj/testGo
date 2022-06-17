package algorithm

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
func TwoSum(nums []int, target int) []int {
	//先定义一个map
	m := make(map[int]int)
	for k, v := range nums {
		x := target - v
		if k1, ok := m[x]; ok {
			return []int{k, k1}
		}
		m[v] = k
	}
	return []int{-1, -1}
}

/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//1、分别计算出两个链表的值，相加之后，倒过来，再生成链表   (这样不行，测试用例中有很多很长的链表，相加会溢出)

	//2、循环两个链表的每个节点，如果有一个节点为空，则将那个节点的值置为0，将两个节点相加，然后得到的合如果大于10则减去10，小于10不做操作，
	//将这个合放入到输出链表里。如果大于10，则记录进位。
	pre := &ListNode{Val: 0}
	cur := pre

	var carry int

	for l1 != nil || l2 != nil {
		var x, y int
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		//计算当前两个链表的值的和是否进一
		carry = sum / 10
		//计算出生成的链表的值
		sum = sum % 10
		//生成一个节点，放入cur中
		cur.Next = &ListNode{Val: sum}
		//开始处理下个节点
		cur = cur.Next
	}
	//如果循环结束后，此时的carry还为1，证明此时最后一位相加也超出了10，则再往后续一个子节点
	if carry == 1 {
		cur.Next = &ListNode{Val: carry}
	}

	return pre.Next
}

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

/*
	//解法1.暴力破解法，计算出字符串的长度，循环每个字符，往后循环，取出最长的不重复子串，然后比较最终的大小
	max := 1
	l := len(s)
	for i := 0; i < l; i++ {
		//新建一个临时的map
		m := make(map[uint8]int)
		m[s[i]] = 1
		temp := 1
		for j := i + 1; j < l; j++ {
			if _, ok := m[s[j]]; ok {
				break
			} else {
				//将这个值放入m中
				m[s[j]] = 1
				temp++
			}
		}
		if temp > max {
			max = temp
		}
	}
	return max
*/
func getMaxLengthDiffString(str string) (int, string) {
	//定义一个map，用于存储这个字符串的字符和位置
	m := make(map[byte]int)
	start := 0
	maxLength := 0
	returnString := ""
	s := []byte(str)

	for i, ch := range []byte(str) {
		//判断，如果当前这个字节已经存储在m中，且位置大于start，将start置为v+1
		if v, ok := m[ch]; ok && v >= start {
			start = v + 1
		}
		//判断最大长度
		if i-start+1 > maxLength {
			maxLength = i - start + 1
			returnString = string(s[start : i+1])
		}
		m[ch] = i
	}
	return maxLength, returnString
}

//给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
//请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//你可以假设 nums1 和 nums2 不会同时为空。
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//思路1:先将两个数组合并，并从小到大排序。然后判断个数奇偶，然后计算中位数
	/*nums := append(nums1, nums2...)
	//冒泡排序
	nums = func(arr []int) []int {
		lens := len(arr)
		for i := 0; i < lens; i++ {
			for j := i + 1; j < lens; j++ {
				if arr[i] > arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
		return arr
	}(nums)
	x := 0.0
	if len(nums)%2 == 1 {
		//奇数，取中间的值
		x = float64(nums[len(nums)/2])
	} else {
		//偶数，取中间的两个值
		x = float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / 2
	}
	return x*/

	//思路2：在两个数组中新增两个游标，来分别比较两个游标对应的值的大小，将小的一方放入一个新的数组中，本质上还是排序
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2
	i, j := 0, 0
	var temp []int

	for {
		if i < l1 && j < l2 {
			if nums1[i] <= nums2[j] {
				temp = append(temp, nums1[i])
				i++
			} else {
				temp = append(temp, nums2[j])
				j++
			}
		} else if i >= l1 && j < l2 {
			temp = append(temp, nums2[j])
			j++
		} else if i < l1 && j >= l2 {
			temp = append(temp, nums1[i])
			i++
		}

		if len(temp) == l/2+1 {
			if l%2 == 1 {
				return float64(temp[l/2])
			} else {
				return float64(temp[l/2-1]+temp[l/2]) / 2.0
			}
		}
	}

}

//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
func LongestPalindrome(s string) string {
	//思路1，暴力破解
	l := len(s)
	if l < 2 {
		return s
	}
	byteArr := []byte(s)
	maxS := string(byteArr[0])
	maxL := 0

	for i := 0; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			//正向字符串
			str := string(byteArr[i:j])
			//反向字符串
			reStr := reverseString(str)
			if str == reStr && len(byteArr[i:j]) > maxL {
				maxL = j - i
				maxS = str
			}
		}
	}
	return maxS

	//思路2，设定一个游标，从第一个值开始，然后往下走，分别比较游标左右的值，如果是回文，则存储下来，再比较左-1和右+1的值
	//将这个字符串用0分割
	/*if len(s) == 0 || len(s) == 1 {
		return s
	}
	s = strings.Join(strings.Split(s, ""), ".")
	s = "." + s + "."
	l := len(s)
	byteArr := []byte(s)
	maxS := string(byteArr[0])
	maxLength := 1

	for i := 0; i < l; i++ {
		if i == 0 && byteArr[0] == byteArr[1] {
			maxS = string(byteArr[0:2])
			maxLength = 2
		}
		if i == l-1 && byteArr[l-1] == byteArr[l-2] && maxLength == 1 {
			maxS = string(byteArr[l-2 : l])
			maxLength = 2
		}
		loop := i
		//判断，当前游标到数组最左边长还是到最右边长
		if l-i-1 < i+1 {
			loop = l - i - 1
		}
		for j := 1; j <= loop; j++ {
			left := byteArr[i-j]
			right := byteArr[i+j]
			if left == right {
				//是回文子串,判断此字符串的长度和maxLength的大小比较
				if maxLength < len(byteArr[i-j:i+j+1]) {
					maxLength = len(byteArr[i-j : i+j+1])
					maxS = string(byteArr[i-j : i+j+1])
				}
			} else if (byteArr[i] == left || byteArr[i] == right) && maxLength == 1 {
				var b []byte
				maxLength = 2
				maxS = string(append(b, byteArr[i], byteArr[i]))
			} else {
				break
			}
		}
	}
	return strings.Replace(maxS, ".", "", 1001)*/
}

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

/**
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

*/
func Zconvert(s string, numRows int) string {
	//以等差法来计算，以字符串为一个数组，键之间的差值满足 d = 2 * numRows - 2
	if numRows <= 1 || len(s) == 0 {
		return s
	}
	rst := ""
	size := 2*numRows - 2
	//循环行数
	for i := 0; i < numRows; i++ {
		//循环这一行的数据，以size来累计加
		for j := i; j < len(s); j += size {
			rst += string(s[j])
			//加上第一个字符之后，判断这一行如果不是第一行或者最后一行，则中间应该有一个值，这个值的位置是固定的，为j + size - 2*i
			tmp := j + size - 2*i
			if i != 0 && i != numRows-1 && tmp < len(s) {
				rst += string(s[tmp])
			}
		}
	}
	return rst
}

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
func Reverse(x int) int {
	MaxInt32 := 1<<31 - 1
	MinInt32 := -1 << 31

	var num, newNum int

	//循环这个值
	for x != 0 {
		//得到当前x的个位的值
		a := x % 10
		//将上一次循环的num*10，和当前x的个位的值相加
		newNum = num*10 + a
		//将这个值付给num
		num = newNum
		//将x减少一个个位，进行下一次循环
		x = x / 10
		if num > MaxInt32 || num < MinInt32 {
			return 0
		}
	}
	return num
}

/**
请你来实现一个 atoi 函数，使其能将字符串转换成整数。
首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。
当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。
该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。
注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。
在任何情况下，若函数不能进行有效的转换时，请返回 0。
*/
func myAtoi(str string) int {
	//step1：去无效字符
	str = strings.TrimSpace(str)
	if str == "" || (len(str) == 1 && (str < "0" || str > "9")) {
		return 0
	}
	//step2：规范首字符
	flag := ""
	if string(str[0]) == "-" {
		flag = "-"
		str = str[1:len(str)]
	} else if string(str[0]) == "+" {
		str = str[1:len(str)]
	}
	//step3：遍历检测数字0-9
	resStr := "0"
	for i := 0; i < len(str); i++ {
		if string(str[i]) < "0" || string(str[i]) > "9" {
			break
		}
		resStr += string(str[i])
	}
	resStr = flag + resStr
	//step4：转换
	res, err := strconv.ParseInt(resStr, 10, 32)

	const MaxUint = ^uint32(0)
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	//step5：转换异常处理
	if err != nil {
		if flag == "" {
			return MaxInt
		}
		return MinInt
	}
	return int(res)
}

/**
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
*/
func IsPalindrome(x int) bool {
	//思路1、整数转string，然后通过字符串翻转来做
	/*str := strconv.Itoa(x)
	strR := reverseString(str)
	if str == strR {
		return true
	} else {
		return false
	}*/

	//思路2 不将整数转为字符串
	//1、先判断是否为负数，如果是，则直接返回false
	if x < 0 {
		return false
	}

	//2、翻转这个整数，判断前后是否相等
	x1 := x
	var num, numNew int
	for x1 != 0 {
		//将x1除以10，得到余数
		a := x1 % 10
		numNew = num*10 + a
		num = numNew
		x1 = x1 / 10
	}

	if x == num {
		return true
	} else {
		return false
	}

}

/**
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
说明：你不能倾斜容器，且 n 的值至少为 2。
*/
func MaxArea(height []int) int {
	//思路1，暴力计算，循环每一个垂直线之间的差，计算出水的容积
	/*var max int
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			h := 0
			if height[j] >= height[i] {
				h = height[i]
			} else {
				h = height[j]
			}
			if max < h*(j-i) {
				max = h * (j - i)
			}
		}
	}
	return max*/

	//思路2，指针移动法。假设有两个指针，一个指向数组的第一个值，第二个指向数组的最后一个值，求出当前的面积，然后存入一个值中
	//然后判断第一个指针对应的值和第二个指针对应的值哪个大，将小的那个指针往中间移动一格，然后重复
	l := 0
	r := len(height) - 1
	max := 0
	for r-l >= 0 {
		//计算面积
		newMax := 0
		if height[l] >= height[r] {
			newMax = height[r] * (r - l)
			r--
		} else {
			newMax = height[l] * (r - l)
			l++
		}
		if max < newMax {
			max = newMax
		}
	}
	return max

}

/**
罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。
*/
func IntToRoman(num int) string {
	var str string
	var qian, bai, shi, ge int
	qian = num / 1000
	bai = (num - (qian * 1000)) / 100
	shi = (num - (qian*1000 + bai*100)) / 10
	ge = num % 10

	if qian > 0 {
		for i := 0; i < qian; i++ {
			str += "M"
		}
	}

	if bai > 0 {
		if bai <= 3 {
			for i := 0; i < bai; i++ {
				str += "C"
			}
		} else if bai == 4 {
			str += "CD"
		} else if bai == 9 {
			str += "CM"
		} else {
			str += "D"
			if bai >= 6 {
				for i := 1; i <= bai-5; i++ {
					str += "C"
				}
			}
		}
	}

	if shi > 0 {
		if shi <= 3 {
			for i := 0; i < shi; i++ {
				str += "X"
			}
		} else if shi == 4 {
			str += "XL"
		} else if shi == 9 {
			str += "XC"
		} else {
			str += "L"
			if shi >= 6 {
				for i := 1; i <= shi-5; i++ {
					str += "X"
				}
			}
		}
	}

	if ge > 0 {
		if ge <= 3 {
			for i := 0; i < ge; i++ {
				str += "I"
			}
		} else if ge == 4 {
			str += "IV"
		} else if ge == 9 {
			str += "IX"
		} else {
			str += "V"
			if ge >= 6 {
				for i := 1; i <= ge-5; i++ {
					str += "I"
				}
			}
		}
	}
	return str
}

/**
罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。
*/
func RomanToInt(s string) int {
	num := 0
	var prev byte
	for _, v := range []byte(s) {
		if v == 'M' {
			//判断上一个是不是C
			if prev == 'C' {
				num += 800 //这里为什么是800而不是900，因为它前面如果是一个C的话，已经加过一次100了亦如是
			} else {
				num += 1000
			}
		}
		if v == 'D' {
			//判断上一个是不是C
			if prev == 'C' {
				num += 300
			} else {
				num += 500
			}
		}

		if v == 'C' {
			//判断上一个是不是X
			if prev == 'X' {
				num += 80
			} else {
				num += 100
			}
		}

		if v == 'L' {
			//判断上一个是不是X
			if prev == 'X' {
				num += 30
			} else {
				num += 50
			}
		}

		if v == 'X' {
			//判断上一个是不是I
			if prev == 'I' {
				num += 8
			} else {
				num += 10
			}
		}

		if v == 'V' {
			//判断上一个是不是I
			if prev == 'I' {
				num += 3
			} else {
				num += 5
			}
		}

		if v == 'I' {
			num += 1
		}

		prev = v
	}

	return num
}

/**
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。
示例 1:

输入: ["flower","flow","flight"]
输出: "fl"

示例 2:
输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
*/
func LongestCommonPrefix(strs []string) string {
	var str []byte

	if len(strs) == 0 {
		return ""
	}

	//获取数组第一个元素的长度
	l := len(strs[0])

	//按照这个第一个元素的长度来循环
	for i := 0; i < l; i++ {
		var x byte
		y := -1
		//判断第一个元素是否都是一样的
		for k, v := range strs {
			//如果是循环的数组的第一个元素，则直接将第一个元素的第i个值赋予x
			if k == 0 {
				x = v[i]
			} else {
				//判断当前这个值是否等于x，如果不等于，则证明第i个元素不是公共元素。 如果当前循环的字符串长度比较短，没有第i个元素，也同理
				if len(v)-1 < i || v[i] != x {
					break
				}
			}
			//如果循环的是最后一个元素，且也通过了，则证明这个i是公共元素，将其放入str中
			if k == len(strs)-1 {
				y = k
			}
		}

		if y == len(strs)-1 {
			str = append(str, x)
		} else {
			//证明不是第i个元素不是每个字符串都通过了，则跳出循环
			break
		}
	}

	return string(str)
}

/**
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。
例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

*/
func ThreeSum(nums []int) [][]int {

	//1、思路1 暴力破解

	//2、思路2，指针法
	//先将数组从小到大排序
	sort.Sort(IntSlice(nums))

	var res [][]int

	l := len(nums)

	m := make(map[int][]int)

	//指针从第二位开始，到倒数第二位结束
	for i := 1; i <= l-1; i++ {
		first := 0
		last := l - 1
		for last > i && i > first {

			//fmt.Println("转换前：", first, i, last, nums[first], nums[i], nums[last])
			if nums[first]+nums[last]+nums[i] == 0 {

				if v, ok := m[nums[first]]; ok {
					if v[0] == nums[i] || v[0] == nums[last] {
						//fmt.Println("已经存在了，跳过", nums[first], nums[i], nums[last])
						first = first + 1
						last = last - 1
						continue
					}
				}
				//fmt.Println("数据ok，存入", nums[first], nums[i], nums[last])
				res = append(res, []int{nums[first], nums[last], nums[i]})
				m[nums[first]] = []int{nums[i], nums[last]}
				first = first + 1
				last = last - 1

			} else if nums[first]+nums[last]+nums[i] < 0 {
				first = first + 1
			} else {
				last = last - 1
			}

			//fmt.Println("转换后：", first, i, last, nums[first], nums[i], nums[last])

		}
	}

	return res
}

type IntSlice []int

func (s IntSlice) Len() int { return len(s) }

func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

/**
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
*/
func ThreeSumClosest(nums []int, target int) int {

	//1、暴力法    此法会超时
	/*var sum int
	var minCha int
	minCha = 10000

	l := len(nums)
	//循环数组
	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			for k := j + 1; k < l; k++ {
				cha := nums[i] + nums[j] + nums[k] - target
				fmt.Println(nums[i], nums[j], nums[k], cha, minCha)
				if cha < 0 {
					cha = -cha
				}
				if minCha > cha {
					minCha = cha
					sum = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}

	return sum*/

	//2、左右游标法
	var sum int
	minCha := 10000

	l := len(nums)
	//先将数组从小到大排序
	sort.Sort(IntSlice(nums))

	for i := 1; i <= l-1; i++ {
		first := 0
		last := l - 1

		for last-i > 0 && i-first > 0 {
			s := nums[i] + nums[first] + nums[last]
			cha := s - target
			if s-target == 0 {
				return s
			} else if s-target > 0 {
				//太大了，右侧左移一位
				last -= 1
			} else {
				//太小了，左侧右移一位
				first += 1
			}
			//取绝对值
			if cha < 0 {
				cha = -cha
			}
			if cha < minCha {
				minCha = cha
				sum = s
			}
		}
	}

	return sum
}

/**
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
*/
func LetterCombinations(digits string) []string {
	//定义好一个map，用来存放数字对应的字符
	m := make(map[byte][]string)
	m['2'] = []string{"a", "b", "c"}
	m['3'] = []string{"d", "e", "f"}
	m['4'] = []string{"g", "h", "i"}
	m['5'] = []string{"j", "k", "l"}
	m['6'] = []string{"m", "n", "o"}
	m['7'] = []string{"p", "q", "r", "s"}
	m['8'] = []string{"t", "u", "v"}
	m['9'] = []string{"w", "x", "y", "z"}

	result := make([]string, 0)

	//这里为了避免digits为""的情况
	if len(digits) == 0 {
		return result
	}

	f("", digits, &result, m)

	return result
}

/**
回溯算法。输入一个将被拼装的字符串和要回溯的字符串，一个结果集
如果next_digits为空，则证明后续没有需要回溯的字符串了，直接将拼接好的字符串放入结果集中
否则，则获取要回溯的字符串的第一个字符，去m中查到其对应的字符串，然后将这些字符串拼到combination中，并将next_digits从下一位开始，继续调用回溯方法
*/
func f(combination, next_digits string, result *[]string, m map[byte][]string) {
	if len(next_digits) == 0 {
		*result = append(*result, combination)
	} else {
		for _, v := range m[next_digits[0]] {
			f(combination+v, next_digits[1:], result, m)
		}
	}
}

/**
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
注意：
答案中不可以包含重复的四元组。
示例：
给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

*/
func FourSum(nums []int, target int) [][]int {

	var res [][]int

	l := len(nums)

	if l < 4 {
		return res
	}

	//对数组进行排序
	sort.Sort(IntSlice(nums))

	//定义四个指针，k, i, j, h 。k从0开始遍历，i从k+1开始遍历，留下j和h。j指向i+1，h指向数组最后一位
	for k := 0; k < l-3; k++ {
		//当k的值与上一个值相等时，跳过当前循环
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		//获取当前k的最小值，如果最小值直接大于了target，则可以直接跳过循环了
		var min1 = nums[k] + nums[k+1] + nums[k+2] + nums[k+3]
		if min1 > target {
			continue
		}
		//获取当前k的最大值，如果最大值直接小于了target，则也可以直接跳过循环
		var max1 = nums[k] + nums[l-1] + nums[l-2] + nums[l-3]
		if max1 < target {
			continue
		}

		//第二层循环
		for i := k + 1; i < l-2; i++ {
			//当i的值与上一个值相等时，跳过当前循环
			if i > k+1 && nums[i] == nums[i-1] {
				continue
			}
			//定义指针
			var j, h = i + 1, l - 1
			//fmt.Println("最小值：", k, i, j, h, nums[k], nums[i], nums[j], nums[j+1])
			//获取当前i的最小值，如果最小值直接大于了target，则可以直接跳过循环了
			var min2 = nums[k] + nums[i] + nums[j] + nums[j+1]
			if min2 > target {
				continue
			}
			//fmt.Println("最大值：", k, i, j, h, nums[k], nums[i], nums[h-1], nums[h])
			//获取当前k的最大值，如果最大值直接小于了target，则也可以直接跳过循环
			var max2 = nums[k] + nums[i] + nums[h-1] + nums[h]
			if max2 < target {
				continue
			}

			//这里开始操作指针j和h
			for j < h {
				//fmt.Println(k, i, j, h, nums[k], nums[i], nums[j], nums[h])
				sum := nums[k] + nums[i] + nums[j] + nums[h]
				if sum == target {
					x := []int{nums[k], nums[i], nums[j], nums[h]}
					jumpFlag := false
					//如果sum刚好等于目标值，则去重后将其放入到res中
					for _, v := range res {
						if reflect.DeepEqual(v, x) {
							j++
							h--
							jumpFlag = true
							break
						}
					}
					if jumpFlag {
						continue
					}
					//将数组放入res中，
					res = append(res, x)
					//移动指针
					j++
					h--
				} else if sum < target {
					//如果sum小于目标值，则将j往右侧移动一位
					j++
				} else {
					//如果sum大于目标值，则将h往左侧移动一位
					h--
				}
			}
		}
	}

	return res
}

/**
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

*/
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	//递归思想，循环获取需要删除的元素，并将倒数n-1个元素的Next指定为倒数n+1个元素
	//解题思路1、先计算出链表的长度，然后获取倒数第n个值在链表中正向的顺序
	/*l1 := head
	var length = 1
	for l1.Next != nil {
		l1 = l1.Next
		length++
	}

	if length == 1 && n == 1 {
		return nil
	}
	//如果长度和倒数位相等，则证明需要去除第一个值，直接返回head的子元素即可
	if length == n {
		return head.Next
	}

	l := head

	for i := 1; i <= length-n; i++ {
		if i == length-n {
			//将l的Next指向n+1
			l.Next = l.Next.Next
			break
		}
		l = l.Next
	}
	return head*/

	//阶梯思路2、双指针。设定两个指针，一个指向正向顺序为n+1的地方，一个指向正向顺序为0的地方。然后将两个指针一步一步的正向往后移动，一旦第一个指针移动到末尾指向空时
	//第二个指针就指向了要被删除的上一个元素。此时令这个元素的子元素等于其子子元素即可
	dummy := &ListNode{Val: 0, Next: head}
	first := dummy
	second := dummy
	//将第一个指针移动到n+1的位置
	for i := 1; i <= n+1; i++ {
		first = first.Next
	}
	//开始一个单位一个单位的共同正向移动，一旦first为空，则此时的second即为要删除的上一个元素
	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return dummy.Next
}

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

*/
func IsValid(s string) bool {
	//可以这么理解，设定一个栈。每读取到一个元素，则判断其情况。如果是左括号，则将其放入栈中，如果是右括号，则检查栈顶的元素是否是跟其匹配的左括号，
	//如果是，则将左括号弹出（pop）。如果没有，则证明表达式无效。如果最后循环完成，栈中仍有数据，则也证明表达式无效
	b := make([]byte, 0)

	for _, v := range []byte(s) {
		switch v {
		case '(', '[', '{':
			b = append(b, v)
		case ')':
			if len(b) >= 1 && b[len(b)-1] == '(' {
				b = b[:len(b)-1]
			} else {
				return false
			}
		case ']':
			if len(b) >= 1 && b[len(b)-1] == '[' {
				b = b[:len(b)-1]
			} else {
				return false
			}
		case '}':
			if len(b) >= 1 && b[len(b)-1] == '{' {
				b = b[:len(b)-1]
			} else {
				return false
			}
		}
		//fmt.Println(string(b))
	}

	if len(b) > 0 {
		return false
	}
	return true
}

/**
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

*/
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	cur := pre
	var x, y int
	//先循环第一个链表
	for l1 != nil {
		x = l1.Val
		//如果第二个链表不为空，则比较第一个链表和第二个链表的大小，将小的放入新的链表中，并将对应的链表往后移动一位
		if l2 != nil {
			y = l2.Val
			if x > y {
				cur.Next = &ListNode{Val: y}
				l2 = l2.Next
				cur = cur.Next
			} else if x < y {
				cur.Next = &ListNode{Val: x}
				l1 = l1.Next
				cur = cur.Next
			} else {
				//如果两个值相等，则往新链表中塞两个节点，并且同时移动l1和l2
				cur.Next = &ListNode{Val: x}
				cur.Next.Next = &ListNode{Val: y}
				cur = cur.Next.Next
				l1 = l1.Next
				l2 = l2.Next
			}
		} else {
			//如果l2为空的话，则直接将l1的值塞入新的链表中，并移动l1
			cur.Next = &ListNode{Val: x}
			l1 = l1.Next
			cur = cur.Next
		}
	}
	//如果第一个循环完了第二个还有节点，则直接将剩下的节点拼到后面
	for l2 != nil {
		cur.Next = &ListNode{Val: l2.Val}
		cur = cur.Next
		l2 = l2.Next
	}

	return pre.Next
}

func MergeTwoLists1(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// 创建最终返回头的前置节点
	pre := &ListNode{}
	node := pre
	for l1 != nil && l2 != nil {
		//判断，如果当前l1的值大于l2的值，则将l2置为node.Next，反之亦然
		if l1.Val > l2.Val {
			node.Next = l2
			l2 = l2.Next
		} else {
			node.Next = l1
			l1 = l1.Next
		}
		//将node往后移动一位
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}

	return pre.Next
}

/**
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

*/
func GenerateParenthesis(n int) []string {
	var res []string
	//使用回溯方法
	generateF(&res, "", 0, 0, n)
	return res
}

func generateF(res *[]string, cur string, left, right, max int) {
	//判断，如果当前字符串的长度已经达到了n的两倍，则证明其已经达到最大长度，将其放入res中
	if len(cur) == max*2 {
		*res = append(*res, cur)
		return
	}
	//如果左括号的数量小于n，则可以往其中放入左括号
	if left < max {
		generateF(res, cur+"(", left+1, right, max)
	}
	//如果右括号的数量小于左括号，则可以往其中放入右括号
	if right < left {
		generateF(res, cur+")", left, right+1, max)
	}
}

/**
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/
func MergeKLists(lists []*ListNode) *ListNode {
	//获取所有链表的长度，得到一个总长
	/*res := &ListNode{}
	cur := res
	var allL int
	for _, v := range lists {
		v1 := v
		l := 0
		for {
			if v1 == nil {
				break
			}
			l++
			v1 = v1.Next
		}
		allL += l
	}

	if allL == 0 {
		return nil
	}

	//从第一位循环到这个总长
	for i := 1; i <= allL; i++ {
		min := 100000
		key := 0
		for k, v := range lists {
			//比较每个v的值，取其中最小的一个，放入res中，并将其往后移动一位
			if v != nil && v.Val < min {
				min = v.Val
				key = k
				//fmt.Println("当前循环值的键值和最小值：", k, v.Val, min)
			}
		}
		//fmt.Println("获取当前最小值：", min, "最小值的所属值在lists中的键：", key)
		cur.Val = min
		if i != allL {
			cur.Next = &ListNode{}
			cur = cur.Next
		}

		lists[key] = lists[key].Next
	}

	return res*/

	//解法2、递归，将K个链表转换为两个链表的合并问题

	l := len(lists)
	if l == 0 {
		return nil
	}

	if l == 1 {
		return lists[0]
	}
	//这里假设是一个三个元素的数组，则此时lists[:l/2]为第0个元素，lists[l/2:]为第一个和第二个元素
	//此时第一个和第二个元素继续递归，会先合并成一个链表，再跟第0个合成一个链表
	return MergeTwoLists(MergeKLists(lists[:l/2]), MergeKLists(lists[l/2:]))

}

/**
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。


示例:
给定 1->2->3->4, 你应该返回 2->1->4->3.

*/
func SwapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	//获取链表的长度
	cur := head
	l := 0
	for cur != nil {
		l++
		cur = cur.Next
	}

	if l == 1 {
		return head
	}

	cur = head
	parent := &ListNode{Val: 0, Next: head}
	p1 := parent
	for i := 1; i <= l; i += 2 {
		//交换当前节点和其子节点
		printList(p1)
		printList(cur)
		//避免长度为奇数时的报错
		if cur.Next == nil {
			break
		}
		//将当前节点的父级的子节点指向当前节点的子节点
		p1.Next = cur.Next
		printList(p1)
		//将当前节点的子节点切换为其孙子节点
		cur.Next = cur.Next.Next
		printList(cur)
		//将当前父节点的孙子节点换为当前节点
		p1.Next.Next = cur
		printList(p1)

		//往后移动两位父节点
		p1 = p1.Next.Next
		//往后移动一位当前节点
		cur = cur.Next
	}

	return parent.Next
}

/**
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

示例 :
给定这个链表：1->2->3->4->5
当 k = 2 时，应当返回: 2->1->4->3->5   12345
当 k = 3 时，应当返回: 3->2->1->4->5

说明 :
你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/
func ReverseKGroup(head *ListNode, k int) *ListNode {
	//设定两个游标，先指向第一个节点跟第k个节点，然后将第一个节点和第k个节点交换
	//完成后，将左边的游标往右移动一位，游标的游标往左移动一位。这个时候要判断，如果中间剩余长度没有小于2，则不处理
	//然后再循环k+1到2k，依次这么做
	if k == 1 {
		return head
	}
	//先获取链表的长度
	var l int
	lhead := head
	for lhead != nil {
		l++
		if lhead.Next == nil {
			break
		}
		lhead = lhead.Next
	}
	//循环l/k次
	for i := 1; i <= l; i += k {
		//如果此时l-i小于k，则跳过循环
		if l-i+1 < k {
			break
		}
		x, y := i, i+k-1
		//交换第i个和第k个元素   1234567
		for y-x >= 1 {
			head = swapIAndK(head, x, y)
			x++
			y--
		}
	}

	return head
}

//交换一个链表中的第i个节点和第k个节点
func swapIAndK(list *ListNode, i, k int) *ListNode {
	cur := list
	//获取第i个元素和其父元素
	iParent, iNode := getListIndex(cur, i)
	cur = list
	//获取第k个元素和其父元素
	kParent, kNode := getListIndex(cur, k)

	//交换i和k的值
	//这里判断i的下一个元素是不是k，如果是，则直接两两交换
	if iNode.Next == kNode {
		//1、先令i的父节点的子节点为k
		iParent.Next = kNode
		//2、再令i的子节点为k的子节点
		iNode.Next = kNode.Next
		//3、令k的子节点为i
		kNode.Next = iNode
	} else {
		//这里获取到iParent的子节点的子节点
		x := iParent.Next.Next
		//1、将i的子节点设置为k的子节点 假设i是1k是3  得到i = 14567
		iNode.Next = kNode.Next
		//2、把k的父节点的子节点设置为i  得到214567
		kParent.Next = iNode
		//3、将i的父节点的子节点设置为k 得到34567
		iParent.Next = kNode
		//4、将x设置为k的父节点
		kNode.Next = x
	}
	if i == 1 {
		return iParent.Next
	} else {
		return list
	}
}

//这个方法用于获取链表的第index个节点和其父节点
func getListIndex(list *ListNode, index int) (*ListNode, *ListNode) {
	x := 1
	parent := &ListNode{Val: 0, Next: list}
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

func ReverseKGroup1(head *ListNode, k int) *ListNode {
	//为什么要这么做？因为定义一个头结点可以回避很多特殊情况
	dummy := &ListNode{Val: 0, Next: head}
	//定义一个开始节点，一个结束节点，用于翻转
	pre := dummy
	end := dummy

	for end.Next != nil {
		//将end节点移动到本次翻转的最后一个节点
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		//如果当前的end节点为nil，则证明后续节点的个数不足本次翻转，直接跳过
		if end == nil {
			break
		}
		//定义本次翻转的开始节点
		start := pre.Next
		//下次翻转的开始节点
		next := end.Next
		//令本次翻转的最后一个节点的下一个节点为nil
		end.Next = nil
		//执行翻转  由于执行了end.Next = nil   所以本次翻转只会翻转需要的个数字符串
		fmt.Println("翻转前的start")
		printList(start)
		pre.Next = reverseList(start)
		fmt.Println("翻转后")
		printList(pre.Next)
		//然后将开始节点置为下一个要翻转的字符串的头结点
		start.Next = next
		pre = start
		end = pre
	}
	//这里为什么返回dummy.Next？ 因为每次变换的只是pre和end，当pre和end完成一次变换之后，会重新赋值继续下一次变换
	//只有dummy.Next的位置是始终不变的
	return dummy.Next
}

//翻转链表元素
func reverseList(head *ListNode) *ListNode {
	pre := &ListNode{}
	pre = nil
	curr := head
	for curr != nil {
		//将curr的下一个链表记录下来（当前循环只处理当前节点）
		next := curr.Next
		//将当前节点的子节点设置为pre，如果是第一次循环，pre为nil
		curr.Next = pre
		//将pre设置为当前节点，此时pre为当前节点 + 之前的pre
		pre = curr
		//将curr设置为下一个链表，开始下一次的循环
		curr = next
	}
	return pre
}

/**
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
*/
func RemoveDuplicates(nums []int) int {
	var old int
	old = -100

	l := len(nums)

	for i := 0; i < l; i++ {
		if i >= len(nums) {
			break
		}
		if old == nums[i] {
			//删除当前元素
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			old = nums[i]
		}
	}

	return len(nums)
}

/**
给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

*/
func RemoveElement(nums []int, val int) int {
	l := len(nums)

	for i := 0; i < l; i++ {
		if i >= len(nums) {
			break
		}
		if val == nums[i] {
			//删除当前元素
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}

/**
实现 strStr() 函数。
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:
输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:
输入: haystack = "aaaaa", needle = "bba"
输出: -1
*/
func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	var pos int

	//定义两个游标，一个指向haystack的第一个字符，一个指向needle的第一个字符。 如果两个游标的值相等，则将两者的游标往后移动一位
	//如果此时都不相等，则将needle的游标调整到第一位，重新开始匹配。 如果needle都匹配完了，则此时返回haystack匹配第一个字符的位置

	lh := len(haystack)
	ln := len(needle)

	if ln > lh {
		return -1
	}

	nFlag := 0
	success := false

	for i := 0; i < lh; i++ {
		//fmt.Println(string(haystack[i]), string(needle[nFlag]), nFlag, pos)
		if haystack[i] == needle[nFlag] {
			if nFlag == 0 {
				pos = i
			}
			//如果匹配成功，开始匹配下个值
			nFlag++
			//如果当前的游标等于了ln，此时证明完全匹配成功，跳过整个循环
			if nFlag == ln {
				success = true
				break
			}
		} else {
			//匹配不成功，继续匹配下一个值
			if nFlag != 0 {
				//将i重新置于pos开始的位置，开始pos+1个循环
				i = pos
				nFlag = 0
			}
		}
		if i == lh-1 && !success {
			pos = -1
		}
	}

	return pos
}

/**
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
返回被除数 dividend 除以除数 divisor 得到的商。

示例 1:
输入: dividend = 10, divisor = 3
输出: 3
示例 2:
输入: dividend = 7, divisor = -3
输出: -2

说明:
被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1

*/
func Divide(dividend int, divisor int) int {
	MaxInt32 := 1<<31 - 1
	MinInt32 := -1 << 31

	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend > MinInt32 {
			return -dividend
		}
		return MaxInt32
	}

	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}

	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	res := div(dividend, divisor)

	if sign == 1 {
		if res > MaxInt32 {
			return MaxInt32
		} else {
			return res
		}
	} else {
		return -res
	}
}

//递归除法
func div(a, b int) int {
	if a < b {
		return 0
	}
	//由于a>=b，则默认会有一个count
	count := 1
	//重新赋值的原因是为了不改变b
	tb := b
	//这里循环，每次将tb变成2倍。  判断是否依然小于a，如果是，则将count赋值为本身的两倍
	for tb+tb <= a {
		count = count + count
		tb = tb + tb
	}
	//循环结束，证明这时2tb已经大于了a，用a-tb，得到剩下不足一个tb的值，再用这个值继续除以b
	return count + div(a-tb, b)
}

/**
给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
示例 1：
输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。

示例 2：
输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]

ababababab , [a, b]
*/
func FindSubstring(s string, words []string) []int {
	//思路1、循环words，组合好每个可能，放入一个新的slice中，循环这个slice，看s中有哪些符合   这种会超时

	//思路2、用两个map来解决。首先，把所有的单词存到map中，键为单词，值为单词出现的次数。然后扫描子串中的单词，如果当前扫描的单词在之前的
	//map中，就证明这个单词有效，放入第二个map中，并判断第二个map中的该单词次数是否大于第一个map中的单词次数，如果是（或者第一个map中压根没有），
	//则证明当前子串中的该单词已经超了，当前子串不符合要求，跳过开始循环下一个子串。
	//如果扫描子串完毕，且map2中的单词数量和map1数量相等，则证明这个子串是我们要找的子串，将其保存。
	var res []int

	ls := len(s)
	wordNum := len(words)
	if ls == 0 || wordNum == 0 {
		return res
	}

	//数组中单词的大小，个数
	oneWordLen := len(words[0])

	//建立单词和单词在数组中个数的映射
	m1 := make(map[string]int)
	for _, v := range words {
		m1[v]++
	}

	//遍历所有子串
	for i := 0; i < ls-wordNum*oneWordLen+1; i++ {
		//map2存储当前扫描的字符串含有的单词
		m2 := make(map[string]int)
		var num int
		for num < wordNum {
			//从子串中取出当前需要比对的单词（当前需要比对的单词为num开始的oneWordLen个长度）
			w := s[i+num*oneWordLen : i+(num+1)*oneWordLen]
			if _, ok := m1[w]; ok {
				//这个字符串在m1中存在，将其在m2中的个数加一
				m2[w]++
				if m2[w] > m1[w] {
					break
				}
			} else {
				//这个字符串在m1中不存在，证明这个子串不合格，直接跳过循环
				break
			}
			num++
		}
		if num == wordNum {
			res = append(res, i)
		}
	}
	fmt.Println(res)
	return res
}

func FindSubstring1(s string, words []string) []int {
	//思路3、在思路2的基础上。存在三种情况。
	// 第一种是当前子串完全匹配。 这个之后直接往后移动oneWordLen位。移动完成后，这个时候无需将m2清空，只需要将上一个子串在m2中的数量减1就可以了
	//第二种是，当前子串在匹配的时候出现了不在m1中的单词时，可以直接移动到这个单词的下一位开始循环下一个子串
	//第三种是，出现了符合要求的单词，但是次数超了。那此时只需要将这个超次数的单词移除即可
	var res []int

	ls := len(s)
	wordNum := len(words)
	if ls == 0 || wordNum == 0 {
		return res
	}

	//数组中单词的大小，个数
	oneWordLen := len(words[0])

	//建立单词和单词在数组中个数的映射
	m1 := make(map[string]int)
	for _, v := range words {
		m1[v]++
	}

	//将所有移动分为oneWordLen种情况
	for i := 0; i < oneWordLen; i++ {
		m2 := make(map[string]int)
		var num int
		//这里相对于思路2，每次循环oneWordLen个长度
		for j := i; j < ls-wordNum*oneWordLen+1; j = j + oneWordLen {
			//这里定义一个变量，防止情况3移除后，情况1继续移除
			hasRemoved := false
			for num < wordNum {
				w := s[j+num*oneWordLen : j+(num+1)*oneWordLen]
				if _, ok := m1[w]; ok {
					//如果单词在m1中，则继续
					m2[w]++
					if m2[w] > m1[w] {
						//出现情况3，遇到符合的单词，但是次数超过了
						hasRemoved = true
						var removeNum int
						for m2[w] > m1[w] {
							//这里一直循环去除，直到两者相等为止
							firstWord := s[j+removeNum*oneWordLen : j+(removeNum+1)*oneWordLen]
							m2[firstWord]--
							removeNum++
						}
						num = num - removeNum + 1
						j = j + (removeNum-1)*oneWordLen
						break
					}

				} else {
					//如果单词不在m1中，则出现了情况2，将下次开始的子串位置直接置为当前单词的下一位
					m2 = make(map[string]int) //清空m2
					j = j + num*oneWordLen    //跳转到当前单词的下一个值
					num = 0
					break
				}
				num++
			}
			if num == wordNum {
				res = append(res, j)
			}

			//这里判断是否会出现情况1,如果出现了，则将上一个子串的第一个单词从m2中移除
			if num > 0 && !hasRemoved {
				firstWord := s[j : j+oneWordLen]
				m2[firstWord]--
				num = num - 1
			}
		}
	}
	return res
}

/**
给定一个字符串数组，长度为l, 将这个字符串数组聚合为各种顺序打乱的字符串
*/
func findSubstringF(combine string, words []string) []string {
	//回溯的思想
	l := len(words)
	if l == 0 {
		return []string{combine}
	} else if l == 1 {
		return []string{combine + words[0]}
	} else {
		var res []string
		for i := 0; i < l; i++ {
			//注意，这里不能直接用等于复制，否则为引用赋值，改了lastWords原本的words也会改
			lastWords := make([]string, l)
			copy(lastWords, words)
			//下一个字符串
			lastS := combine + words[i]
			//剔除当前元素
			if i == l-1 {
				lastWords = lastWords[:i]
			} else {
				lastWords = append(lastWords[:i], lastWords[i+1:]...)
			}
			r := findSubstringF(lastS, lastWords)
			res = append(res, r...)
		}
		return res
	}
}

/**
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
必须原地修改，只允许使用额外常数空间。
以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

题干的意思是：找出这个数组排序出的所有数中，刚好比当前数大的那个数
比如当前 nums = [1,2,3]。这个数是123，找出1，2，3这3个数字排序可能的所有数，排序后，比123大的那个数 也就是132
如果当前 nums = [3,2,1]。这就是1，2，3所有排序中最大的那个数，那么就返回1，2，3排序后所有数中最小的那个，也就是1，2，3 -> [1,2,3]

*/
func NextPermutation(nums []int) {
	//思路  其实只要从末尾开始，从右到左判断就好了，如果左侧比右侧小，则记录当前左侧的值，取到当前值右侧大于当前值的最小值，跳出循环即可。
	//如果一直比较下来也没有符合的，那就直接将数组翻转
	swap := func(nums []int, i, j int) {
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
	}
	reverse := func(nums []int, start int) {
		i := start
		j := len(nums) - 1
		for i < j {
			swap(nums, i, j)
			i++
			j--
		}
	}

	l := len(nums)

	//如果此时i+1 > i,则循环停止，获取到i的值
	var i = l - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	//如果i大于等于0，则从数组的最右边再开始往i的位置找，
	if i >= 0 {
		j := l - 1
		for j >= i && nums[j] <= nums[i] {
			j--
		}
		swap(nums, i, j)
	}
	//这一步的目的是为了将i到数组最右侧的一段翻转，因为这一段只可能是从小到大的顺序
	reverse(nums, i+1)
	fmt.Println(nums)
}

/**
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"
*/
func LongestValidParentheses(s string) int {
	//1、暴力法，获取字符串中的每个偶数子串，并判断其是否为有效括号，如果是，和当前maxlength比较   (这种会超时，不采取)
	/*isValid := func(str string) bool {
		//这里还是采用栈的思想，新建一个栈，将s的字符串一个一个的往这个栈里塞，如果取到的是右括号，则判断栈顶上是不是左括号，如果是，则将栈里的左括号弹出，和当前的右括号
		//组合成有效括号，并将当前有效字符串的长度加2。如果不是，则跳过当前的右括号。如果当前是左括号，则将左括号塞入到栈中。
		stack := make([]byte, 0)
		for _, v := range str {
			if v == '(' {
				stack = append(stack, '(')
			} else {
				//右括号，从栈顶端弹出一个值，如果是左括号，则抵消
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		}

		if len(stack) > 0 {
			return false
		}
		return true
	}

	ls := len(s)
	max := 0

	for i := 0; i < ls; i++ {
		for j := i + 2; j <= ls; j += 2 {
			if isValid(s[i:j]) {
				if max < j-i {
					max = j - i
				}
			}
		}
	}
	return max*/

	/*动态规划  定义一个dp数组，数组的每个下标为i的值保存到当前字符串下标i为止的有效字符个数，例如
	())((())
	02000024
	在循环第i个元素的时候，有两种情况，一种是'('这种数组下标必定为0，因为'('不闭合。所以只需要判断是')'的情况即可。是')'的情况下，有两种可能
	1、上一个元素是'('这种情况下，上个元素和这个元素直接组合成'()'只需要将上上个元素的数组值加2即可
	2、上一个元素是')'这种情况下，获取上个元素的数组值，如果此数组值不为0且值为x，证明此数组的前x个元素刚好可以组合成有效字符。则获取这些有效字符的再前一个值
	如果这个值是'('，则可以组合为有效数组，将前一个数组元素的值加2即可。 还没完，如果前面还有一段是有效字符，如()((()))这种情况，则需要将前面的有效字符的长度
	加起来
	*/
	/*var max int
	ls := len(s)
	dp := make([]int, ls)

	for i := 1; i < ls; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] = 2
				if i >= 2 {
					dp[i] += dp[i-2]
				}
			} else if i-dp[i-1] >= 1 && s[i-dp[i-1]-1] == '(' {
				dp[i] = dp[i-1] + 2
				if i-dp[i-1]-2 >= 0 { //这里防止溢出
					dp[i] += dp[i-dp[i-1]-2]
				}
			}
			if dp[i] > max {
				max = dp[i]
			}
		}

	}
	return max*/

	/**
	3、用栈
	对于遇到的每个 ‘(’ ，我们将它的下标放入栈中。
	对于遇到的每个 ‘)’ ，我们弹出栈顶的元素并将当前元素的下标与弹出元素下标作差，得出当前有效括号字符串的长度。通过这种方法，
	我们继续计算有效子字符串的长度，并最终返回最长有效子字符串的长度。

	*/
	var max int
	stack := make([]int, 0)
	stack = append(stack, -1)
	ls := len(s)
	for i := 0; i < ls; i++ {
		fmt.Println("stack1:", stack)
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			//弹出栈中的最上面一个元素
			stack = stack[:len(stack)-1]
			fmt.Println("stack2:", stack)
			if len(stack) == 0 {
				//将当前元素放入到栈中
				stack = append(stack, i)
				fmt.Println("stack3:", stack)
			} else {
				j := stack[len(stack)-1] //这里的j是i前面的没有被消除掉的最小坐标，可能是-1，也可能是')'所在的下标
				fmt.Println("stack4:", i, j, max)
				if i-j > max {
					max = i - j
				}
			}
		}
	}
	return max
}

/**
假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
你可以假设数组中不存在重复的元素。
你的算法时间复杂度必须是 O(log n) 级别。

示例 1:
输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
示例 2:
输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1
*/
func Search(nums []int, target int) int {
	//典型的二分查找，由于此数组是升序数组的翻转，则符合一个定理，在翻转节点后的数组值，一定小于翻转节点前的数组值（如4,5,6,7,0,1,2）
	//判断，如果当前中点值大于左侧的值，则证明数组左侧是升序。否则则证明右侧是升序
	var left, right = 0, len(nums) - 1
	mid := left + (right-left)/2

	for left <= right {
		if nums[mid] == target {
			return mid
		}
		//左边升序
		if nums[left] <= nums[mid] {
			//如果目标在中间节点的左侧，且左侧为升序，则只能从左侧找
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			//右边升序
			//如果目标在中间节点的右侧
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		mid = left + (right-left)/2
	}
	return -1
}

/**
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
你的算法时间复杂度必须是 O(log n) 级别。
如果数组中不存在目标值，返回 [-1, -1]。

示例 1:
输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:
输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/
func SearchRange(nums []int, target int) []int {
	//二分查找走你
	/*erfen := func(nums []int, target int) int {
		var left, right = 0, len(nums) - 1
		var mid = left + (right-left)/2
		for left <= right {
			mid = left + (right-left)/2
			if nums[mid] == target {
				return mid
			} else if nums[mid] > target {
				//从左侧取
				right = mid - 1
			} else if nums[mid] < target {
				//从右侧开始取
				left = mid + 1
			}
		}
		return -1
	}*/
	//查找最左侧的满足数据的二分查找
	erfen2 := func(nums []int, target int) int {
		if len(nums) == 0 {
			return -1
		}
		var left, right = 0, len(nums) //这里跟基本的二分不一样
		var mid = left + (right-left)/2
		for left < right { //[left, right)
			mid = left + (right-left)/2
			if nums[mid] == target {
				//收紧右侧
				right = mid //为什么这里不+1？ 因为右侧是开区间  [left, right)
			} else if nums[mid] > target {
				//收紧右侧
				right = mid
			} else if nums[mid] < target {
				//从左侧开始取
				left = mid + 1
			}
		}
		//如果left已经加到了数组的长度，则证明没有找到，返回-1
		if left == len(nums) {
			return -1
		} else {
			//如果此时left的值确实等于target
			if nums[left] == target {
				//这里返回left或right都一样，因为两个值最后相等
				return left
			} else {
				return -1
			}
		}
	}

	//查找最左侧的满足数据的二分查找
	erfen3 := func(nums []int, target int) int {
		if len(nums) == 0 {
			return -1
		}
		var left, right = 0, len(nums) //这里跟基本的二分不一样
		var mid = left + (right-left)/2
		for left < right { //[left, right)
			mid = left + (right-left)/2
			if nums[mid] == target {
				//收紧左侧
				left = mid + 1 //为什么这里是+1？ 因为左侧是闭区间，右侧是开区间[left, right)
			} else if nums[mid] > target {
				//收紧右侧
				right = mid
			} else if nums[mid] < target {
				//从左侧开始取
				left = mid + 1
			}
		}

		if left == 0 {
			return -1
		} else {
			if nums[left-1] == target {
				//由于left = mid+1
				return left - 1
			} else {
				return -1
			}
		}
	}

	return []int{erfen2(nums, target), erfen3(nums, target)}
}

/**
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
你可以假设数组中无重复元素。
示例 1:
输入: [1,3,5,6], 5
输出: 2
示例 2:
输入: [1,3,5,6], 2
输出: 1
示例 3:
输入: [1,3,5,6], 7
输出: 4
示例 4:
输入: [1,3,5,6], 0
输出: 0

*/
func SearchInsert(nums []int, target int) int {
	l := len(nums)

	for i := 0; i < l; i++ {
		if nums[i] == target {
			return i
		} else if nums[i] > target {
			//此时nums[i] > target，而且没有返回。证明没有target的值，则直接返回
			return i
		}

		if i == l-1 {
			return i + 1
		}
	}

	return 0
}

/**
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。
数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
数独部分空格内已填入了数字，空白格用 '.' 表示。

*/
func IsValidSudoku(board [][]byte) bool {
	rows := make([]map[int]int, 9)
	columns := make([]map[int]int, 9)
	boxes := make([]map[int]int, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[int]int)
		columns[i] = make(map[int]int)
		boxes[i] = make(map[int]int)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num != '.' {
				n := int(num)
				boxIndex := i/3*3 + j/3
				rows[i][n]++
				columns[j][n]++
				boxes[boxIndex][n]++

				if rows[i][n] > 1 || columns[j][n] > 1 || boxes[boxIndex][n] > 1 {
					return false
				}
			}
		}
	}
	return true
}

/**
编写一个程序，通过已填充的空格来解决数独问题。
一个数独的解法需遵循如下规则：
数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
*/
func SolveSudoku(board [][]byte) {
	//新建三个bool的数组表明行、列、3*3
	rowUsed := [9][10]bool{}
	colUsed := [9][10]bool{}
	boxUsed := [3][3][10]bool{}

	l := len(board)
	l1 := len(board[0])
	//初始化，将已经填入的数字设置为true
	for row := 0; row < l; row++ {
		for col := 0; col < l1; col++ {
			if board[row][col] != '.' {
				//这里减48是为了抵消byte转int时的值
				num := int(board[row][col]) - 48
				rowUsed[row][num] = true
				colUsed[col][num] = true
				boxUsed[row/3][col/3][num] = true
			}
		}
	}
	//递归尝试填充数组
	recusiveSolveSudoku(board, rowUsed, colUsed, boxUsed, 0, 0)
}

func recusiveSolveSudoku(board [][]byte, rowUsed, colUsed [9][10]bool, boxUsed [3][3][10]bool, row, col int) bool {
	//边界校验，如果已经完成填充，返回true，表示结束
	if col == len(board[0]) {
		col = 0
		//开始填充行
		row++
		if row == len(board) {
			return true
		}
	}

	if board[row][col] == '.' {
		//开始尝试1-9的数字
		for num := 1; num <= 9; num++ {
			//如果行，列，3*3里面的num都为false，则证明此数字可以填
			if !(rowUsed[row][num] || colUsed[col][num] || boxUsed[row/3][col/3][num]) {
				//开始填充
				rowUsed[row][num] = true
				colUsed[col][num] = true
				boxUsed[row/3][col/3][num] = true
				s := strconv.FormatInt(int64(num), 10)
				board[row][col] = s[0]

				//开始看下一个节点，如果下一个节点ok，则返回ok
				if recusiveSolveSudoku(board, rowUsed, colUsed, boxUsed, row, col+1) {
					return true
				}
				//否则将当前节点都置空，重新填充
				board[row][col] = '.'
				rowUsed[row][num] = false
				colUsed[col][num] = false
				boxUsed[row/3][col/3][num] = false
			}
		}
	} else {
		return recusiveSolveSudoku(board, rowUsed, colUsed, boxUsed, row, col+1)
	}
	return false
}

/**
「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。

这题的意思是，下一项是对当前项的描述，比如第一项是1，第二项为了描述1，则就是11，读作1个1，第三项为了描述11，则就是21，读作2个1
*/
func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}

	x := "11"

	for i := 3; i <= n; i++ {
		//解释m[i-1]
		str := x
		l := len(str)
		cur := ""

		var last byte //当前字符串的上一个字节
		var num int   //当前字符出现的次数
		for j := 0; j < l; j++ {
			if j == 0 {
				last = str[j]
				num++
			} else {
				if str[j] == last {
					num++
					if j == l-1 {
						cur += strconv.FormatInt(int64(num), 10) + string(last)
					}
				} else {
					//当前字符和上一个字符不同，则处理上一个字符
					cur += strconv.FormatInt(int64(num), 10) + string(last)
					num = 0
					last = str[j]
					j-- //处理完毕后，继续处理当前的字符
				}
			}
		}
		x = cur
	}

	return x
}

/**
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的数字可以无限制重复被选取。
说明：
所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1:
输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]

*/
func CombinationSum(candidates []int, target int) [][]int {
	//这题可以用回溯的方法做。先将数组排序，然后将目标值减去数组的每一个值，得到n个新的目标值，并判断目标值减去数组的每个值是否为0，如果是，则将其放入res
	//然后继续减每个值，如果为0，则将其放入到res中
	sort.Ints(candidates)
	res := [][]int{}
	combinationSumF(candidates, []int{}, &res, target, 0)
	return res
}

func combinationSumF(candidates, last []int, res *[][]int, target, i int) bool {
	if target == 0 {
		//当前target为0，则已经减到0了，满足，将值放入到res中
		cur := make([]int, len(last))
		copy(cur, last)
		//将cur放入res中
		*res = append(*res, cur)
		//由于后续的数组没必要循环了，所以直接return
		return true
	} else if target < 0 {
		//当前target小于0，则当前情况不符合，返回false（由于candidates是排序的，则上一级的下一次循环也没必要了）
		return false
	}
	//从当前的i开始循环
	for j := i; j < len(candidates); j++ {
		//将当前的值放入到数组中
		last = append(last, candidates[j])
		//开始递归，判断，如果返回值为false，则下面的循环都没必要了，直接中断。
		if !combinationSumF(candidates, last, res, target-candidates[j], j) {
			break
		}
		//如果为true，则将candidates[j]剔除数组，继续循环j+1
		last = last[:len(last)-1]
	}
	return true
}

/**
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用一次。
说明：
所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:
输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
*/
func CombinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	combinationSumF2(candidates, []int{}, &res, target, 0)
	return res
}

func combinationSumF2(candidates, last []int, res *[][]int, target, i int) bool {
	if target == 0 {
		//当前target为0，则已经减到0了，满足，将值放入到res中
		cur := make([]int, len(last))
		copy(cur, last)
		//将cur放入res中
		*res = append(*res, cur)
		//由于后续的数组没必要循环了，所以直接return
		return true
	} else if target < 0 {
		//当前target小于0，则当前情况不符合，返回false（由于candidates是排序的，则上一级的下一次循环也没必要了）
		return false
	}
	//从当前的i开始循环
	for j := i; j < len(candidates); j++ {
		//这里和上一题不一样，这里需要判断，当j和i属于同一层级时，且当前值等于上个值，则跳过循环
		if j > i && candidates[j] == candidates[j-1] {
			continue
		}
		//将当前的值放入到数组中
		last = append(last, candidates[j])
		//开始递归，判断，如果返回值为false，则下面的循环都没必要了，直接中断。
		//fmt.Println("递归：", last, res, target-candidates[j], j+1)
		if !combinationSumF2(candidates, last, res, target-candidates[j], j+1) {
			break
		}
		//如果为true，则将candidates[j]剔除数组，继续循环j+1
		last = last[:len(last)-1]
	}
	return true
}

/**
给定一个未排序的整数数组，找出其中没有出现的最小的正整数。

示例 1:
输入: [1,2,0]
输出: 3
示例 2:
输入: [3,4,-1,1]
输出: 2
示例 3:
输入: [7,8,9,11,12]
输出: 1
*/
func FirstMissingPositive(nums []int) int {
	/*var res = 1

	l := len(nums)
	if l == 0 {
		return res
	}

	sort.Ints(nums)
	//如果最大值是负数，那直接返回1
	if nums[len(nums)-1] <= 0 {
		return res
	}
	for k, v := range nums {
		if v == 1 {
			res = 2
			continue
		}
		if v > 1 {
			if k == 0 {
				return 1
			}
			if k >= 1 && nums[k-1] <= 0 {
				return 1
			}
			//判断和前一个值的差
			if v-nums[k-1] > 1 {
				return nums[k-1] + 1
			} else {
				res = v + 1
			}
		}
	}

	return res*/

	abs := func(n int) int {
		if n <= 0 {
			return 0 - n
		}
		return n
	}

	found1 := false
	length := len(nums)
	//这一步将nums中是负数的或者大于当前数组长度的值替换为1
	for i := 0; i < length; i++ {
		if nums[i] == 1 {
			found1 = true
		} else if nums[i] > length || nums[i] <= 0 {
			nums[i] = 1
		}
	}
	//如果没有找到1，则返回1即可
	if !found1 {
		return 1
	}
	//做替换操作，类似将nums当成一个记录（因为本地不可使用额外的内存空间）
	//取到某个值后，将这个值在nums中的位置设置为负数（标识这个值已经有了）
	for i := 0; i < length; i++ {
		num := abs(nums[i])
		if num == length {
			nums[0] = -abs(nums[0])
		} else {
			nums[num] = -abs(nums[num])
		}
	}
	//这里从第一位开始取，如果大于0，代表这个值在原数组中没有出现过，则直接返回即可
	for i := 1; i < length; i++ {
		if nums[i] > 0 {
			return i
		}
	}
	if nums[0] > 0 {
		return length
	}

	return length + 1
}

/**
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:
输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6

*/
func Trap(height []int) int {
	//思路1，取到最高那个柱子，然后从1开始循环，将每一个高度的柱子能接住的雨水计算出来，想加，就是最终面积（注意剔除最前面和最后面的部分）
	/*l := len(height)
	if l == 0 {
		return 0
	}
	var max = height[0]
	for i := 0; i < l; i++ {
		if height[i] > max {
			max = height[i]
		}
	}
	areaSum := 0
	//从第1层循环到最高一层的下一级
	for i := 1; i <= max; i++ {
		//计算第i层能接多少面积的雨水
		//找到这一层的起点和终点
		var start, end, area int
		for j := 0; j < l; j++ {
			if height[j] >= i {
				start = j
				break
			}
		}
		for j := l - 1; j > 0; j-- {
			if height[j] >= i {
				end = j
				break
			}
		}
		if end-start > 0 {
			area = end - start + 1
			for j := start; j <= end; j++ {
				if height[j] >= i {
					area--
				}
			}
		}
		areaSum += area
	}
	return areaSum*/
	//思路2 ，遍历数组，获取数组左右端的最大高度，则当前元素下的面积即为min(leftMax, rightMax) - 当前元素的高度
	/*l := len(height)
	if l == 0 || l == 1 {
		return 0
	}
	var areaSum int
	leftMax := make(map[int]float64)
	rightMax := make(map[int]float64)
	leftMax[0] = float64(height[0])
	for i := 1; i < l; i++ {
		leftMax[i] = math.Max(float64(height[i]), float64(leftMax[i-1]))
	}
	rightMax[l-1] = float64(height[l-1])
	for i := l - 2; i >= 0; i-- {
		rightMax[i] = math.Max(float64(height[i]), float64(rightMax[i+1]))
	}
	for i := 1; i < l-1; i++ {
		areaSum += int(math.Min(leftMax[i], rightMax[i])) - height[i]
	}
	return areaSum*/

	//思路3，双指针。定义两个指针指向数组的最左侧和最右侧。 判断，如果左侧小，则处理左侧，否则处理右侧。
	//以左侧为例。如果当前元素的高度大于左边的最高值，则将左边的最高值设定为当前元素的高度。否则，当前元素
	//承接的面积为左侧的最高值减去当前元素的高度，处理完后移动指针
	l := len(height)
	if l == 0 || l == 1 {
		return 0
	}
	var areaSum, leftMax, rightMax int
	left, right := 0, l-1
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				areaSum += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				areaSum += rightMax - height[right]
			}
			right--
		}
	}

	return areaSum
}

/**
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:
输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:
输入: num1 = "123", num2 = "456"
输出: "56088"
说明：
num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。

*/
func Multiply(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	result := make([]int, len1+len2)
	for i := len2 - 1; i >= 0; i-- {
		for j := len1 - 1; j >= 0; j-- {
			temp := int(num2[i]-'0')*int(num1[j]-'0') + result[i+j+1]
			if temp >= 10 {
				result[i+j] += temp / 10
				result[i+j+1] = temp % 10
				//fmt.Printf("add:%d mod:%d\n",temp / 10,temp%10)
			} else {
				result[i+j+1] = temp
			}
		}
	}
	//fmt.Println(result)
	//num1和num2较小时，去除首位0
	if result[0] == 0 {
		result = result[1:]
	}
	str := ""
	for _, v := range result {
		str += strconv.Itoa(v)
	}
	return str
}

/**
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。
两个字符串完全匹配才算匹配成功。

说明:
s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
示例 1:
输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:
输入:
s = "aa"
p = "*"
输出: true
解释: '*' 可以匹配任意字符串。
示例 3:
输入:
s = "cb"
p = "?a"
输出: false
解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
示例 4:
输入:
s = "adceb"
p = "*a*b"
输出: true
解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
示例 5:
输入:
s = "acdcb"
p = "a*c?b"
输入: false

*/
func IsMatch(s string, p string) bool {

	m := len(s)
	n := len(p)
	if p == "*" {
		return true
	}

	//设定两个游标，和'*'对应值的位置
	var i, j, iStar, jStar = 0, 0, -1, -1

	//循环第一个字符串
	for m > i {
		fmt.Println(fmt.Sprintf("循环前，i:%d, j:%d, m:%d, n:%d, iStar:%d, jStar:%d", i, j, m, n, iStar, jStar))
		//如果两个值相等
		if j < n && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
		} else if j < n && p[j] == '*' {
			//如果碰到了*，记录i值的位置和j值的位置，并将j向右移动一位，由于*可以匹配空字符串，所以i不动
			iStar = i
			jStar = j
			j++
		} else if iStar >= 0 {
			//如果此时没匹配上，则将i往后移动一位。iStar也往后移动一位，就算当前元素匹配上了。开始匹配下一个。如果下一个没匹配上，继续移动
			//如果出现了cac匹配*ab的情况，此时c是匹配的，a也是匹配的，由于最后的c和b不匹配，则将j移动回原来的a这里，然后假设cac都是匹配的*，继续排cac的下一个
			//这个方法是，如果遇到了*，那么一定会把s排完，哪怕p排不完，最后再处理p
			iStar++
			i = iStar
			j = jStar + 1
		} else {
			//如果都没匹配上，且iStar也为0，证明p中没有*，直接返回false
			return false
		}
		fmt.Println(fmt.Sprintf("循环后，i:%d, j:%d, m:%d, n:%d, iStar:%d, jStar:%d", i, j, m, n, iStar, jStar))
	}

	//如果p还没遍历完，且当前的j为*，那么将j往后移动一位
	for j < n && p[j] == '*' {
		j++
	}
	//如果j此时也移动到了末尾，那么就是true，反之则为false
	return j == n
}

/**
给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。

示例:
输入: [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
说明:
假设你总是可以到达数组的最后一个位置。
*/
func Jump(nums []int) int {
	//思路，其实可以从后往前跳。 有一个固定公式。数组的长度-1-当前下标 = 数组的值 。如果成立，则此时从该下标跳到数组的末尾，将会用到自身最大的值
	//此时，这个下标后面的值都是可以不用跳的，则到该下标为止的跳跃次数应该是最小。从后往前依次获取这个最小的下标值。
	/*var jumpTimes = 1
	l := len(nums)
	if l == 1 {
		return 0
	} else {
		var min = l - 2
		for i := l - 2; i >= 0; i-- {
			//证明从当前i跳转到l-1，需要nums[i]步。需要记录此时的下标
			if l-1-i <= nums[i] {
				min = i
			}
		}
		//如果当前min即为数组第一个元素，则证明数组只有两个值了，此时直接返回1即可
		if min == 0 {
			return 1
		} else {
			fmt.Println(min, nums[:min+1])
			//获取到了最近的min后，再从min往前取
			//return jumpTimes
			return jumpTimes + jump(nums[:min+1])
		}
	}*/

	//思路2：贪婪算法，每次跳尽量多的值。比如2,3,1,1,4,5，第一个值是2，可以跳到3或者1，那么3明显可以跳的更远，那就跳到3
	//然后3可以跳到1,1,4,那4可以跳的更远，那就跳到4，最后再跳5即可。    2, 3, 1, 1, 4, 5, 7, 10, 11, 10
	var jumpTimes = 0
	l := len(nums)
	//记录最大的位置，此位置为下标加上下标对应的值
	maxPos := 0
	//记录当前i的边界值
	end := 0
	for i := 0; i < l-1; i++ {
		//判断当前的下标和其对应的值相加得到的值是否大于当前的maxPos，如果是，则将其置为maxPos
		if nums[i]+i > maxPos {
			maxPos = nums[i] + i
		}
		//为当前的值设置边界。比如2,3,1,1,4,5 i=0时，边界此时为0+2=2。证明以i=0开始，最多往后移动2位。当从0到1,1到2之后，发现第1位的值想加最大
		//得到1+3=4。此时证明，从第1位开始往后跳3位，最大可以跳到4，此时边界就设置为4
		if i == end {
			end = maxPos
			jumpTimes++
		}
		fmt.Println(i, maxPos, end)
	}
	return jumpTimes
}

/**
给定一个没有重复数字的序列，返回其所有可能的全排列。

示例:
输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

*/
func Permute(nums []int) [][]int {
	//思路1，回溯。
	res := make([][]int, 0)
	l := len(nums)
	permuteF(nums, []int{}, -100, l, &res)
	return res
}

func permuteF(nums, cur []int, n, l int, res *[][]int) *[][]int {
	if n != -100 {
		cur = append(cur, n)
	}
	if l == len(cur) {
		*res = append(*res, cur)
		return res
	} else {
		//如果此时会走到这里，代表nums里面必然还有值
		lc := len(nums)
		for i := 0; i < lc; i++ {
			curNums := make([]int, len(nums))
			copy(curNums, nums)
			next := curNums[i]
			if lc == 1 {
				permuteF([]int{}, cur, next, l, res)
			} else {
				if i+1 < l {
					permuteF(append(curNums[:i], curNums[i+1:]...), cur, next, l, res)
				} else {
					permuteF(curNums[:i], cur, next, l, res)
				}
			}
		}
	}
	return res
}

/**
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:
输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
*/
func PermuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	l := len(nums)
	permuteUniqueF(nums, []int{}, -100, l, &res)
	return res
}

func permuteUniqueF(nums, cur []int, n, l int, res *[][]int) *[][]int {
	if n != -100 {
		cur = append(cur, n)
	}
	if l == len(cur) {
		*res = append(*res, cur)
		return res
	} else {
		//如果此时会走到这里，代表nums里面必然还有值
		lc := len(nums)
		m := make(map[int]int)
		for i := 0; i < lc; i++ {
			//判断，如果当前的值跟上一个值一致，则跳过当前值
			if _, ok := m[nums[i]]; ok {
				continue
			}
			curNums := make([]int, len(nums))
			copy(curNums, nums)
			next := curNums[i]
			//注意这里跟上面不一样，需要将当前cur重新赋值给一个切片，否则会导致结果重复
			curN := make([]int, len(cur))
			copy(curN, cur)

			if lc == 1 {
				permuteUniqueF([]int{}, curN, next, l, res)
			} else {
				if i+1 < l {
					permuteUniqueF(append(curNums[:i], curNums[i+1:]...), curN, next, l, res)
				} else {
					permuteUniqueF(curNums[:i], curN, next, l, res)
				}
			}
			m[nums[i]] = 1
		}
	}
	return res
}

/**
给定一个 n × n 的二维矩阵表示一个图像。
将图像顺时针旋转 90 度。
说明：
你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

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
func Rotate(matrix [][]int) {
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
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
示例:
输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：
所有输入均为小写字母。
不考虑答案输出的顺序。
*/
func GroupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	res := make([][]string, 0)

	//字符串排序算法
	f := func(w string) string {
		s := strings.Split(w, "")
		sort.Strings(s)
		return strings.Join(s, "")
	}

	for _, v := range strs {
		//这里将v按照ASCII码排序，判断m中是否存在，如果不存在，则新建，如果存在，则存入map对应的string切片中
		str := f(v)
		if _, ok := m[str]; ok {
			m[str] = append(m[str], v)
		} else {
			m[str] = make([]string, 0)
			m[str] = append(m[str], v)
		}
	}

	for _, v := range m {
		res = append(res, v)
	}

	return res
}

/**
实现 pow(x, n) ，即计算 x 的 n 次幂函数。

示例 1:
输入: 2.00000, 10
输出: 1024.00000
示例 2:
输入: 2.10000, 3
输出: 9.26100
示例 3:
输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25
说明:
-100.0 < x < 100.0
n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。
*/
func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1 / x
	}
	//取当前n的一半返回的值
	var half = MyPow(x, n/2)
	//如果n是偶数，rest是1，如果n是奇数，则rest是x本身。
	//以x=2 n=5为例，2的五次方换成2的2次方乘以2的2次方乘以2的1次方
	//以x=2 n=4为例，2的四次方换成2的2次方乘以2的2次方乘以1
	//以x=2 n=-5为例，2的负五次方换成2的负二次方乘以2的负二次方乘以2的负一次方
	//以x=2 n=-4为例，2的负四次方换成2的负2次方乘以2的负2次方乘以1
	var rest = MyPow(x, n%2)
	return rest * half * half
}

/**
计算那个素数妹子的微信号
*/
func GetMeizi() {
	var ji = 7140229933
	//计算乘机为上述值的两个素数
	for i := 2; i < ji/2; i++ {
		//判断此时的i是否为素数，如果是，则判断ji除以i是否为素数
		if sushu(i) {
			fmt.Println(i)
			if ji%i == 0 && sushu(ji/i) {
				fmt.Println(i, ji/i)
				break
			}
		}
	}
}

/**
判断输入的值是否为素数
*/
func sushu(x int) bool {
	var i = 2.0
	for i = 2.0; i < math.Sqrt(float64(x)); i++ {
		if x%int(i) == 0 {
			break
		}
	}
	if i <= math.Sqrt(float64(x)) {
		return false
	} else {
		return true
	}
}

/**
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
说明：
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:
输入: [2,2,1]
输出: 1
示例 2:
输入: [4,1,2,1,2]
输出: 4
*/
func SingleNumber(nums []int) int {
	//思路1，由于不使用额外的空间，那么循环数组，得到当前值之后，再循环数组，从当前值的下一位开始，判断是否能找到和其相等的元素，时间复杂度O(n²)
	/*var res int
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		flag := false
		if i == len(nums)-1 {
			return nums[i]
		} else {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] == cur {
					flag = true
					//将j从数组中去除
					if j == len(nums)-1 {
						nums = nums[:j]
					} else {
						nums = append(nums[:j], nums[j+1:]...)
					}
					break
				}
			}
		}
		if !flag {
			res = cur
			break
		}
	}
	return res*/

	//思路2   使用亦或的方法做。循环数组，由于相同的两个数亦或之后等于0（比如4为0100，4^4，得到0000。）将数组中所有的值相互亦或一遍，得到的就是没有被消除的数字
	//以4,1,2,1,2,4,5举例，0^4是4,4^1是5,5^2是7，7^1是6， 6^2是4，4^4是0，0^5是5，最后得到5
	var n int
	for _, i := range nums {
		n = n ^ i
	}
	return n

}

/**
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1:
输入: [3,2,3]
输出: 3
示例 2:
输入: [2,2,1,1,1,2,2]
输出: 2
*/
func MajorityElement(nums []int) int {
	m := make(map[int]int)
	l := len(nums)
	for _, v := range nums {
		m[v]++
		if m[v] > l/2 {
			return v
		}
	}
	return 0
}

/**
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

示例:
现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15, 20],
  [2,   5,  8, 12, 19, 25],
  [3,   6,  9, 16, 22, 27],
  [10, 13, 14, 17, 24, 28],
  [18, 21, 23, 26, 30, 31]
]
给定 target = 5，返回 true。
给定 target = 20，返回 false。
*/
func SearchMatrix(matrix [][]int, target int) bool {
	//思路：从最后一排的第一个元素开始找。 如果当前元素等于目标值，则返回true，如果当前元素小于目标值，由于当前元素已经是这一列最大的一个了，那么比目标值大的只能是其右边了，所以将
	//行指针往右移动一位。如果当前元素大于目标值，此时，上一个元素小于目标值，当前元素大于目标值，则证明目标值只有可能在当前元素的列的上面，此时将列的游标往上移动一位。

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	//设定两个游标，一个指向每个一维数组的第一个值，一个指向二维数组的最后一行
	i, j := len(matrix)-1, 0
	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			//如果当前元素大于目标值，则将行数减少一位
			i--
		} else {
			//如果当前元素小于目标值，则将列数加1。  由于一开始就已经是最后一行了，如果此时都没有大于目标值的，那么其他地方也没有了。
			j++
		}
	}
	return false
}

/**
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
说明:
初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

示例:
输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3
输出: [1,2,2,3,5,6]
*/
func Merge(nums1 []int, m int, nums2 []int, n int) {
	//思路，双指针，从后往前移动
	i := m - 1
	j := n - 1
	k := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
			k--
		} else {
			nums1[k] = nums2[j]
			k--
			j--
		}
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}

/**
你将获得 K 个鸡蛋，并可以使用一栋从 1 到 N  共有 N 层楼的建筑。
每个蛋的功能都是一样的，如果一个蛋碎了，你就不能再把它掉下去。
你知道存在楼层 F ，满足 0 <= F <= N 任何从高于 F 的楼层落下的鸡蛋都会碎，从 F 楼层或比它低的楼层落下的鸡蛋都不会破。
每次移动，你可以取一个鸡蛋（如果你有完整的鸡蛋）并把它从任一楼层 X 扔下（满足 1 <= X <= N）。
你的目标是确切地知道 F 的值是多少。
无论 F 的初始值如何，你确定 F 的值的最小移动次数是多少？

示例 1：
输入：K = 1, N = 2
输出：2
解释：
鸡蛋从 1 楼掉落。如果它碎了，我们肯定知道 F = 0 。
否则，鸡蛋从 2 楼掉落。如果它碎了，我们肯定知道 F = 1 。
如果它没碎，那么我们肯定知道 F = 2 。
因此，在最坏的情况下我们需要移动 2 次以确定 F 是多少。
示例 2：
输入：K = 2, N = 6
输出：3
示例 3：
输入：K = 3, N = 14
输出：4

提示：
1 <= K <= 100
1 <= N <= 10000
*/
func SuperEggDrop(K int, N int) int {
	T := 1
	for helper(K, T) <= N {
		T++
	}
	return T
}

func helper(K int, T int) (N int) {
	if T == 1 || K == 1 {
		return T + 1
	}
	N = helper(K-1, T-1) + helper(K, T-1)
	return
}

/**
验证回文串
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:
输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:
输入: "race a car"
输出: false
*/
func IsPalindrome1(s string) bool {
	//双指针。
	l := len(s)
	var i, j = 0, l - 1

	s = strings.ToLower(s)

	for j > i {
		//判断，如果不是数字或者字母，则跳过
		if !unicode.IsLetter(rune(s[j])) && !unicode.IsDigit(rune(s[j])) {
			j--
			continue
		}
		if !unicode.IsLetter(rune(s[i])) && !unicode.IsDigit(rune(s[i])) {
			i++
			continue
		}

		if s[i] == s[j] {
			i++
			j--
			continue
		} else {
			return false
		}

	}
	return true
}

/**
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
返回 s 所有可能的分割方案。

示例:
输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]
*/
func Partition(s string) [][]string {

	//思路，动态规划加回溯。
	//新建一个动态规划用的bool二维数组
	dp := make([][]bool, len(s)) // dp[l][r]是否是回文字符串，l,r为左右区间，帮助剪枝
	res := [][]string{}
	bytes := []byte(s)
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		for j := 0; j <= i; j++ {
			fmt.Println(i, j, bytes[i], bytes[j])
			if bytes[i] == bytes[j] && (i-j <= 2 || dp[j+1][i-1]) {
				dp[j][i] = true
			}
			fmt.Println(dp)
		}
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			fmt.Print(dp[j][i])
		}
		fmt.Println()
	}

	var track []string
	trackBack(bytes, track, dp, 0, &res)
	return res
}

func trackBack(s []byte, track []string, dp [][]bool, st int, res *[][]string) {
	if len(s) == st {
		var temp []string
		temp = append(temp, track...)
		*res = append(*res, temp)
	}

	// 切一段下来
	for i := st; i < len(s); i++ {
		if !dp[st][i] { //剪枝
			continue
		}
		//选择路径
		track = append(track, string(s[st:i+1]))
		//递归
		trackBack(s, track, dp, i+1, res)
		//撤销选择
		track = track[:len(track)-1]
	}
}

/**
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
说明：
拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。

示例 1：
输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。

示例 2：
输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。

示例 3：
输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false
*/
func WordBreak(s string, wordDict []string) bool {
	//思路：回溯。一次从s中截取字典数组中存在的单词，将剩下的s继续截取，直到最后看是否能完全截取完毕

	for _, v := range wordDict {
		if i := strings.Index(s, v); i != -1 {
			s = string(append([]byte(s[:i]), []byte(s[i+len(v)-1:])...))
			if s == "" {
				return true
			}
		}
	}

	return false
}

/**
在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。
说明:
如果题目有解，该答案即为唯一答案。
输入数组均为非空数组，且长度相同。
输入数组中的元素均为非负数。
示例 1:
输入:
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]
输出: 3
解释:
从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
因此，3 可为起始索引。

示例 2:
输入:
gas  = [2,3,4]
cost = [3,4,3]
输出: -1
解释:
你不能从 0 号或 1 号加油站出发，因为没有足够的汽油可以让你行驶到下一个加油站。
我们从 2 号加油站出发，可以获得 4 升汽油。 此时油箱有 = 0 + 4 = 4 升汽油
开往 0 号加油站，此时油箱有 4 - 3 + 2 = 3 升汽油
开往 1 号加油站，此时油箱有 3 - 3 + 3 = 3 升汽油
你无法返回 2 号加油站，因为返程需要消耗 4 升汽油，但是你的油箱只有 3 升汽油。
因此，无论怎样，你都不可能绕环路行驶一周。
*/
func CanCompleteCircuit(gas []int, cost []int) int {
	//思路，首先选择起点。能作为起点的只有gas[i] > cost[i]的节点，如非如此，前往下个节点的油必然不够。 以贪心算法为例，先找gas[i] - cost[i] 最大的节点。
	l := len(gas)
	for i := 0; i < l; i++ {
		fmt.Println("i:", i)
		//如果此时满足gas[i] > cost[i]，则当前加油站可以作为起点。  此时依次模拟，观察是否满足条件
		if gas[i] >= cost[i] {
			fmt.Println(gas[i], cost[i])
			//汽油总数
			gasNum := gas[i]
			//以i为起点模拟
			for j := i; j < l+i; j++ {
				fmt.Println("j:", j)
				//判断，如果j<l，则正常取索引，如果j>=l，则索引为j-l
				if j < l {
					fmt.Println("gasNum:", gasNum, "cost[j]:", cost[j])
					if gasNum < cost[j] {
						break
					} else {
						if j+1 == l {
							gasNum = gasNum - cost[j] + gas[0]
						} else {
							gasNum = gasNum - cost[j] + gas[j+1]
						}
					}
				} else {
					fmt.Println("gasNum:", gasNum, "cost[j-l:", cost[j-l])
					if gasNum < cost[j-l] {
						break
					} else {
						gasNum = gasNum - cost[j-l] + gas[j-l+1]
					}
				}
				//如果走到了终点，证明以i为起点是可以的，返回
				if j == l+i-1 {
					return i
				}
			}
		}
	}
	return -1
}

/**
问题1: 为什么应该将起始站点设为k+1？
因为k->k+1站耗油太大，0->k站剩余油量都是不为负的，每减少一站，就少了一些剩余油量。所以如果从k前面的站点作为起始站，剩余油量不可能冲过k+1站。
问题2: 为什么如果k+1->end全部可以正常通行，且rest>=0就可以说明车子从k+1站点出发可以开完全程？
因为，起始点将当前路径分为A、B两部分。其中，必然有(1)A部分剩余油量<0。(2)B部分剩余油量>0。
所以，无论多少个站，都可以抽象为两个站点（A、B）。(1)从B站加满油出发，(2)开往A站，车加油，(3)再开回B站的过程。
重点：B剩余的油>=A缺少的总油。必然可以推出，B剩余的油>=A站点的每个子站点缺少的油。
*/
func CanCompleteCircuit1(gas []int, cost []int) int {
	//total为所有加油站的总的油量减去总的里程。  sum为当前走过的加油站的所有油量。start为开始的加油站坐标。
	var total, sum, start int
	for i := 0; i < len(cost); i++ {
		total += gas[i] - cost[i]
		sum += gas[i] - cost[i]
		//如果在往第i个加油站走的时候，油不够了，那么就将开始的坐标设定为i+1
		if sum < 0 {
			start = i + 1
			sum = 0
		}
	}
	if total < 0 {
		return -1
	}
	return start
}

/**
编写一个算法来判断一个数是不是“快乐数”。
一个“快乐数”定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是无限循环但始终变不到 1。如果可以变为 1，那么这个数就是快乐数。

示例:
输入: 19
输出: true
解释:
1² + 9² = 82
8² + 2² = 68
6² + 8² = 100
1² + 0² + 0² = 1

*/
func IsHappy(n int) bool {
	//思路，快慢指针破解循环。   由于如果一个数是快乐数，那么最终它一定会变成1，如果一个数不是快乐数，那最终一定会是一个循环

	//这里写一个计算当前数的下一个数的函数
	bitSquareSum := func(n int) int {
		sum := 0
		for n > 0 {
			//取当前n的最后一位数
			bit := n % 10
			sum += bit * bit
			n = n / 10
		}
		return sum
	}
	//设定两个指针，初始值都等于n
	slow, fast := n, n
	//这里的思路为，如果n是快乐数，那么最后不管是快指针还是慢指针，一定会变成1，所以最后会返回1
	//如果n不是快乐数，那么它一定是个无线循环数，以一定的规律无限循环，那么快指针和慢指针一定会在某个节点相等
	//此时如果相等了切慢指针不是1，那么它就不是快乐数
	for {
		//慢指针走一步，快指针走两步
		slow = bitSquareSum(slow)
		fast = bitSquareSum(fast)
		fast = bitSquareSum(fast)
		if slow == fast {
			break
		}
	}

	return slow == 1
}

/**
371. 两整数之和
不使用运算符 + 和 - ​​​​​​​，计算两整数 ​​​​​​​a 、b ​​​​​​​之和。

示例 1:
输入: a = 1, b = 2
输出: 3
示例 2:
输入: a = -2, b = 3
输出: 1
*/
func GetSum(a int, b int) int {
	/**
	首先看十进制是如何做的： 5+7=12，三步走
	第一步：相加各位的值，不算进位，得到2。
	第二步：计算进位值，得到10. 如果这一步的进位值为0，那么第一步得到的值就是最终结果。
	第三步：重复上述两步，只是相加的值变成上述两步的得到的结果2和10，得到12。
	同样我们可以用三步走的方式计算二进制值相加： 5---101，7---111
	第一步：相加各位的值，不算进位，得到010，二进制每位相加就相当于各位做异或操作，101^111。
	第二步：计算进位值，得到1010，相当于各位进行与操作得到101，再向左移一位得到1010，(101&111)<<1。
	第三步重复上述两步，各位相加 010^1010=1000，进位值为100=(010 & 1010)<<1。
	继续重复上述两步：1000^100 = 1100，进位值为0，跳出循环，1100为最终结果。
	结束条件：进位为0，即a为最终的求和结果。
	*/
	for b != 0 {
		temp := a ^ b
		b = (a & b) << 1
		a = temp
	}
	return a
}

/**
写一个程序，输出从 1 到 n 数字的字符串表示。
1. 如果 n 是3的倍数，输出“Fizz”；
2. 如果 n 是5的倍数，输出“Buzz”；
3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。

示例：
n = 15,
返回:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]

*/
func FizzBuzz(n int) []string {
	var res []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 != 0 {
			res = append(res, "Fizz")
		} else if i%5 == 0 && i%3 != 0 {
			res = append(res, "Buzz")
		} else if i%15 == 0 {
			res = append(res, "FizzBuzz")
		} else {
			res = append(res, strconv.FormatInt(int64(i), 10))
		}
	}
	return res
}

/**
给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。

示例 1:
输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:
输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
*/
/*func maxProduct(nums []int) int {
	//思路，找整个数组中是否有0，是否有偶数个负数。
}*/

/**
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
示例 1:
输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1
示例 2:
输入: coins = [2], amount = 3
输出: -1
说明:
你可以认为每种硬币的数量是无限的。
*/
func CoinChange(coins []int, amount int) int {
	//典型的动态规划问题。用dp table记录每个值得最优解。  同时申请一个备忘录，如果当前的值出现过，则直接返回，无需再次计算（剪枝）
	m := make(map[int]int)
	return coinChangeDp(coins, amount, m)
}

func coinChangeDp(coins []int, amount int, memo map[int]int) int {
	//如果当前的金额在备忘录中，那么直接返回
	if val, ok := memo[amount]; ok {
		return val
	}
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	var res = 1<<31 - 1
	for _, v := range coins {
		//子问题为当前金额减去某个硬币的面值
		subProblem := coinChangeDp(coins, amount-v, memo)
		//如果当前的子问题返回-1（无解）那证明当前的v也是无解的
		if subProblem == -1 {
			continue
		} else {
			if res > subProblem+1 {
				res = subProblem + 1
			}
		}
	}
	//将res记录到备忘录中   注意这里即便是-1也要记录，否则会有大量重复的计算
	if res == 1<<31-1 {
		res = -1
	}
	memo[amount] = res
	return res
}

//第二种思路，用数组记录，不使用递归   这种效率高一些。 leetcode中有点奇怪，如果单独写一个min函数放在外面，效率会高很多
func coinChange(coins []int, amount int) int {
	//维护一个长度为金额+1的数组，并将数组的所有值都初始化为金额+1.
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	//循环数组，当当前的金额i大于某一个面值的硬币时，则比较i-coin（某个面值的硬币）+1和dp[i]的大小
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				if dp[i] > dp[i-coin]+1 {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}
	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}

/**
给定一个非空二叉树，返回其最大路径和。
本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
示例 1:
输入: [1,2,3]

       1
      / \
     2   3

输出: 6

示例 2:
输入: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

输出: 42
*/
var ans = -1 << 31

func MaxPathSum(root *TreeNode) int {
	//这个其实就是二叉树的后序遍历。先取当前节点的左子树的最大路径，再取右字数的最大路径。然后取其中的大的值，加上当前节点的路径，返回
	//需要注意一点就是，需要记录一个最大的结果，这个最大的结果每次更新
	maxPathSumF(root)
	return ans
}

func maxPathSumF(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := max(0, maxPathSumF(root.Left))
	right := max(0, maxPathSumF(root.Right))
	ans = max(ans, left+right+root.Val)
	return max(left, right) + root.Val
}

/**
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/
func MaxDepth(root *TreeNode) int {
	//思路，递归遍历
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := max(0, MaxDepth(root.Left))
	right := max(0, MaxDepth(root.Right))

	return max(left, right) + 1
}

/**
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

*/
var res = make([]int, 0)

func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return res
	}
	//递归法：
	InorderTraversal(root.Left)
	res = append(res, root.Val)
	InorderTraversal(root.Right)
	return res
}

/**
迭代法，栈+颜色标记节点  由于一般的栈方法比较难理解，下面是颜色标记法，非常便于理解，同时便于转换不同的顺序，
对中序，前序，后序可以写出结构一致的代码
*/
func InorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return res
	}
	//一个结构体，里面有节点和节点的颜色，节点的颜色分为两种，1是处理过的，0是未处理过的
	type Node struct {
		Color    int
		TreeNode *TreeNode
	}
	//新建一个栈
	var stack = make([]*Node, 0)
	//假定有一个栈。 存储需要遍历的节点。
	stack = append(stack, &Node{0, root})
	for len(stack) > 0 {
		//出栈操作
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur == nil {
			continue
		}
		//假如cur未处理过，则正常处理
		if cur.Color == 0 {
			cur.Color = 1
			//中序，左中右的顺序，右先入，中再入，左再入
			if cur.TreeNode.Right != nil {
				stack = append(stack, &Node{0, cur.TreeNode.Right})
			}
			stack = append(stack, cur)
			if cur.TreeNode.Left != nil {
				stack = append(stack, &Node{0, cur.TreeNode.Left})
			}
		} else {
			res = append(res, cur.TreeNode.Val)
		}
	}
	return res
}

/**
给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。

示例:
输入: 3
输出:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
解释:
以上的输出对应以下 5 种不同结构的二叉搜索树：

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

*/

func GenerateTrees(n int) []*TreeNode {
	if n == 0 {
		res := make([]*TreeNode, 0)
		return res
	}
	//1、回溯法。我们从1到n之间选取一个点i作为根节点，那么由于二叉树的性质，1到i-1将会组成这个二叉树的左子树，i+1会组成这个二叉树的右子树。递归完成之后再将两个子树组合即可
	return generateTreesF(1, n)
}

func generateTreesF(start, end int) []*TreeNode {
	var res = make([]*TreeNode, 0)

	if start > end {
		res = append(res, nil)
		return res
	}

	for i := start; i <= end; i++ {
		//获取左子树
		leftTree := generateTreesF(start, i-1)
		//获取右子树
		rightTree := generateTreesF(i+1, end)
		//将左右两个子树结合
		for _, lv := range leftTree {
			for _, rv := range rightTree {
				root := &TreeNode{Val: i}
				root.Left = lv
				root.Right = rv
				res = append(res, root)
			}
		}
	}
	return res
}

/**
动态规划法。所有的回溯递归都可以转换为动态规划的思路
*/
func GenerateTreesD(n int) []*TreeNode {
	//动态规划dp数组
	dp := make([][]*TreeNode, n+1)
	dp[0] = make([]*TreeNode, 0)

	if n == 0 {
		return dp[0]
	}
	dp[0] = append(dp[0], nil)
	//循环，长度从1到n，dp[len]代表长度为len的情况下有多少种树的组合情况。那么依照动态规划的思想
	//求长度为n的情况，只需要将从1到n分别作为根节点，然后计算出左子树的长度和右子树的长度，然后合并即可
	for len := 1; len <= n; len++ {
		dp[len] = make([]*TreeNode, 0)
		//这里将从1到len的不同的数字作为根节点
		for root := 1; root <= len; root++ {
			left := root - 1    //左子树的长度
			right := len - root //右子树的长度
			for _, lv := range dp[left] {
				for _, rv := range dp[right] {
					treeRoot := &TreeNode{Val: root}
					treeRoot.Left = lv
					//克隆右子树并加上偏差值
					treeRoot.Right = generateTreesClone(rv, root)
					dp[len] = append(dp[len], treeRoot)
				}
			}
		}
	}
	return dp[n]
}

//克隆树，并将树做一定的偏移操作
func generateTreesClone(tree *TreeNode, offset int) *TreeNode {
	if tree == nil {
		return nil
	}
	node := &TreeNode{Val: tree.Val + offset}           //偏移根节点
	node.Left = generateTreesClone(tree.Left, offset)   //偏移左节点
	node.Right = generateTreesClone(tree.Right, offset) //偏移右节点
	return node
}

/**
96. 不同的二叉搜索树
给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

示例:
输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/
func NumTrees(n int) int {
	//直接用动态规划的思路来解
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	//状态转移方程。
	for len := 2; len <= n; len++ {
		cur := 0
		//以每个长度作为根节点
		for root := 1; root <= len; root++ {
			left := root - 1    //左子树数量
			right := len - root //右子树数量

			//如果当前左子树为0，那么直接加上右子树的数量
			if left == 0 {
				cur += dp[right]
			} else if right == 0 {
				//如果当前右子树为0，那么直接加上左子树的数量
				cur += dp[left]
			} else {
				//否则左右子树交替组合
				cur += dp[left] * dp[right]
			}
		}
		dp[len] = cur
	}
	return dp[n]
}

/**
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
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
func isValidBST(root *TreeNode) bool {
	//前序遍历
	return isValidBSTF(root, -1<<63, 1<<63-1)
}

func isValidBSTF(root *TreeNode, min, max int) bool {
	return root == nil || min < root.Val && root.Val < max &&
		isValidBSTF(root.Left, min, root.Val) &&
		isValidBSTF(root.Right, root.Val, max)
}

var lastIsValidBST = -1 << 31

func IsValidBST(root *TreeNode) bool {
	//中序遍历，如果要满足左子节点<当前节点<右子节点，那么只需要中序遍历，依次遍历左中右，记录每次的值，
	//但凡有一次不满足上面的大于条件，则即为false
	if root == nil {
		return true
	}
	if !IsValidBST(root.Left) {
		return false
	}
	if lastIsValidBST >= root.Val {
		return false
	}
	lastIsValidBST = root.Val
	return IsValidBST(root.Right)
}

/**
二叉搜索树中的两个节点被错误地交换。
请在不改变其结构的情况下，恢复这棵树。

示例 1:
输入: [1,3,null,null,2]

   1
  /
 3
  \
   2

输出: [3,1,null,null,2]

   3
  /
 1
  \
   2

示例 2:
输入: [3,1,4,null,null,2]

  3
 / \
1   4
   /
  2

输出: [2,1,4,null,null,3]

  2
 / \
1   4
   /
  3

*/
var last, first, second *TreeNode

func RecoverTree(root *TreeNode) {
	//设定三个树结构，分别代表最后遍历的一个树，要交换的第一个树和第二个树，先都赋值为nil
	last, first, second = nil, nil, nil
	//中序遍历
	RecoverTreeDfs(root)
	first.Val, second.Val = second.Val, first.Val
}

func RecoverTreeDfs(root *TreeNode) {
	if root == nil {
		return
	}
	RecoverTreeDfs(root.Left)
	//由于中序遍历正常的二叉搜索树应该为升序，如果上一个节点不为nil，且上一个节点大于当前节点的值，那么证明上个节点跟当前节点的位置是错的
	if last != nil && root.Val <= last.Val {
		//如果第一个节点是nil，则将上个节点和当前节点分别赋值给first和second
		if first == nil {
			first, second = last, root
		} else {
			//如果第一个节点不为空，那么证明first和second之间还有其他的正常节点（如3,2,1这种情况）。此时将second置为当前节点
			second = root
			return //剪枝
		}
	}
	last = root
	RecoverTreeDfs(root.Right)
}

/**
给定两个二叉树，编写一个函数来检验它们是否相同。
如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

示例 1:
输入:       1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

输出: true

示例 2:
输入:      1          1
          /           \
         2             2

        [1,2],     [1,null,2]
输出: false

示例 3:
输入:       1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

输出: false

*/
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	//二叉树前序遍历
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	if !IsSameTree(p.Left, q.Left) {
		return false
	}
	if !IsSameTree(p.Right, q.Right) {
		return false
	}
	return true
}

/**
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3

说明:
如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricF(root.Left, root.Right)
}

func isSymmetricF(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	//注意这里是传的左树的左子树和右树的右子树，以及左树的右子树和右树的左子树
	return isSymmetricF(p.Left, q.Right) && isSymmetricF(p.Right, q.Left)
}

/**
迭代的思路来做
*/
func IsSymmetricA(root *TreeNode) bool {
	if root == nil {
		return true
	}
	//这里准备一个队列
	queue := make([]*TreeNode, 0)
	//将根节点的左子树和右子树放入队列中
	queue = append(queue, root.Left, root.Right)
	for len(queue) > 0 {
		//从头部弹出两个元素
		p := queue[0]
		q := queue[1]
		queue = queue[2:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		//将p的左子树和q的右子树入队列
		queue = append(queue, p.Left, q.Right)
		//将p的右子树和q的左子树入队列
		queue = append(queue, p.Right, q.Left)
	}
	return true
}

/**
给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
*/

func LevelOrder(root *TreeNode) [][]int {
	//用队列的思路来做，先进先出
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	type TREE struct {
		treeNode *TreeNode
		level    int
	}

	queue := make([]*TREE, 0)
	queue = append(queue, &TREE{treeNode: root, level: 0})
	//当前level结果
	cur := make([]int, 0)
	//上一个级别
	lastLevel := 0

	for len(queue) > 0 {
		//从队列中取出第一个元素
		p := queue[0]
		queue = queue[1:]
		//判断，如果这个元素的深度跟lastLevel一样，则直接将其放入到cur的数组中
		if p.level == lastLevel {
			if p.treeNode != nil {
				cur = append(cur, p.treeNode.Val)
			}
		} else {
			//将上个级别的节点放入结果中，并新建当前level的数组，将当前节点放入到当前level的数组中
			res = append(res, cur)
			cur = make([]int, 0)
			if p.treeNode != nil {
				cur = append(cur, p.treeNode.Val)
			}
		}

		//将当前节点的左右节点放入到队列中,并将这些节点的level在当前节点的基础上+1
		if p.treeNode != nil {
			queue = append(queue, &TREE{treeNode: p.treeNode.Left, level: p.level + 1}, &TREE{treeNode: p.treeNode.Right, level: p.level + 1})
		}
		lastLevel = p.level
	}

	return res
}

/**
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
  [15,7]
]

*/
var zigzagRes [][]int

func ZigzagLevelOrder(root *TreeNode) [][]int {
	//为什么这么写是因为leecode如果直接在定义的地方make，会把结果都聚合起来
	zigzagRes = make([][]int, 0)
	zigzagLevelOrderF(root, 0)
	return zigzagRes
}

func zigzagLevelOrderF(tree *TreeNode, level int) {
	if tree == nil {
		return
	}
	if level == len(zigzagRes) {
		t := make([]int, 0)
		zigzagRes = append(zigzagRes, t)
	}
	//如果当前的level是偶数，则正序放入，是奇数则逆序
	if level%2 == 0 {
		zigzagRes[level] = append(zigzagRes[level], tree.Val)
	} else {
		temp := make([]int, 0)
		temp = append(temp, tree.Val)
		temp = append(temp, zigzagRes[level]...)
		zigzagRes[level] = temp
	}
	zigzagLevelOrderF(tree.Left, level+1)
	zigzagLevelOrderF(tree.Right, level+1)
}

//bts(广度优先解法)
func zigzagLevelOrderBts(root *TreeNode) [][]int {
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	if root == nil {
		return res
	}
	//层级
	level := 0
	queue = append(queue, root)
	for len(queue) > 0 {
		temp := make([]int, 0)
		//这个count表示当前级别下在队列中有多少个元素，将这些元素全部从队列中弹出来并且处理
		count := len(queue)
		//循环处理当前level下的所有子元素
		for count > 0 {
			//获取第一个节点
			cur := queue[0]
			queue = queue[1:]
			//如果层级是偶数，则为正序插入，否则为逆序插入
			if level%2 == 0 {
				temp = append(temp, cur.Val)
			} else {
				t := make([]int, 0)
				t = append(t, cur.Val)
				temp = append(t, temp...)
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			count--
		}
		level++
		res = append(res, temp)
	}
	return res
}

/**
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7

*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	//思路：前序遍历数组的第一个值一定是树的根节点。  然后找到这个值在中序遍历中的位置，这个值在中序遍历数组中左侧为树的左子树，右侧为树的右子树
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	res := &TreeNode{Val: preorder[0]}
	//找到当前根节点在中序遍历中的位置
	/**
	我们在 inorder 中找到 mid 为根节点的下标
	由于中序遍历特性，mid 左侧都为左子树节点，所以左子树的节点有 mid 个
	那么同样的，由于前序遍历的特性，preorder 第一个元素（根节点）后跟着的就是它的左子树节点，一共有 mid 个，所以切了 [1:mid+1] 出来
	*/
	var mid int
	for k, v := range inorder {
		if v == preorder[0] {
			mid = k
			break
		}
	}
	//生成res的左子树
	res.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	//生成res的右子树
	res.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return res
}

/**
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出
中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

*/
func buildTreeB(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	//后序遍历的最后一个值一定是根节点
	res := &TreeNode{Val: postorder[len(postorder)-1]}
	//找到这个值的左子树跟右子树
	var mid int
	for k, v := range inorder {
		if v == postorder[len(postorder)-1] {
			mid = k
			break
		}
	}

	res.Left = buildTreeB(inorder[:mid], postorder[:mid])
	res.Right = buildTreeB(inorder[mid+1:], postorder[mid:len(postorder)-1])
	return res
}

/**
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]

*/
var levelOrderBottomRes [][]int

func LevelOrderBottom(root *TreeNode) [][]int {
	//思路1，递归，自上向下层次遍历之后翻转结果数组
	levelOrderBottomRes = make([][]int, 0)
	levelOrderBottomF(root, 0)
	//翻转res
	l := len(levelOrderBottomRes)
	res := make([][]int, l)
	for i := l - 1; i >= 0; i-- {
		res[l-i-1] = levelOrderBottomRes[i]
	}
	return res
}

func levelOrderBottomF(root *TreeNode, level int) {
	if root == nil {
		return
	}
	//长度达到了结果的最大值，需要延伸一位
	if level == len(levelOrderBottomRes) {
		t := make([]int, 0)
		levelOrderBottomRes = append(levelOrderBottomRes, t)
	}
	levelOrderBottomRes[level] = append(levelOrderBottomRes[level], root.Val)
	levelOrderBottomF(root.Left, level+1)
	levelOrderBottomF(root.Right, level+1)
}

/**
广度优先的方式
*/
func LevelOrderBottomS(root *TreeNode) [][]int {
	//思路2，广度优先。在插入到结果集的时候采用从前面插入的方式
	levelOrderBottomRes := make([][]int, 0)
	if root == nil {
		return levelOrderBottomRes
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		//处理当前层级的元素
		cur := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				cur = append(cur, node.Val)
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
		}
		//将cur插入到结果的顶部
		levelOrderBottomRes = append([][]int{cur}, levelOrderBottomRes...)
	}

	return levelOrderBottomRes
}

/**
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定有序数组: [-10,-3,0,5,9],
一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5

*/
func SortedArrayToBST(nums []int) *TreeNode {
	//采用递归的思路。判断nums的长度是奇数还是偶数，确定当前子树的根节点，确定了根节点，那么左子树跟右子树就确定了
	l := len(nums)
	if l == 0 {
		return nil
	}
	var mid int
	mid = l / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = SortedArrayToBST(nums[:mid])
	node.Right = SortedArrayToBST(nums[mid+1:])
	return node
}

/**
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定的有序链表： [-10, -3, 0, 5, 9],
一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5

*/
func sortedListToBST(head *ListNode) *TreeNode {
	//思路1，先将链表转换为数组，再根据上面一体的做法分别处理左子树跟右子树。
	/*var cur = make([]int, 0)
	for head != nil {
		cur = append(cur, head.Val)
		head = head.Next
	}
	l := len(cur)
	if l == 0 {
		return nil
	}
	mid := l / 2
	node := &TreeNode{Val: cur[mid]}
	node.Left = SortedArrayToBST(cur[:mid])
	node.Right = SortedArrayToBST(cur[mid+1:])
	return node*/

	//思路2，用快慢指针，找到链表的中间点，然后将其设置为当前的根节点，将前半部分作为左节点，后半部分作为右节点递归处理
	if head == nil {
		return nil
	}
	var pre *ListNode
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	root := new(TreeNode)
	root.Val = slow.Val
	//(1)如果只有一个点
	if pre == nil {
		return root
	}

	//(2)如果大于一个点，则需要把链表切一次，丢弃中间点
	next := slow.Next
	pre.Next = nil

	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(next)

	return root
}

/**
给定一个二叉树，判断它是否是高度平衡的二叉树。
本题中，一棵高度平衡二叉树定义为：
一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

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
func IsBalanced(root *TreeNode) bool {
	//思路，获取当前节点的左子树的深度和右子树的深度
	res, _ := getDeep(root)
	return res
}

func getDeep(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	//获取左子树的深度和结果
	lr, ld := getDeep(root.Left)
	if !lr {
		return false, 0
	}
	//获取右子树的深度和结果
	rr, rd := getDeep(root.Right)
	if !rr {
		return false, 0
	}
	if ld > rd && ld-rd > 1 || rd > ld && rd-ld > 1 {
		return false, 0
	}
	if ld >= rd {
		return true, ld + 1
	} else {
		return true, rd + 1
	}
}

/**
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明: 叶子节点是指没有子节点的节点。

示例:
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.

*/
func MinDepth(root *TreeNode) int {
	//思路1：递归，获取左子树的最小深度和右子树的最小深度，然后取最小值加1
	if root == nil {
		return 0
	}
	//判断当前节点是否为叶子节点（没有子节点的节点）
	var lM, rM int
	if root.Left != nil && root.Right == nil {
		lM = MinDepth(root.Left)
		rM = 1<<31 - 1
	} else if root.Right != nil && root.Left == nil {
		rM = MinDepth(root.Right)
		lM = 1<<31 - 1
	} else {
		lM = MinDepth(root.Left)
		rM = MinDepth(root.Right)
	}

	if lM >= rM {
		return rM + 1
	} else {
		return lM + 1
	}
}

func MinDepthF(root *TreeNode) int {
	//思路2：深度优先
	if root == nil {
		return 0
	}
	var min int
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		min++
		for i := 0; i < l; i++ {
			//如果当前节点的左右节点为nil(当前节点是叶子节点)
			if queue[i].Left == nil && queue[i].Right == nil {
				return min
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]
	}
	return min
}

/**
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。

*/
func HasPathSum(root *TreeNode, sum int) bool {
	//思路：和减去当前的节点，如果为0了且没有其他子节点，则为true，否则为false
	if root == nil {
		return false
	}
	cur := sum - root.Val
	if root.Left == nil && root.Right == nil {
		if cur == 0 {
			return true
		} else {
			return false
		}
	} else {
		return HasPathSum(root.Left, cur) || HasPathSum(root.Right, cur)
	}
}

/**
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
说明: 叶子节点是指没有子节点的节点。

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

[
   [5,4,11,2],
   [5,8,4,5]
]

*/
var pathSumRes [][]int

func PathSum(root *TreeNode, sum int) [][]int {
	//思路1 回溯
	pathSumRes = make([][]int, 0)
	cur := make([]int, 0)
	pathSumF(root, sum, cur)
	return pathSumRes
}

func pathSumF(root *TreeNode, sum int, cur []int) {
	//判断，如果当前节点没有左右节点，且节点值刚好等于sum，则将当前节点放入到cur中，将cur放入res中
	if root == nil {
		return
	}
	cur = append(cur, root.Val)
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			curN := make([]int, len(cur))
			copy(curN, cur)
			pathSumRes = append(pathSumRes, curN)
		}
	}
	pathSumF(root.Left, sum-root.Val, cur)
	pathSumF(root.Right, sum-root.Val, cur)
}

/**
给定一个二叉树，原地将它展开为链表。
例如，给定二叉树

    1
   / \
  2   5
 / \   \
3   4   6
将其展开为：

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
*/
var lastTree *TreeNode

func flatten(root *TreeNode) {
	lastTree = nil
	flattenF(root)
}

func flattenF(root *TreeNode) {
	if root == nil {
		return
	}
	//递归。不管后面的节点，处理当前节点时，只需要将左侧置空，右侧正常处理即可
	flattenF(root.Right)
	flattenF(root.Left)
	root.Right = lastTree
	root.Left = nil
	lastTree = root
}

func Unzip(s string) string {
	start := strings.Index(s, "[")
	if start == -1 {
		return s
	}
	end := strings.LastIndex(s, "]")
	var temp string
	temp = s[start+1 : end]
	nStart := strings.Index(temp, "|")
	n, _ := strconv.Atoi(temp[:nStart])
	//判断arr[1]中是否有[
	flag := strings.Index(temp[nStart+1:], "[")
	f := temp[nStart+1:]
	if flag != -1 {
		f = Unzip(f)
	}
	var res string
	for i := 1; i <= n; i++ {
		res += f
	}
	if end == len(s)-1 {
		res = s[:start] + res
	} else {
		res = s[:start] + res + s[end+1:]
	}
	return res
}

//翻转链表 1就地翻转
func ReverseList1(node *ListNode) *ListNode {
	//思路，定义一个空链表，遍历传入的链表，每次断开一个元素，并将这个元素拼在定义的那个链表的头部
	if node == nil {
		return nil
	}
	if node.Next == nil {
		return node
	}

	var p *ListNode
	var tmp *ListNode
	for node != nil {
		//将当前节点的后面部分全部放入一个临时变量汇总
		tmp = node.Next
		//将当前节点的Next指向p
		node.Next = p
		//将p置为当前节点的链表
		p = node
		//遍历下一个节点
		node = tmp
	}
	//printList(p)
	return p
}

//翻转链表2，借助栈先进后出的特性
func ReverseList2(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	if node.Next == nil {
		return node
	}
	stack := make([]*ListNode, 0)
	var tmp *ListNode
	for node != nil {
		tmp = node.Next
		node.Next = nil
		stack = append(stack, node)
		node = tmp
	}

	var head = &ListNode{Val: 0, Next: nil}
	var p = head
	for i := len(stack) - 1; i >= 0; i-- {
		p.Next = stack[i]
		p = p.Next
	}
	//printList(head.Next)
	return head.Next
}

/**
两个函数，实现对二叉树的序列化和反序列化
思路1：dfs深度优先 （递归）
*/
func DfsSerializeTree(node *TreeNode) string {
	//按照前序的方式遍历，如果遇到空节点，则存入'$'
	if node == nil {
		return "$"
	} else {
		return strconv.Itoa(node.Val) + "," + DfsSerializeTree(node.Left) + "," + DfsSerializeTree(node.Right)
	}
}

var unSerializeTreeRes []string

func DfsUnSerializeTree(steam string) *TreeNode {
	unSerializeTreeRes = strings.Split(steam, ",")
	return dfsUnSerializeTree()
}

func dfsUnSerializeTree() *TreeNode {
	tmp := unSerializeTreeRes[0]
	unSerializeTreeRes = unSerializeTreeRes[1:]
	if tmp == "$" {
		return nil
	}
	v, _ := strconv.Atoi(tmp)
	root := &TreeNode{Val: v}
	root.Left = dfsUnSerializeTree()
	root.Right = dfsUnSerializeTree()
	return root
}

/**
思路2：广度优先  （二叉树的层次遍历）
*/
func BfsSerializeTree(node *TreeNode) string {
	if node == nil {
		return ""
	}
	//声明一个队列，将二叉树按照层次遍历的方式放入队列中
	var res = make([]string, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, node)
	for len(queue) > 0 {
		cur := queue[0]
		//将当前节点的值放入res中，然后将左右子节点放入队列中
		if cur != nil {
			res = append(res, strconv.Itoa(cur.Val))
			queue = append(queue, cur.Left, cur.Right)
		} else {
			res = append(res, "$")
		}
		queue = queue[1:]
	}
	return strings.Join(res, ",")
}

func BfsUnSerializeTree(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	str := strings.Split(data, ",")
	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(str[0])
	queue := []*TreeNode{root}
	//这里不会溢出，因为哪怕二叉树只有一个节点，传入的data也应该是“1,$,$”这种形式
	str = str[1:]
	for len(queue) > 0 {
		cur := queue[0]
		//生成当前树的左子树，并将其放入队列中
		if str[0] != "$" {
			leftVal, _ := strconv.Atoi(str[0])
			cur.Left = &TreeNode{Val: leftVal}
			queue = append(queue, cur.Left)
		}
		//生成当前树的右子树，并将其放入队列中
		if str[1] != "$" {
			rightVal, _ := strconv.Atoi(str[1])
			cur.Right = &TreeNode{Val: rightVal}
			queue = append(queue, cur.Right)
		}
		queue = queue[1:]
		str = str[2:]
	}
	return root
}

//连续子数组的最大和
func GetMaxSumOfChildArr(param []int) int {
	l := len(param)
	if l == 0 {
		return 0
	}
	//用一个数组记录到i为止的数列和。
	dp := make([]int, l)
	dp[0] = param[0]
	var max = dp[0]
	for i := 1; i < l; i++ {
		//如果前一个值小于0，则到i为止的最大数即为它自身
		if dp[i-1] <= 0 {
			dp[i] = param[i]
		} else {
			//如果前一个值大于0，则到i为止的最大数为它自身加上前一个数
			dp[i] = param[i] + dp[i-1]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	//fmt.Println(dp)
	return max
}

//礼物的最大价值。在一个m*n的棋盘的每一格都放有一个礼物，可以从棋盘的左上角开始拿，每次可以往左或者往右移动一格。 请问能拿到的礼物的最大价值和
func MaxValOfGift(param [][]int) int {
	//典型的动态规划算法。 dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + param[i][j]
	row := len(param)
	if row == 0 {
		return 0
	}
	col := len(param[0])
	if col == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a >= b {
			return a
		} else {
			return b
		}
	}
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = param[0][0]
			} else {
				if i == 0 {
					dp[i][j] = dp[i][j-1] + param[i][j]
				} else if j == 0 {
					dp[i][j] = dp[i-1][j] + param[i][j]
				} else {
					dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + param[i][j]
				}
			}
		}
	}
	return dp[row-1][col-1]
}

//找出最长不含重复字符串的子字符串 例如"arabcacfr" 最长的不含重复字符的子字符串是"acfr"，长度是4
func MaxUnRepeatStr(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	//动态规划。
	dp := make([]int, l)
	m := make(map[byte]int) //用来保存当前的字符的位置
	dp[0] = 1
	m[s[0]] = 0
	max := 1
	for i := 1; i < l; i++ {
		//如果这个值在之前出现过。那么获取这个值的位置。判断当前的i到这个值的最新位置差值跟dp[i-1]谁大
		if v, ok := m[s[i]]; ok {
			if i-v > dp[i-1] {
				dp[i] = dp[i-1] + 1
			} else if i-v == dp[i-1] {
				dp[i] = dp[i-1]
			} else {
				//这里要注意，当i-v小于dp[i]时，需要从上一个重复元素的下一位开始计算，加到当前位
				dp[i] = i - v
			}
		} else {
			//如果这个值没有出现过，那么就直接在dp[i-1]的基础上加1
			dp[i] = dp[i-1] + 1
		}
		if max < dp[i] {
			max = dp[i]
		}
		m[s[i]] = i
	}
	return max
}

/**
和为s的连续整数序列。  输入一个整数S，打印出所有和为s的连续整数序列（至少含有两个数）。例如，输入14，由于1+2+3+4+5=4+5+6=7+8=15，所以打印出{1,2,3,4,5} {4,5,6} {7,8}
*/
func FindContinueSequence(sum int) {
	//设定两个指针，一个指向1，一个指向2。两个指针组成一个连续序列{1,2} 如果这个连续序列相加小于sum，则将指针2往后移动。如果这个连续序列相加大于sum，则将指针1往后移动
	first := 1
	second := 2
	middle := (sum + 1) / 2
	curSum := first + second
	for first < middle {
		if curSum == sum {
			fmt.Println(first, second)
		}

		for curSum > sum && first < middle {
			curSum -= first
			first++
			if curSum == sum {
				fmt.Println(first, second)
			}
		}

		second++
		curSum += second
	}
}

/**
 圆圈中最后剩下的数字。
0,1，....，n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字，求出这个圆圈里剩下的最后一个数字。
*/
func LastRemaining(m, n int) int {
	/**
	直接通过数学规律总结约瑟夫环。在0~n-1这个数列中，删除第m个数字，那么被删除的数字一定是(m-1) % n （这里余n是考虑m大于n的情况 ）我们假设这个数字的位置是k，那么删除k之后，剩下的元素还有0,1，...，k-1，k+1，...，n-1，并且下一次删除送数字k+1开始。 相当于在下一次的数列中，k+1是排在开头的，即下一次的顺序为：k+1, k+2,...，n-1, 0, 1, ..., k-1。
	那么其实有一个映射关系，将这个数列映射到0~n-2（此时已经弹出了一个元素）
	k+1 -> 0
	k+2 -> 1
	...
	n-1 -> n-k-2
	0 -> n-k-1
	1 -> n-k
	...
	k-1 -> n-2
	可以看出，这就是原问题中把n替换成n-1的情况，设最终胜利的那个人在这种编号环境里（已经出列一个元素，编号范围为0~n-2）的编号为x，则我们可以求出这个人在原编号环境（初始编号范围 0~n-1）下的编号（x+k）%n。
	如果我们用f(n)标识n个人的情况下最终结果的编号，那么如何知道f(n-1)呢？ 答案是由f(n-2)得来，这就转换成典型的递归问题。
	f(1) = 0   (当只有最后一个人的时候，无论m为几，最终结果都为0)
	f(n) = (f(n-1) + m) % n
	如果此时要求f(n)，那么只需要从f(1)推算即可。
	*/
	if m < 1 || n < 1 {
		return -1
	}
	last := 0 //此时的n为1
	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}

/**
N叉树的前序遍历
给定一个 N 叉树，返回其节点值的前序遍历。

例如，给定一个 3叉树 :
返回其前序遍历: [1,3,5,6,2,4]。
说明: 递归法很简单，你可以使用迭代法完成此题吗?
 Definition for a Node.
  type Node struct {
      Val int
      Children []*Node
  }
*/

/*var preorderRes []int

//递归法
func preorder(root *Node) []int {
	preorderRes = make([]int, 0)
	var preorderF func(root *Node)
	preorderF = func(root *Node) {
		if root == nil {
			return
		}
		preorderRes = append(preorderRes, root.Val)
		for _, v := range root.Children {
			preorderF(v)
		}
	}
	preorderF(root)
	return preorderRes
}*/

//迭代法
/**
func preorder(root *Node) []int {
	//迭代法，dps
    res := make([]int, 0)
    if root == nil {
        return res
    }
    queue := make([]*Node, 0)
    queue = append(queue, root)
    for len(queue) > 0 {
        //取出最前面一个元素，然后判断
        cur := queue[0]
        queue = queue[1:]
        if cur != nil {
            res = append(res, cur.Val)
            temp := make([]*Node, 0)
            for _, v := range cur.Children {
                temp = append(temp, v)
            }
            queue = append(temp, queue...)
        }
    }
    return res
}
*/

/**
最大树定义：一个树，其中每个节点的值都大于其子树中的任何其他值。

给出最大树的根节点 root。

给出一个树，给一个值。如果这个值大于这个树的根，则将这个树作为当前值的左子树。否则将这个值放入这个树的右子树中
*/
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if val > root.Val {
		return &TreeNode{val, root, nil}
	}
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}

/**
给你一个树，请你 按中序遍历 重新排列树，使树中最左边的结点现在是树的根，并且每个结点没有左子结点，只有一个右子结点。
示例 ：

输入：[5,3,6,2,4,null,8,1,null,null,null,7,9]

       5
      / \
    3    6
   / \    \
  2   4    8
 /        / \
1        7   9

输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]

 1
  \
   2
    \
     3
      \
       4
        \
         5
          \
           6
            \
             7
              \
               8
                \
                 9
*/
var increasingBSTParent *TreeNode
var increasingBSTP *TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	increasingBSTParent = &TreeNode{}
	increasingBSTP = increasingBSTParent
	increasingBSTF(root)
	return increasingBSTParent.Right
}

func increasingBSTF(root *TreeNode) {
	//先中序遍历，得到的每个值放入到increasingBSTRes中
	if root == nil {
		return
	}
	increasingBSTF(root.Left)
	increasingBSTP.Right = &TreeNode{Val: root.Val}
	increasingBSTP = increasingBSTP.Right
	increasingBSTF(root.Right)
}

/**
222. 完全二叉树的节点个数
给出一个完全二叉树，求出该树的节点个数。
说明：
完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。
示例:
输入:
    1
   / \
  2   3
 / \  /
4  5 6

输出: 6

*/
func countNodes(root *TreeNode) int {
	//思路，递归获取左右子树的高度left和right。如果left==right，则证明左右子树等高，所以左子树已经被填满了，此时左子树的个数为pow(2, left)
	//此时的总数量等于pow(2, left) + countNodes(root.Right) 如果left != right，则证明左子树比右子树高，则此时右子树比左子树少一层，但也是填满的，所以右子树的个数为pow(2, right)
	//则此时总数量等于pow(2. right) + countNodes(root.Left)
	if root == nil {
		return 0
	}
	left := countNodesGetTreeDeep(root.Left)
	right := countNodesGetTreeDeep(root.Right)
	if left == right {
		return countNodes(root.Right) + (1 << left)
	} else {
		return countNodes(root.Left) + (1 << right)
	}
}

func countNodesGetTreeDeep(root *TreeNode) uint {
	var res uint
	for root != nil {
		res += 1
		root = root.Left
	}
	return res
}

/**
二叉树剪枝
给定二叉树根结点 root ，此外树的每个结点的值要么是 0，要么是 1。
返回移除了所有不包含 1 的子树的原二叉树。
( 节点 X 的子树为 X 本身，以及所有 X 的后代。)
*/
func pruneTree(root *TreeNode) *TreeNode {
	//其实需要被剪去的节点只符合一种情况。本身是0,且没有左右子节点。 自下向上递归，如果遇到当前节点的左右节点都被剪掉，且当前节点是0，则当前节点也应该被剪去
	if root == nil {
		return nil
	}
	//如果当前节点是0
	if root.Val == 0 {
		//递归处理左右的节点
		root.Left = pruneTree(root.Left)
		root.Right = pruneTree(root.Right)
		//如果左右的节点都是nil，并且当前节点是0,则证明当前节点需要被剪去
		if root.Left == nil && root.Right == nil {
			return nil
		}
	} else {
		root.Left = pruneTree(root.Left)
		root.Right = pruneTree(root.Right)
	}
	return root
}

/**
给定一个根为 root 的二叉树，每个结点的深度是它到根的最短距离。
如果一个结点在整个树的任意结点之间具有最大的深度，则该结点是最深的。
一个结点的子树是该结点加上它的所有后代的集合。
返回能满足“以该结点为根的子树中包含所有最深的结点”这一条件的具有最大深度的结点。

*/
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	//其实这题就是个简单的最深节点的最近公共祖先
	//递归，如果当前左子树和右子树深度一样，则返回root。如果左子树比右子树深，则递归处理左子树，否则递归处理右子树
	if root == nil {
		return root
	}
	l := MaxDepth(root.Left)
	r := MaxDepth(root.Right)
	if l == r {
		return root
	}
	if l > r {
		return subtreeWithAllDeepest(root.Left)
	} else {
		return subtreeWithAllDeepest(root.Right)
	}
}

/**
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
进阶：

你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

*/
func SortList(head *ListNode) *ListNode {
	//思路，由于要使用常数级的空间，所以使用冒泡排序的思路，比较当前节点和其后面节点的大小，如果比后面节点大，则交换双方的值
	tmp := head

	l := 0
	for tmp != nil {
		tmp = tmp.Next
		l++
	}

	for i := 0; i < l; i++ {
		tmp = head
		for tmp != nil {
			if tmp.Next != nil && tmp.Val > tmp.Next.Val {
				tmp.Val, tmp.Next.Val = tmp.Next.Val, tmp.Val
			}
			tmp = tmp.Next
		}
	}
	printList(head)
	return head
}

/**
较大分组的位置
在一个由小写字母构成的字符串 s 中，包含由一些连续的相同字符所构成的分组。
例如，在字符串 s = "abbxxxxzyy" 中，就含有 "a", "bb", "xxxx", "z" 和 "yy" 这样的一些分组。
分组可以用区间 [start, end] 表示，其中 start 和 end 分别表示该分组的起始和终止位置的下标。上例中的 "xxxx" 分组用区间表示为 [3,6] 。
我们称所有包含大于或等于三个连续字符的分组为 较大分组 。
找到每一个 较大分组 的区间，按起始位置下标递增顺序排序后，返回结果。

*/
func largeGroupPositions(s string) [][]int {
	//双指针
	res := make([][]int, 0)
	if len(s) < 3 {
		return res
	}
	p1 := 0
	p2 := p1 + 1
	maxLen := 1
	for p2 < len(s) {
		//如果当前p1和p2的字符不相等，则判断maxLen的长度是否大于等于3，如果是，则将p1和p2-1放入到res中。且此时将p1移动到p2，将p2移动到p2+1
		if s[p1] != s[p2] {
			if maxLen >= 3 {
				res = append(res, []int{p1, p2 - 1})
			}
			p1 = p2
			p2++
			maxLen = 1
		} else {
			//如果当前p1和p2相等，则将maxLen加1
			maxLen++
			p2++
			//如果加完之后p2刚到等于s的长度，则判断最后一组是否符合
			if p2 == len(s) && maxLen >= 3 {
				res = append(res, []int{p1, p2 - 1})
			}
		}
	}
	return res
}

/**
斐波那契数，通常用 F(n) 表示，形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
F(0) = 0，F(1) = 1
F(n) = F(n - 1) + F(n - 2)，其中 n > 1
给你 n ，请计算 F(n) 。
*/
func fib(n int) int {
	if n <= 1 {
		return n
	}
	res := 0
	first := 0
	second := 1
	for i := 2; i <= n; i++ {
		res = first + second
		first = second
		second = res
	}
	return res
}

/**
399. 除法求值
给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。
另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。
返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。

注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。

*/
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	//思路，假如["a", "b"]的值确定了，那毫无疑问["b", "a"]的值也确定了。  如果["a", "b"]的值确定了，["b", "c"]的值也确定了，那么["a", "c"]的值也确定了。
	return nil
}
