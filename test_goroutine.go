package main

import (
	"fmt"
	"sync"
	"time"
)

//计算1到200这两百个数字合

var (
	Num = make(map[int]int, 10)

	lock sync.Mutex
)

func main() {
	for i := 1; i <= 50; i++ {
		//这里会出现concurrent map writes错误，多个协程不能同时并发操作一个map
		go qiuHe(i)
	}
	//休眠十秒钟
	time.Sleep(time.Second * 5)
	//输出map中的值
	for k, v := range Num {
		fmt.Printf("%d的求和是:%d\n", k, v)
	}
}

func qiuHe(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	lock.Lock()
	Num[n] = res
	lock.Unlock()
}
