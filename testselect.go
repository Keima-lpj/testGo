package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 100)
	quit := make(chan int)

	go func() {
		for i := 0; i <= 99; i++ {
			time.Sleep(time.Second)
			ch <- i
		}
	}()

	go func() {
		for {
			select {
			case v := <-ch:
				fmt.Println("get ", v)
			case <-quit:
				fmt.Println("stop!")
				return
			}
		}
	}()

	time.Sleep(time.Second * 10)
	close(quit)
	time.Sleep(time.Second * 90)
}
