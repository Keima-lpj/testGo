package main

import (
	"fmt"
	"time"
)

func main() {
	var s int // 声明变量 s 是需要判断的数
	fmt.Println("输入一个数字：")
	fmt.Scan(&s)

	if s%2 == 0 { //     取 s 处以 2 的余数是否等于 0
		fmt.Print("s 是偶数\n") //如果成立
	} else {
		fmt.Print("s 不是偶数\n") //否则
	}
	fmt.Print("s 的值是：", s)

	//这里sleep是为了看编译之后点击执行的效果
	time.Sleep(time.Second * 5)
}
