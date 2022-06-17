package main

import (
	"fmt"
	"strconv"
	"time"
)

/**
测试select.  select是go语言从语言层面提供的一种注册事件机制，类似linux中的一种事件注册方式
原理是：先注册一堆事件，哪个事件准备好了，就先执行哪个事件
*/
/*func main() {
	ch := make(chan int)

	fmt.Println("主函数执行")

	//开一个goroutine,为ch中写入一个值
	go func(ch chan int) {
		ch <- 1
	}(ch)

	//sleep1秒
	time.Sleep(time.Second)

	select {
	//从channel中读取数据
	case i := <-ch:
		fmt.Println("我从channel中读取到了：" + strconv.Itoa(i))
	default:
		fmt.Println("我只是个默认的程序")
	}
}*/

/**
这里是一个经典的延时控制方法
*/
func main() {
	ch := make(chan int)
	timeOut := make(chan int)

	go goSelect(ch, timeOut)

	//开一个goroutine,为ch中写入一个值
	go func(ch chan int, timeOut chan int) {
		//这里先设置对应的超时时间，假设是10秒
		time.Sleep(time.Second * 1)
		ch <- 1
		ch <- 2
		ch <- 3
		timeOut <- 111
	}(ch, timeOut)

	time.Sleep(time.Second * 10)
}

func goSelect(ch chan int, timeOut chan int) {
	for {
		select {
		//从channel中读取数据
		case i := <-ch:
			fmt.Println("我从channel中读取到了：" + strconv.Itoa(i))
		case j := <-timeOut:
			fmt.Println("我超时了！" + strconv.Itoa(j))
			/*default:
			fmt.Println("我只是个默认的程序")*/
		}
	}
}
