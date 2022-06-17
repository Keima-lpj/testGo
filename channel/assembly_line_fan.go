package main

import (
	"fmt"
	"sync"
)

//生产者
func producerFan(nums ...int) chan int {
	fmt.Println("生产者开始了")
	out := make(chan int)
	go func() {
		fmt.Println("生产者协程开始了")
		defer close(out)
		for _, n := range nums {
			fmt.Println("生产者协程存入数据", n)
			out <- n
		}
		fmt.Println("生产者协程结束了")
	}()
	fmt.Println("生产者结束了")
	return out
}

//消费者
func squareFan(in chan int, num int) chan int {
	fmt.Println("消费者", num, "开始了")
	out := make(chan int)
	go func() {
		fmt.Println("消费者", num, "协程开始了")
		defer close(out)
		for n := range in {
			fmt.Println("消费者", num, "协程从任务列表中取出了", n)
			fmt.Println("消费者", num, "存入数据", n*n)
			out <- n * n
		}
		fmt.Println("消费者", num, "协程结束了")
	}()
	fmt.Println("消费者", num, "结束了")
	return out
}

//合并流出的任务通道
func mergeFan(outs ...chan int) chan int {
	out := make(chan int)

	wait := sync.WaitGroup{}
	wait.Add(len(outs))

	for _, c := range outs {
		go func() {
			defer wait.Done()
			for v := range c {
				out <- v
			}
		}()
	}

	go func() {
		wait.Wait()
		close(out)
	}()

	return out
}

func main() {
	in := producerFan(1, 2, 3, 4, 5, 6)
	//获取两个消费者
	out1 := squareFan(in, 1)
	out2 := squareFan(in, 2)

	out := mergeFan(out1, out2)

	for ret := range out {
		fmt.Printf("%3d", ret)
	}

}
