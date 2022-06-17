package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {
	//声明一个channel
	ch := make(chan int)

	//以单核模式来跑程序
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 1; i <= 100; i++ {
			if i == 10 {
				//主动要求让出cpu
				//runtime.Gosched()
				//或者协程阻塞,此时这里的channel是非缓冲channel，读不出来数据就会阻塞在这里。
				<-ch
			}
			fmt.Println("coroutine 1:" + strconv.Itoa(i))
		}
	}()

	go func() {
		for i := 100; i <= 200; i++ {
			fmt.Println("coroutine 2:" + strconv.Itoa(i))
		}
		ch <- 1
	}()

	time.Sleep(time.Second)

}
