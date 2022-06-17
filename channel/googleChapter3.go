package main

import (
	"fmt"
	"sync"
)

type atomic struct {
	value int
	lock  sync.Mutex
}

func (a *atomic) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()

	//tips:如果想要加锁的操作对一个代码块起作用，那么就使用匿名函数就好了

	a.value++
}

func (a *atomic) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a atomic
	a.increment()
	go func() {
		a.increment()
	}()
	fmt.Println(a.get())
}
