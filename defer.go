package main

import (
	"fmt"
)

func main(){
	fmt.Println("Start")

	defer fmt.Println("a1")

	defer fmt.Println("a2")

	testDefer()

	defer fmt.Println("a3")


	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("End")
}


func testDefer(){

	fmt.Println("111")

	defer fmt.Println("我是testDefer里面的第一句")

	defer fmt.Println("我是testDefer里面的第二句")

	fmt.Println("222")
}
