package main

import "fmt"

func main() {
	numChan := make(chan int, 50)
	sumMap := make(chan map[int]int)

	for i := 1; i <= 50; i++ {
		numChan <- i
	}
	close(numChan)

	for i := 1; i <= 8; i++ {
		go leiJia(numChan, sumMap, i)
	}

	//这里遍历sumMap
	for i := 1; i <= 8; i++ {
		v := <-sumMap
		for k1, v1 := range v {
			fmt.Println(k1, v1)
		}
	}

}

func leiJia(numChan chan int, sumMap chan map[int]int, x int) {
	//从numChan中取值，并且计算其累加值放入sumChan中
	map111 := make(map[int]int)
	for v := range numChan {
		//这里计算V的累加
		res := 0
		for i := 1; i <= v; i++ {
			res += i
		}
		map111[v] = res
	}
	sumMap <- map111
	//读完了，往完成chan里写一个信号
	fmt.Println(x, "已经完成了任务")
}
