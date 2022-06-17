package main

import (
	//"fmt"
	//"runtime"
	//"sync"
	"fmt"
)

/*
func main() {
	//创建一个存储bool的channel
	c := make(chan bool)
	go func() {
		fmt.Println("Lets Go!")
		//执行之后将bool值存入channel中
		c <- true
	}()
	//线程阻塞，如果取到bool值，才会继续执行
	<-c
}
*/

/*func main() {
	//创建一个channel
	c := make(chan bool)
	go func() {
		fmt.Println("lets Go!!!")
		//执行之后将bool值存入channel中
		c <- true
		//关闭这个channel(如果不关闭channel,则会造成所有的goroutine等待，死锁。所以必须在某个地方关闭这个channel)
		close(c)
	}()
	//循环读取channel中的值
	for v := range c {
		fmt.Println(v)
	}
}*/

/*func main() {
	//创建一个channel
	c := make(chan bool, 1)
	go func() {
		fmt.Println("lets Go!!!")
		for v := range c {
			fmt.Println(v)
		}
		//fmt.Println(<-c)
	}()
	c <- true
	c <- false
}*/

/*func main() {
	//fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(8)
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i <= 1000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	c <- true
}*/

/*func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//创建一个同步工作组
	waitGroup := sync.WaitGroup{}
	//给这个工作组添加10个任务
	waitGroup.Add(10)
	for i := 0; i < 10; i++ {
		go Go(&waitGroup, i)
	}
	waitGroup.Wait()
}

func Go(waitGroup *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i <= 1000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	waitGroup.Done()
}*/

/*func main() {
	//新建两个channel.实现channel间的互相通信
	c1, c2 := make(chan int, 2), make(chan int)
	//首先先给c1中添加两个元素
	c1 <- 1
	c1 <- 2
	go Go(c1, c2)

}

func Go(c1 chan int, c2 chan int) {
	//接收到一个channel，将其中的值循环并存入另一个channel
	for v := range c1 {
		i := v + 5
		c2 <- i
	}
}*/

var c chan string

func pingPang() {
	i := 0
	for {
		//pingPang中先取数据，再存数据
		fmt.Println(<-c)
		c <- fmt.Sprintf("我是pingPang中的参数：#%d", i)
		i++
	}
}

func main() {
	c = make(chan string)
	go pingPang()
	for i := 0; i < 10; i++ {
		//先往c中存入数据
		c <- fmt.Sprintf("我是main中的参数：#%d", i)
		//再从c中取出
		fmt.Println(<-c)
	}
}
