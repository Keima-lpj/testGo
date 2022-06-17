package main

import "fmt"

type animal struct {
	run string
}

type catOne struct {
	animal *animal
}

type cheeseCat struct {
	cat *catOne
}

func newAnimal(run string) *animal {
	return &animal{
		run: run,
	}
}

func newCatOne(animal *animal) *catOne {
	return &catOne{
		animal: animal,
	}
}

func newCheeseCat(catOne *catOne) *cheeseCat {
	return &cheeseCat{
		cat: catOne,
	}
}

func (cat *cheeseCat) run() {
	fmt.Println(cat.cat.animal.run)
}

func main() {
	/*animal := newAnimal("run!")
	catOne := newCatOne(animal)
	cheese := newCheeseCat(catOne)
	cheese.run()*/
	cheese := InitializeCheeseCat("run")
	cheese.run()
}
