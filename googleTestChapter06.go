package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//testGoroutine()
	testThread()
}

func testGoroutine() {
	var a [10]int
	//如何去理解协程是自己主动交出控制权呢，这里加一句设定为单核CPU运行（如果不加，则默认是多核运行，默认是可以多个协程同时跑的）
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		//这里要注意，要将i做为参数传进去。为什么一定要传进去？
		//因为如果不传参数，那么下面的goroutine就会作为一个闭包，会引用上面的i
		//这个时候，由于main函数最后一次循环执行完成后，i会执行最后一次++  变成10
		//此时的a[10]就会提示越界 index out of range
		go func(i int) {
			//这里的println由于是一个IO操作，调度器会自动切换协程
			//fmt.Println("当前goroutine执行", i)
			//如果换成了不是IO操作的操作呢
			for {
				//由于上面是单核CPU运行，则第一个协程因为是死循环，永远不会退出，main程序也取不到CPU的控制权。
				a[i]++
				//这个时候需要加上一句runtime.Gosched()，这个函数的意思是交出当前协程的控制权。
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Microsecond)

	//注意！！！ 如果是不设置GOMAXPROCS的数量，则默认会以当前机器的核数来运行程序
	//我的电脑是双核四线程，则默认值是4，但是为什么每次a最多只有三个数呢，因为还有一个核被main函数占住了。
	fmt.Println(a)
}

func testThread() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			fmt.Println("Hello world from goroutine %d\n", i)
		}(i)
	}
	time.Sleep(time.Minute)
}
