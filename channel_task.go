package main

import (
	"fmt"
)

func Task(taskch, resch chan int, exitch chan bool, i int) {
	defer func() { //异常处理
		err := recover()
		if err != nil {
			fmt.Println("do task error：", err)
			return
		}
	}()
	for t := range taskch { //  处理任务
		fmt.Println("第", i, "个消费者正在 do task :", t)
		resch <- t //
	}
	fmt.Println(fmt.Sprintf("task:%d 任务完成了", i))
	exitch <- true //处理完发送退出信号
}

func main() {
	taskch := make(chan int, 20) //任务管道
	resch := make(chan int, 20)  //结果管道
	exitch := make(chan bool, 5) //退出管道
	go func() {
		for i := 0; i < 10; i++ {
			taskch <- i
		}
		close(taskch)
	}()

	for i := 0; i < 5; i++ { //启动5个goroutine做任务
		go Task(taskch, resch, exitch, i)
	}

	go func() { //等5个goroutine结束
		for i := 0; i < 5; i++ {
			<-exitch
		}
		close(resch)  //任务处理完成关闭结果管道，不然range报错
		close(exitch) //关闭退出管道
	}()

	for res := range resch { //打印结果
		fmt.Println("task res：", res)
	}
}
