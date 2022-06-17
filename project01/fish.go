package project1

import "fmt"

type Fish struct {
	Name string
}

func FishSay() {
	fmt.Println("I is fish")
}

func (f *Fish) FishAge() {
	age := getInt()
	fmt.Printf("I am %d age old", age)
}
