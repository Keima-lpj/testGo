package main

import (
	"fmt"
)

func main() {

	/*	fmt.Println("start")

		a := 1
		b := "a"
		c := "c"

		d := A(a, b, c)
		fmt.Println(d)
		fmt.Println(C())

		x := func() {
			fmt.Println("I am ni ming han shu")
		}
		x()*/

	/*	a := closure(5)
		fmt.Println(a(1))
		fmt.Println(a(2))

		for i := 0; i < 3; i++ {
			defer func() {
				fmt.Println(i)
			}()
		}

		X()
		Y()
		Z()*/

	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i) //第一次入栈   i = 0
		defer func() {
			fmt.Println("defer_closure i = ", i) //第二次入栈   i = 0
		}()
		fs[i] = func() {
			fmt.Println("closure i = ", i)
		}
	}

	for _, f := range fs {
		f()
	}
}

func A(a int, b, c string) int {
	return 1
}

func C() (a, b, c int) {
	a, b, c = 1, 2, 3
	return a, b, c
}

func D(a ...int) {
	fmt.Println(a)
}

func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func X() {
	fmt.Println("func X")
}

func Y() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in Y")
		}
	}()
	panic("panic in Y")
}

func Z() {
	fmt.Println("func Z")
}
