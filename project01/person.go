package project1

import "fmt"

type person struct {
	Name string
	age  int
}

type student struct {
	person
}

type Worker struct {
	person
}

func (s *student) study() {
	fmt.Println("我是学生，我会学习")
}

func (w *Worker) Worker() {
	fmt.Println("我是工人，我要搬砖")
}
