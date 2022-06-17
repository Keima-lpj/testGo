package main

import (
	"fmt"
)

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
)

func main() {

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	switch a := 2; {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough
	case a >= 1:
		fmt.Println("a>=1")
		fallthrough
	default:
		fmt.Println("undefined a")
	}

LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			break LABEL
		}
	}
}
