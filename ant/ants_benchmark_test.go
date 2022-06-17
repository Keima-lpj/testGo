package main

import (
	ants "github.com/panjf2000/ants/v2"
	"sync"
	"testing"
)

const (
	//任务数量
	benchmarkTaskNum = 1000000
	antsPollSize     = ants.DefaultAntsPoolSize
	//antsPollSize = 500000
)

//测试使用常规方案
func BenchmarkGoroutines(b *testing.B) {
	//runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(benchmarkTaskNum)
		for j := 0; j < benchmarkTaskNum; j++ {
			go func() {
				demoFunc()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

//测试使用ants
func BenchmarkAntsPool(b *testing.B) {
	//runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	p, _ := ants.NewPool(antsPollSize)
	defer p.Release()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(benchmarkTaskNum)
		for j := 0; j < benchmarkTaskNum; j++ {
			go func() {
				_ = p.Submit(func() {
					demoFunc()
					wg.Done()
				})
			}()
		}
		wg.Wait()
	}
	b.StopTimer()
}

//测试一下关闭的channel是否可以被读取（关闭的channel可以被读取）
/*func Test_go(t *testing.T) {
	c := make(chan int, 10)
	for i := 1; i <= 3; i++ {
		c <- i
	}
	close(c)
	for v := range c {
		fmt.Println(v)
	}
}*/
