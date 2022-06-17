package algorithm

import "sort"

/**
面试题 08.11. 硬币
硬币。给定数量不限的硬币，币值为25分、10分、5分和1分，编写代码计算n分有几种表示法。
(结果可能会很大，你需要将结果模上1000000007)
示例1:
 输入: n = 5
 输出：2
 解释: 有两种方式可以凑成总金额:
5=5
5=1+1+1+1+1
示例2:
 输入: n = 10
 输出：4
 解释: 有四种方式可以凑成总金额:
10=10
10=5+5
10=5+1+1+1+1+1
10=1+1+1+1+1+1+1+1+1+1

*/
func WaysToChange(n int) int {
	//一看就是动态规划，没啥好说的。 25,10,5,1
	dp := make([]int, n+1)
	dp[0] = 1
	coins := []int{1, 5, 10, 25}
	for i := 0; i < 4; i++ {
		for j := 1; j <= n; j++ {
			if j-coins[i] >= 0 {
				dp[j] += dp[j-coins[i]]
			}
		}
	}
	return dp[n] % 1000000007
}

/**
面试题 08.12. 八皇后
设计一种算法，打印 N 皇后在 N × N 棋盘上的各种摆法，其中每个皇后都不同行、不同列，也不在对角线上。
这里的“对角线”指的是所有的对角线，不只是平分整个棋盘的那两条对角线。

注意：本题相对原题做了扩展
示例:
 输入：4
 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
 解释: 4 皇后问题存在如下两个不同的解法。
[
 [".Q..",  // 解法 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // 解法 2
  "Q...",
  "...Q",
  ".Q.."]
]
*/

func solveNQueens(n int) [][]string {
	var tmp = make([][]byte, n)
	// 初始化
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			tmp[r] = append(tmp[r], '.')
		}
	}
	var res [][]string
	backTrack(0, tmp, &res)
	return res
}

func backTrack(row int, tmp [][]byte, res *[][]string) {
	if row == len(tmp) {
		var subRes []string
		for i := 0; i < row; i++ {
			subRes = append(subRes, string(tmp[i]))
		}
		*res = append(*res, subRes)
		return
	}
	for c := 0; c < len(tmp); c++ {
		if !ok(row, c, tmp) {
			continue
		}
		//当前位置符合条件，将皇后放入
		tmp[row][c] = 'Q'
		//处理下一行的皇后
		backTrack(row+1, tmp, res)
		//当前位置可以，但是没说后面的位置就不行，那么将当前位置放入. 并继续循环后续的位置
		tmp[row][c] = '.'
	}
}

func ok(row int, col int, res [][]byte) bool {
	for i := 0; i < col+1; i++ {
		//列检测
		if res[row][i] == 'Q' {
			return false
		}
	}
	for j := 0; j < row+1; j++ {
		//行检测
		if res[j][col] == 'Q' {
			return false
		}
	}
	//右下角至左上角检测
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if res[i][j] == 'Q' {
			return false
		}
	}
	//左下角至右上角检测
	for i, j := row-1, col+1; i >= 0 && j < len(res); i, j = i-1, j+1 {
		if res[i][j] == 'Q' {
			return false
		}
	}
	return true
}

