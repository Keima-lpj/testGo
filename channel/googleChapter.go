package main

import (
	"fmt"
	"time"
)

func main() {
	//channelDemo()
	//channelDemo2()
	//channelClose()
	//channelA()
}

func channelDemo() {
	var channels [10]chan int
	//这里开十个工人
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go work(i, channels[i])
	}

	//给这十个工人分配任务
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	//time.Sleep(time.Microsecond)

}

func work(id int, c chan int) {
	/*for {
		//如果c没有被关闭，则从c中读取任务
		if n, ok := <-c; ok {
			fmt.Printf("worker id %d，reviced work %c\n", id, n)
		} else {
			//如果c被关闭了，则跳出循环
			break
		}
	}*/
	//还有一种更简便的做法,这种做法等同于上面的做法，如果c被关闭，则会跳出循环
	for n := range c {
		fmt.Printf("worker id %d，reviced work %c\n", id, n)
	}
}

func channelDemo2() {
	var channels [10]chan<- int
	//这里开十个工人
	for i := 0; i < 10; i++ {
		channels[i] = createWork(i)
	}

	//给这十个工人分配任务
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Microsecond)
}

//这个work跟第一个不同，这个work其实是创建了一个goroutine在等待读取数据，同时返回了一个channel
//所以它本质上并不是一个worker，他应该是一个worker的创建者
//如果要严格一点，则可以在返回channel的时候定义好，这个chan是用来写数据的，因为我已经有goroutine在读数据了
//如果要求这个channel只能读，则定义为chan<-    如果要求只能写，则定义为<-chan
func createWork(id int) chan<- int {
	c := make(chan int)
	go work(id, c)
	return c
}

func channelClose() {
	c := make(chan int)
	go work(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Microsecond)
}

func channelA() {
	c := make(chan int)
	go func() {
		for n := range c {
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	//此时不关闭c，则上面的goroutine会阻塞等待channel写入数据
	//此时关闭c,则上面的goroutine会一直从channel中读取chan类型的0值
	//如果不用死循环，用for-range的形式，则会从channel中读数据一直到channel关闭为止，如果不显式的关闭channel，则这个循环会阻塞直到main函数退出
	close(c)
	time.Sleep(time.Microsecond)
}
