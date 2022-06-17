package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	c := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			c <- i
			i++
		}
	}()
	return c
}

func workForSelect(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker id %d，reviced work %d\n", id, n)
	}
}

func createWorkForSelect(id int) chan<- int {
	c := make(chan int)
	go workForSelect(id, c)
	return c
}

func main() {
	c1, c2 := generator(), generator()
	worker := createWorkForSelect(0)

	//写法1，这种写法可以满足，但是存在一个问题。每次取到数据往worker中塞的时候
	//由于worker是一个阻塞的chan，如果此时worker没有处理完，则后续的操作都阻塞了
	//只有等worker处理完毕之后，才能继续往其中塞数据
	/*for {
		select {
		case n := <-c1:
			worker <- n
		case n := <-c2:
			worker <- n
		}
	}*/

	//写法2，这种写法避免了写法1中的阻塞问题，c1和c2都可以源源不断的流入数据，但是如果c1和c2的数据流入速度大于worker的消费速度
	//则此时n的数据会不断被刷新，没来得及消费的数据会丢失
	/*n := 0
	hasValue := false
	for {
		//这里为什么一定要定义一个空值，因为要使用nil chan的特性，当select调度到nil chan时，不会触发case
		var nilWorker chan<- int
		if hasValue {
			nilWorker = worker
		}
		select {
		case n = <-c1:
			hasValue = true
		case n = <-c2:
			hasValue = true
		case nilWorker <- n:
			hasValue = false
		}
	}*/

	//After函数会返回一个只可以取数据的chan
	exit := time.After(10 * time.Second)
	//Tick函数会返回一个只取数据的chan，并且每隔设置的时间发送一次信号
	tick := time.Tick(time.Second)
	//写法3，思路，将接受的数据存入一个数组中，然后从数组中读取数据进行消费
	//这样就解决了写法2中丢失数据的问题。
	//这就是go语言中电信的csp模型，没有通过锁，就完成了通过通信共享内存这件事
	var task []int
	for {
		//这里为什么一定要定义一个空值，因为要使用nil chan的特性，当select调度到nil chan时，不会触发case
		var nilWorker chan<- int
		var preTask int
		//其中有数据要处理
		if len(task) > 0 {
			nilWorker = worker
			preTask = task[0]
		}
		select {
		case n := <-c1:
			task = append(task, n)
		case n := <-c2:
			task = append(task, n)
		case <-time.After(800 * time.Millisecond):
			fmt.Println("本次select timeout")
		case <-tick:
			fmt.Println("queue lens:", len(task))
		case nilWorker <- preTask:
			fmt.Println("任务开始：", task)
			task = task[1:]
			fmt.Println("任务剩余：", task)
		case <-exit:
			fmt.Println("到时间了，不玩了")
			return
		}
	}

}
