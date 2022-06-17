package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connecter
}

type Connecter interface {
	Connecter()
}

type phoneConnecter struct {
	name string
	age  int
}

func (pc phoneConnecter) Name() string {
	return pc.name
}

func (pc phoneConnecter) Connecter() {
	fmt.Println("connecter")
}

func disconnecter(usb USB) {

	pc, ok := usb.(phoneConnecter)
	fmt.Println(pc, ok)

	if ok {
		fmt.Println(pc, "is a interface")
	} else {
		fmt.Println("is not a interface")
	}

	switch v := usb.(type) {
	case phoneConnecter:
		fmt.Println("is a interface", v)
	default:
		fmt.Println("unknow device")
	}
}

func main() {
	a := phoneConnecter{name: "phone"}
	name := a.Name()
	fmt.Println(name)

	a.Connecter()

	disconnecter(a)
}
