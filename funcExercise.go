package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var caishuziCount int

func main() {

	/*var n int
	fmt.Println("请输入一个数字")
	fmt.Scanln(&n)

	jinZiTa(n)

	jiujiuChengFa(n)*/

	/*err := monthTest()
	if err != nil {
		fmt.Println(err)
	}*/

	//caiShuZi()

	fbn(6)

	maoPao()

	arr := []int{24, 96, 80, 57, 13}

	MaoPao1(&arr)

	fmt.Println(arr)

	fmt.Println(erFenFind(arr, 13, 0, 4))
}

//输入一个数，输出一个金字塔形
func jinZiTa(n int) {

	for i := 1; i <= n; i++ {
		//第i行，前面要输出n-i个空格，然后输出2(i+1)个*
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/**
输入一个n，打印一到n的乘法表
*/
func jiujiuChengFa(n int) {
	//i表示层数
	for i := 1; i <= n; i++ {
		//j表示列数
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d ", i, j, i*j)
		}
		fmt.Println()
	}
}

/**
输入年份，月份，打印该月份的天数
*/
func monthTest() error {
	var year int
	var month int
	var isRunYear bool
	fmt.Println("请输入年：")
	fmt.Scanln(&year)
	fmt.Println("请输入月：")
	fmt.Scanln(&month)

	//这里判断输入的月份是否有误
	if month < 1 || month > 12 {
		return errors.New("你输错了月份！")
	}

	fmt.Printf("您输入的日期为：%d年%d月\n", year, month)

	//判断该年是否为闰年，判断该月份有多少天
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		isRunYear = true
	} else {
		isRunYear = false
	}

	var monthDay int

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		monthDay = 31
	case 4, 6, 9, 11:
		monthDay = 30
	case 2:
		if isRunYear {
			monthDay = 29
		} else {
			monthDay = 28
		}
	}

	fmt.Printf("您输入的月份有%d天\n", monthDay)
	return nil
}

/**
随机生成一个1到100的整数，有十次机会
如果第一次就猜中，提示你真是个天才
如果2-3次猜中，提示你很聪明。
如果4-9次猜中，提示一般般
如果最后一次猜中，提示可算猜对了
如果一次都没猜中，提示你这个智障
*/
func caiShuZi() {
	//给定一个时间的种子
	rand.Seed(time.Now().Unix())
	num := rand.Intn(100)

	fmt.Println("生成的数字为：", num)

	var putIn int

	for i := 1; i <= 10; i++ {
		fmt.Scanln(&putIn)
		if putIn == num {
			caishuziCount = i
			break
		} else {
			fmt.Println("没猜对，继续")
		}
	}

	if caishuziCount == 1 {
		fmt.Println("你真是个天才")
	} else if caishuziCount >= 2 && caishuziCount <= 3 {
		fmt.Println("你很聪明")
	} else if caishuziCount >= 4 && caishuziCount <= 9 {
		fmt.Println("一般般")
	} else {
		fmt.Println("妈的智障")
	}
}

//编写一个函数，接收n int，将斐波那契数列放入切片中
func fbn(n int) {
	//声明一个长度为n的切片
	var intSlice = make([]uint64, n)

	intSlice[0] = 1
	intSlice[1] = 1

	for i := 2; i < n; i++ {
		intSlice[i] = intSlice[i-1] + intSlice[i-2]
	}

	fmt.Println(intSlice)
}

//写一个冒泡排序
func maoPao() {
	arr := [5]int{24, 69, 80, 57, 13}
	var temp int

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println(arr)
}

func MaoPao1(arr *[]int) {
	fmt.Println("排序前的数组：", *arr)
	var temp int

	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j+1]
				(*arr)[j+1] = (*arr)[j]
				(*arr)[j] = temp
			}
		}
	}
}

//写一个二分查找，返回查到的切片的下标。(此切片必须为一个有序切片)
func erFenFind(arr []int, num, leftIndex, rightIndex int) int {
	middle := (leftIndex + rightIndex) / 2
	//判断数组中间的值是大于num还是小于num
	if arr[middle] == num {
		return middle
	} else if arr[middle] > num {
		return erFenFind(arr, num, leftIndex, middle-1)
	} else {
		return erFenFind(arr, num, middle+1, rightIndex)
	}
}
