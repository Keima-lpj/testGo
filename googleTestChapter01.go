package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"strconv"
	"strings"
)

func main() {
	//验证byte和uint8是同一个类型
	var a uint8 = 10
	var b byte
	b = a
	fmt.Println(a, b)

	//验证rune和int32是同一类型
	var c int32 = 100
	var d rune
	d = c
	fmt.Println(c, d)

	euler()

	fmt.Println(convertIntToBin(5), convertIntToBin(13))

	//checkSlice()

	checkMap()

	fmt.Println(getLongString("qwerrtzq"))
	fmt.Println(getLongString("abcbcbc"))
	fmt.Println(getLongString("bbbbbbbb"))
	fmt.Println(getLongString("asdfaskdhf"))

}

//验证欧拉公式
func euler() {
	//这是一个复数，3为实部，4i为虚部
	x := 3 + 4i
	//取模打印出5
	fmt.Println(cmplx.Abs(x))

	//欧拉公式   e的iπ次方加上1等于0
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

//将int整形转为2进制
func convertIntToBin(n int) (result string) {
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return
}

//验证slice
func checkSlice() {
	arr := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println(s1, s2) //[2 3 4 5] [5 6]
	//s3 := s1[-1:2]  负数是不行的，会报错

	s1 = []int{0, 2, 4, 6, 8}
	printValLenCap(s1)
	//slice的copy操作
	s3 := make([]int, 16, 32)
	copy(s3, s1)
	printValLenCap(s3)

	//如果此时我要讲s3中间的某个元素去掉，比如去掉第三个元素，则思路应该为s3 = s3[:3] + s3[4:]，使用append实现
	s3 = append(s3[:3], s3[4:]...)
	printValLenCap(s3)

	//从头部弹出元素
	front := s3[0]
	s3 = s3[1:]

	//从尾部弹出元素
	end := s3[len(s3)-1]
	s3 = s3[:len(s3)-1] //这一步操作后，会改变s3的cap

	fmt.Sprintln(front, end)
	printValLenCap(s3)

}

func checkMap() {

	m := map[int]string{
		1: "1",
		2: "2",
		3: "3",
		4: "4",
	}

	fmt.Println(m) //每次输出的顺序可能不同，因为map是无序的，是一个hash表

	for k, v := range m {
		fmt.Println(k, v) //这里也一样
	}

	fmt.Println(m[5]) //打印map中不存在的键时，由于此时值是string类型，则会打出一个空字符串, 如果是int，则会打印出0

	//如何判断这个键在map中是否存在呢，使用如下方式，如果存在，ok会为true，否则会为false
	a, ok := m[4]
	fmt.Println(a, ok)
	b, ok := m[5]
	fmt.Println(b, ok)

	if c, ok := m[5]; ok {
		fmt.Println(c)
	}

	//删除map中的元素
	delete(m, 4)
	fmt.Println(m)

}

//根据输入的字符串，找出其中不含重复字符串的最长子串
func getLongString(str string) (int, string) {
	//定义一个map
	m := make(map[byte]int)
	start := 0
	maxLength := 0
	returnString := ""

	strSlice := strings.Split(str, "")
	//a s d f a s k d h f
	//0 1 2 3 4 5 6 7 8 9

	//其中k代表字符byte的位子，v代表具体的字符
	for k, v := range []byte(str) {
		//1、判断这个值和start的位置关系，如果这个值大于start，则证明这个值已经在m中出现了，
		// 例如【a】【s】【d】【f】【a】【s】【k】【d】【h】【f】，在出现第二个a的时候，将start从第一个a的位置移动到了第一个s的位置，则开始从第一个s计数，第一个a可以视为被抛弃了。
		// 走到第二个s时，则将start从第二个s的位置移动到d的位置，则此时又从d开始计数
		// 如果小于start或者压根没有，则可以拿进来
		if val, ok := m[v]; ok && val >= start {
			start = m[v] + 1
		}
		//2、更新maxLength
		if k-start+1 > maxLength {
			maxLength = k - start + 1
			returnString = strings.Join(strSlice[start:k+1], "")
		}
		//3、将v的位置和字符放入m中
		m[v] = k
	}

	return maxLength, returnString

}

func printValLenCap(s []int) {
	fmt.Printf("Val:%v, Len:%d, Cap:%d\n", s, len(s), cap(s))
}
