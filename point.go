package main

import (
	"fmt"
)

func main() {
	var a, b = 40, 80
	var p = &a
	fmt.Println(*p)
	fmt.Println(p)
	*p = 20
	fmt.Println(a)

	p = &b
	*p = *p / 20
	fmt.Println(b)
}

