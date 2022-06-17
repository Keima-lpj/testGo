package main

import (
	"fmt"
	"strconv"
)

func read(ch chan int) {
	value := <-ch
	fmt.Println("value:" + strconv.Itoa(value))
}

func write(ch chan int, x int) {
	if x == 10 {
		ch <- x
	}
}

func testChannel(ch chan int) {
	ch <- 1
	fmt.Println("testChannel test func 111")

	ch <- 1
	fmt.Println("testChannel test func 222")
}

func main() {
	ch := make(chan int, 1)
	//go read(ch)
	go write(ch, 10)
	//ch <- 1

	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		default:
			fmt.Println("111")
			break
		}
	}

	fmt.Println(<-ch)
	fmt.Println("end")
	/*ch := make(chan int, 1)
	go testChannel(ch)

	fmt.Println("main func")
	<-ch

	time.Sleep(time.Second * 2)*/

}
