package algorithm

/**
无重复元素的排序数组里找符合条件的值
*/
func bSearch(a []int, target int) int {
	l := len(a)
	low, high := 0, l-1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] == target {
			return mid
		} else if a[mid] < target {
			low = mid + 1
		} else if a[mid] > target {
			high = mid - 1
		}
	}
	return -1
}

/**
找第一个等于target的元素（数组可能重复）
*/
func bSearch1(a []int, target int) int {
	l := len(a)
	low, high := 0, l-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] > target {
			high = mid - 1
		} else if a[mid] < target {
			low = mid + 1
		} else {
			//如果当前的mid为0，直接返回。如果mid-1不等于目标值，则代表当前值就是第一个等于target的元素
			if mid == 0 || a[mid-1] != target {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

/**
找最后一个等于target的元素（数组可能重复）
*/
func bSearch2(a []int, target int) int {
	l := len(a)
	low, high := 0, l-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] > target {
			high = mid - 1
		} else if a[mid] < target {
			low = mid + 1
		} else {
			//如果当前的mid为0，直接返回。如果mid-1不等于目标值，则代表当前值就是第一个等于target的元素
			if mid == l-1 || a[mid+1] != target {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

/**
找第一个大于给定目标值的元素
*/
func bSearch3(a []int, target int) int {
	l := len(a)
	low, high := 0, l-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] >= target {
			if mid == 0 || a[mid-1] < target {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}

/**
找最后一个小于给定目标值的元素
*/
func bSearch4(a []int, target int) int {
	l := len(a)
	low, high := 0, l-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] <= target {
			if mid == l-1 || a[mid+1] > target {
				return mid
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}
	}
	return -1
}

/**
假设a是一个有序数组被旋转过之后的结果   例如：a = [4,5,6,1,2,3]
*/
func bSearch5(a []int, target int) int {
	l := len(a)
	low, high := 0, l-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] == target {
			return mid
		}
		//证明0到mid之间升序
		if a[0] <= a[mid] {
			//如果target在a[0]到a[mid]这个升序数组之间，则high = mid - 1
			if a[0] <= target && target < a[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			//0到mid之间不是升序，则转折点必在0到mid之间，则mid到l-1之间升序
			//如果target在a[mid]到a[l-1]之间，则有
			if a[mid] < target && target <= a[l-1] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
