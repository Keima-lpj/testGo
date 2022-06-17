package main

import "fmt"

func main() {
	s1 := "Hello Golong!"

	s2 := "Yes廖璞健!"

	fmt.Println(len(s1), len(s2)) //输出两个13

	/****************************遍历****************************/
	for k, v := range s1 {
		fmt.Printf("(%d, %X) ", k, v) //(0, 48) (1, 65) (2, 6C) (3, 6C) (4, 6F) (5, 20) (6, 47) (7, 6F) (8, 6C) (9, 6F) (10, 6E) (11, 67) (12, 21)
	}
	fmt.Println()

	for k, v := range s2 {
		fmt.Printf("(%d, %X) ", k, v) //(0, 59) (1, 65) (2, 73) (3, 5ED6) (6, 749E) (9, 5065) (12, 21)
	}
	fmt.Println()

	/****************************遍历[]byte****************************/

	for k, v := range []byte(s1) {
		//这里跟直接遍历s1一样
		fmt.Printf("(%d, %X) ", k, v)
	}
	fmt.Println()

	for k, v := range []byte(s2) {
		//这里跟直接遍历s2不一样，会将占用了三个字节的中文每一个都拆开成一个字节
		//(0, 59) (1, 65) (2, 73) (3, E5) (4, BB) (5, 96) (6, E7) (7, 92) (8, 9E) (9, E5) (10, 81) (11, A5) (12, 21)
		fmt.Printf("(%d, %X) ", k, v)
	}
	fmt.Println()

	/****************************遍历[]rune****************************/
	for k, v := range []rune(s1) {
		//这个跟直接遍历s1一样
		fmt.Printf("(%d, %X) ", k, v)
	}
	fmt.Println()

	for k, v := range []rune(s2) {
		fmt.Printf("(%d, %x) ", k, v)
	}
	fmt.Println()
	for k, v := range []rune(s2) {
		//这里使用了c%（将unicode码值转换为正常值），可以正常输出字符串的内容和其在字符串中的位置
		fmt.Printf("(%d, %c) ", k, v)
	}

}
