package main

import (
	"fmt"
	"strings"
)


//切片
func main() {
	s := []int{2, 4, 6, 8, 9}
	fmt.Println("s ==", s)

	fmt.Println("s[1:4] == ", s[0:4])

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	game := [][]string{
		[]string{},
		[]string{},
		[]string{},
	}

	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"

	printBoard(game)

	var z []int
	fmt.Println(z, len(z), cap(z))

	printSlice("a", z)

	z = append(z, 0)

	printSlice("a", z)

	z = append(z, 1)

	printSlice("a", z)

	pow := []int{1, 2, 4, 8, 16, 32}

	fmt.Printf("\n")

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}


func printBoard(s [][]string){
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}

func printSlice(s string, x []int){
	fmt.Printf("%s len = %d, cap = %d $v\n", s, len(x), cap(x), x)
}
