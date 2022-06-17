package main

import "fmt"

func main() {
	c := make(chan bool, 100)

	for i:=1; i<=101; i++ {
		go func() {
			c <- true
			fmt.Println(1)
		}()
	}

	for i:=1; i<=100; i++ {
		<-c
	}
}
