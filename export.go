package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var (
	a = 3
	b = "no!"
	c = true
	d = cmplx.Sqrt(-5 + 12i)
)

func add (x, y int) int {
	return x + y
}

func main(){
	const f = "%T(%v)\n"
	fmt.Println(add(10, 20))
	fmt.Println(swap("a", "b"))
	fmt.Println(split(9))
	fmt.Printf(f, a, a)
	fmt.Printf(f, b, b)
	fmt.Printf(f, c, c)
	fmt.Printf(f, d, d)

	var x, y int = 3, 4
	var e = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(e)
	fmt.Println(x, y, z)

}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (int, int) {
	var x,y int
	x = sum * 4 / 9
	y = sum - x
	return x, y
}



