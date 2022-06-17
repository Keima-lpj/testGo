package main

import "github.com/google/wire"

func InitializeCheeseCat(msg string) *cheeseCat {
	wire.Build(newAnimal, newCatOne, newCheeseCat)
	return &cheeseCat{}
}
