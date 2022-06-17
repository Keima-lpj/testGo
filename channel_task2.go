package main

import "fmt"

var (
	task_ch   = make(chan int, 20)
	result_ch = make(chan int, 20)
)

func main() {
	//新增20个任务
	go func() {
		for i := 0; i <= 19; i++ {
			task_ch <- i
		}
		close(task_ch)
	}()

	//开五个goroutine来做任务
	for i := 1; i <= 5; i++ {
		go doTask(i)
	}
	//从result_ch中读取值
	for v := range result_ch {
		fmt.Println(v)
	}
}

//做任务
func doTask(i int) {
	//从task_ch中取出任务，将值取平方之后，放入result_ch中
	fmt.Println("我是任务", i, "我开始做任务了")
	for task := range task_ch {
		fmt.Println("我取到了值：", task)
		result_ch <- task * task
	}
	fmt.Println("我已经做完了，没有任务了")
	//向结束chan发送结束信息
	return
}