/**
面试题 17.08. 马戏团人塔
有个马戏团正在设计叠罗汉的表演节目，一个人要站在另一人的肩膀上。出于实际和美观的考虑，
在上面的人要比下面的人矮一点且轻一点。已知马戏团每个人的身高和体重，请编写代码计算叠罗汉最多能叠几个人。

示例：
输入：height = [65,70,56,75,60,68] weight = [100,150,90,190,95,110]
输出：6
解释：从上往下数，叠罗汉最多能叠 6 层：(56,90), (60,95), (65,100), (68,110), (70,150), (75,190)
提示：
height.length == weight.length <= 10000
*/
func bestSeqAtIndex(height []int, weight []int) int {
	l := len(height)
	if l == 0 {
		return 0
	}
	//将数组按照身高升序排列
	person := make([][2]int, l)
	for i := 0; i < l; i++ {
		person[i] = [2]int{height[i], weight[i]}
	}

	//将person按照身高降序排列，如果身高相等，则按照体重升序排列（这里是为了在身高相等的时候，取体重最大的人）
	sort.Slice(person, func(i, j int) bool {
		if person[i][0] == person[j][0] {
			return person[i][1] < person[j][i]
		}
		return person[i][0] > person[j][0]
	})

	var result [][2]int
	//这样的话，身高一定满足从大到小。
	for _, v := range person {
		//在结果中查询第一个不能p不能叠在上面的人，二分查找
		j := sort.Search(len(result), func(i int) bool {
			c := result[i]
			return c[0] <= v[0] || c[1] <= v[1]
		})
		//将这个人替换成p
		if j == len(result) {
			result = append(result, [2]int{v[0], v[1]})
		} else {
			result[j] = [2]int{v[0], v[1]}
		}
	}
	return len(result)
}

/**
面试题 08.13. 堆箱子
堆箱子。给你一堆n个箱子，箱子宽 wi、深 di、高 hi。箱子不能翻转，将箱子堆起来时，下面箱子的宽度、高度和深度必须大于上面的箱子。
实现一种方法，搭出最高的一堆箱子。箱堆的高度为每个箱子高度的总和。
输入使用数组[wi, di, hi]表示每个箱子。

示例1:
 输入：box = [[1, 1, 1], [2, 2, 2], [3, 3, 3]]
 输出：6
示例2:
 输入：box = [[1, 1, 1], [2, 3, 4], [2, 6, 7], [3, 4, 5]]
 输出：10
*/
func pileBox(box [][]int) int {
	l := len(box)
	if l == 0 {
		return 0
	}
	//将盒子按照升序排序
	sort.Slice(box, func(i, j int) bool {
		if box[i][0] == box[j][0] {
			if box[i][1] == box[i][1] {
				return box[i][2] < box[i][2]
			}
			return box[i][1] < box[i][1]
		}
		return box[i][0] < box[j][0]
	})

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	dp := make([]int, l+1)
	res := 0
	for i := 0; i < l; i++ {
		//dp[i]表示用第i个为底可以堆多高
		dp[i] = box[i][2]
		//由于box是升序排列，所以以第i个箱子为底，只能叠其前面的箱子
		for j := i - 1; j >= 0; j-- {
			if box[i][0] > box[j][0] && box[i][1] > box[j][1] && box[i][2] > box[j][2] {
				dp[i] = max(dp[i], dp[j]+box[i][2])
			}
		}
		res = max(res, dp[i])
	}
	return res
}

/**
面试题 10.03. 搜索旋转数组
搜索旋转数组。给定一个排序后的数组，包含n个整数，但这个数组已被旋转过很多次了，次数不详。
请编写代码找出数组中的某个元素，假设数组元素原先是按升序排列的。若有多个相同元素，返回索引值最小的一个。

示例1:
 输入: arr = [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], target = 5
 输出: 8（元素5在该数组中的索引）
示例2:
 输入：arr = [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], target = 11
 输出：-1 （没有找到）
*/
func Search2(arr []int, target int) int {
	//二分查找。
	var i, j = 0, len(arr) - 1
	for i < j {
		mid := i + (j-i)/2
		if arr[mid] == target {
			//判断左测是否还有符合条件的值
			for arr[mid] == target {
				mid--
			}
			return mid + 1
		} else {
			if arr[mid] >= arr[j] {
				//此时转折点在mid右侧
				if target < arr[mid] {
					j = mid - 1
				} else {
					i = mid + 1
				}
			} else {
				//此时转折点在左侧
				if target < arr[mid] {
					j = mid - 1
				} else {
					i = mid + 1
				}
			}
		}
	}
	return -1
}
