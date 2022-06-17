package main

import (
	"fmt"
)

type worker struct {
	c    chan int
	done chan bool
	//这里定义了一个完成的函数，这个函数可以由用户自己高度自定义
	doneFunc func(done chan bool)
}

func main() {
	channelDemo3()
}

func doWork(id int, w worker) {
	for n := range w.c {
		fmt.Printf("worker id %d，reviced work %c\n", id, n)
	}
	//工人今天的事情做完了，给一个信号
	w.doneFunc(w.done)
}

func channelDemo3() {
	//定义10个工人
	var workers [10]worker

	//循环
	for i := 0; i < 10; i++ {
		//新建工人
		workers[i] = worker{
			c:    make(chan int),
			done: make(chan bool),
			//定义了一个完成工作的函数，将信号写入done中
			doneFunc: func(done chan bool) {
				done <- true
			},
		}
		//这个工人开始工作
		go doWork(i, workers[i])
	}

	//分配两次任务
	for k, v := range workers {
		//为每个工人分配任务，任务写入v.c中
		v.c <- 'a' + k
	}
	for k, v := range workers {
		v.c <- 'A' + k
	}
	//关闭所有的c，此时工人接受不到信号了，则认定为今天的工作做完了，开始往done中写数据
	for _, v := range workers {
		close(v.c)
	}
	//等待任务执行完成
	for _, v := range workers {
		<-v.done
	}
}
