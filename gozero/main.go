package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/mr"
	"log"
)

func main() {
	val, err := mr.MapReduce(func(source chan<- interface{}) {
		// generator
		for i := 0; i < 10; i++ {
			source <- i
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		// mapper
		i := item.(int)
		writer.Write(i * i)
	}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
		// reducer
		var sum int
		for i := range pipe {
			sum += i.(int)
		}
		writer.Write(sum)
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", val)
}
