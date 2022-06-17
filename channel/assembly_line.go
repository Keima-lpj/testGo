package main

import (
	"fmt"
	"time"
)

func producer(nums ...int) chan int {
	fmt.Println("主函数1开始了")
	out := make(chan int)
	go func() {
		fmt.Println("协程1开始了")
		defer close(out)
		for _, n := range nums {
			fmt.Println("协程1存入数据", n)
			out <- n
		}
		fmt.Println("协程1结束了")
	}()
	fmt.Println("主函数1结束了")
	return out
}

func square(inCh <-chan int) chan int {
	fmt.Println("主函数2开始了")
	out := make(chan int)
	go func() {
		fmt.Println("协程2开始了")
		defer close(out)
		for n := range inCh {
			fmt.Println("协程2从协程1中取出了", n)
			fmt.Println("协程2存入数据", n*n)
			out <- n * n
		}
		fmt.Println("协程2结束了")
	}()
	fmt.Println("主函数2结束了")
	return out
}

//这个例子可以看出来，channel并不需要在完全写入之后关闭才能使用for range。可以一边写入一边用for range读取，只要写入完毕之后有close操作即可
func testCloseChan() {
	//新建一个无缓冲的信道
	ch := make(chan int)
	//新开一个协程,写入数据
	go func() {
		defer func() {
			time.Sleep(time.Second * 2)
			fmt.Println("1结束")
			close(ch)
		}()
		for i := 0; i <= 3; i++ {
			fmt.Println("写入数据", i)
			ch <- i
		}
	}()
	//开一个协程，读取数据
	go func() {
		for v := range ch {
			fmt.Println("读取数据", v)
			fmt.Println(v)
		}
		fmt.Println("2结束")
	}()
}

func main() {
	in := producer(1, 2, 3, 4)
	ch := square(in)
	// consumer
	for ret := range ch {
		fmt.Printf("%3d", ret)
	}
	fmt.Println()

	testCloseChan()

	time.Sleep(time.Second * 5)
}
