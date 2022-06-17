package main

import "fmt"

func main() {
	i := 10
	fmt.Println("i的地址是", &i)

	//ptr是一个指针变量，类型为*int(意思是指向int的指针)，本身的值是&i
	var ptr *int = &i
	fmt.Printf("value:%v\n", ptr)
	//用*prt就能取出ptr指针所指向的内存地址的值
	fmt.Println("ptr指向的值：", *ptr)

	fmt.Println("ptr的地址是", &ptr)

	//通过指针去修改原始的变量的值
	*ptr = 100
	fmt.Println("修改之后的i的值：", i)

}
