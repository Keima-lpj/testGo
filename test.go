package main

import(
	"fmt"
)

const(
	BIG = 1 << 100
	SMALL = BIG >> 99
)

func needInt (x int) int {
	return x * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main(){
	fmt.Println(needInt(SMALL))
	fmt.Println(needFloat(BIG))
	fmt.Println(needFloat(SMALL))

	s1 := make([]int, 0)
	printfSlice("s1", s1)
	s2 := make([]int, 0, 5)
	s2 = append(s2, 1, 2, 3, 4, 5)
	printfSlice("s2", s2)
	c := s2[:2]   //0,1
	printfSlice("c", c)
	d := c[2:5]   //0,0,0
	printfSlice("d", d)

/*	s3 := [10]int{1, 2, 3, 4, 5}

	s4 := s3[1:]
	printfSlice("s4", s4)*/


}

func printfSlice(s string, x []int){
	fmt.Printf("%s, len = %d, cap = %d, %v, %p\n", s, len(x), cap(x), x, x)
}
