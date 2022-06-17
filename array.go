package main

import (
	"fmt"
)

func main() {
	a := [3]string{"a", "1", 2: "0"}
	fmt.Println(a)
	b := [...]int{50: 1}

	fmt.Println(b)

	p1 := new([10]int)
	fmt.Println(p1)
	fmt.Printf("%T, %p\n", p1, p1)
	fmt.Println(*p1)

	c := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}
	fmt.Println(c)

	var x = [...]int{9, 7, 15, 66, 28, 17, 1}
	len1 := len(x)
	for i := 0; i < len1; i++ {
		for j := i + 1; j < len1; j++ {
			if x[i] < x[j] {
				temp := x[i]
				x[i] = x[j]
				x[j] = temp
			}
		}
	}
	fmt.Println(x)

	testArr := [3]int{10, 20, 30}
	test01(testArr)
	fmt.Println(testArr)
	test02(&testArr)
	fmt.Println(testArr)



	var zimu [26]byte
	for i := 0; i < len(zimu); i++{
		zimu[i] = byte('A' + i)
	}

	fmt.Println(zimu)

	arrSum := [5]float64{1, 1.1, 5.3, 4.4, 2.8}
	fmt.Println(sumArr(&arrSum))
}

//数组传值
func test01(arr [3]int){
	arr[2] = 88
}

func test02(arr *[3]int){
	(*arr)[2] = 88
}

/**
对一个float64数组求和
 */
func sumArr(arr *[5]float64) (sum float64){
	for _, v := range *arr{
		sum += v
	}
	return
}