package algorithm

//快速排序
var quickRes []int

func QuickSortMain(nums []int) []int {
	quickRes = nums
	quickSort(0, len(quickRes)-1)
	return quickRes
}

func quickSort(left, right int) {
	if left > right {
		return
	}
	i, j := left, right
	//找一个基准值，默认为最左侧的值
	tmp := quickRes[left]

	for i < j {
		//右侧先动（由于以最左侧的值作为基准），从右往左找一个小于基准值的值
		for quickRes[j] >= tmp && i < j {
			j--
		}
		//左侧再动
		for quickRes[i] <= tmp && i < j {
			i++
		}
		//如果此时left<right，则交换两个值的位置
		if i < j {
			quickRes[i], quickRes[j] = quickRes[j], quickRes[i]
		}
	}
	//将基准值归位
	quickRes[left], quickRes[i] = quickRes[i], tmp
	//递归处理基准值归位后左右的数组
	quickSort(left, i-1)
	quickSort(i+1, right)
}
