package project1

import "fmt"

//这里给Fish结构体新增一个方法

func (f *Fish) Say() {
	fmt.Println("I is fish AAA")
}

func getInt() int {
	return 100
}
