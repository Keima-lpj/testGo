package main

import (
	"fmt"
)

//使用ch <- v发送一个值v到channel。发送值到channel可能会有多种结果，即可能成功，也可能阻塞，甚至还会引发panic，取决于当前channel在什么状态。
//使用 v, ok <- ch 接收一个值。第二个遍历ok是可选的，它表示channel是否已关闭。接收值只会又两种结果，要么成功要么阻塞，而永远也不会引发panic。

//总结来看，为什么会死锁？非缓冲信道上如果发生了流入无流出，或者流出无流入，也就导致了死锁。
// 或者这样理解，Go启动的所有goroutine里的非缓冲信道一定要一个线里存数据，一个线里取数据，要成对才行 。

//缓冲channel:c可以缓存一定的数据。也就是说，放入一个数据，c并不会挂起当前线, 再放够容量才会挂起当前线直到第一个数据被其他goroutine取走.
// 也就是只阻塞在容量一定的时候，不达容量不阻塞。  缓冲信道会在满容量的时候加锁

func main() {
	//声明一个无缓冲channel
	//ch := make(chan int)
	//直接往里面塞一个值  这个时候会直接报panic，因为前面没有任何的goroutine在读取这个channel
	/*ch <- 1
	<-ch*/

	//直接往里面塞一个值，后面加一个goroutine来读取数据。此时也会直接报panic，理由同上
	/*ch <- 1
	go func() {
		fmt.Println(<-ch)
	}()*/

	//声明一个goroutine,里面从ch中读取一个值。此时这个goroutine由于暂时取不到值，被阻塞，等到将值放入到ch中后，此goroutine继续执行
	/*go func() {
		fmt.Println(<-ch)
	}()
	ch <- 1*/

	//阻塞主goroutine，新开一个goroutine往ch中写入数据，此时ch中写入数据之后，主goroutine才会继续执行
	/*go func() {
		time.Sleep(time.Second)
		ch <- 1
		fmt.Println(222)

	}()
	<-ch
	fmt.Println(111)*/

	//如果开一个goroutine，往无缓冲信道中存入一个数据，主函数中没有另外的goroutine往出取，此时并不会panic，只是这个goroutine会一直阻塞，主函数会继续执行直到结束
	/*go func() {
		time.Sleep(time.Second)
		ch <- 1
		fmt.Println(222)
	}()*/

	//上面的例子可以说明，如果想要某个goroutine阻塞，只需要在这个goroutine中获取无缓冲channel的值即可。

	//下面是for-range   for-range语法可以用到通道上。循环会一直接收channel里面的数据，直到channel关闭。 如果不关闭channel，则会panic
	//注意，for-range对应的channel不能是只写channel。
	/*go func() {
		ch <- 1
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}*/

	//声明一个缓冲信道，可以存放5个值
	ch1 := make(chan int, 5)
	//往其中存五个值，并不会panic
	/*for i := 0; i < 5; i++ {
		ch1 <- i
	}
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)*/

	//往其中存6个值，则会直接panic
	/*for i := 0; i < 6; i++ {
		ch1 <- i
	}*/

	/*go func() {
		fmt.Println("副线程开始取")
		//此时阻塞当前goroutine,等待有数据流入ch1
		fmt.Println(<-ch1)
		fmt.Println(<-ch1)
		fmt.Println(<-ch1)
		fmt.Println(<-ch1)
		fmt.Println(<-ch1)
	}()
	for i := 0; i < 10; i++ {
		//这里存入一个，上面的goroutine则取走一个，一共取走五个值，则从5到9还留在ch1中
		ch1 <- i
	}
	fmt.Println("主线程开始取")
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)*/

	fmt.Println("我好了")
}
