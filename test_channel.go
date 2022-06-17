package main

import "fmt"

type Cheese struct {
	Name string
	Age  int
}

func main() {
	//testNoCacheChannel()
	//testCacheChannel()
	//testAllChannel()

	testDemo()
}

func testNoCacheChannel() {

	intChan := make(chan int)

	fmt.Printf("intChan的值：%v, intChan本身的地址：%p\n", intChan, &intChan)

	go func() {
		fmt.Println(<-intChan)
	}()

	//向这个intChan中写入数据
	//可以看到，如果在主线程中写入数据，且写入数据之前没有任何其他的goroutine取出数据的话，会造成deadlock。
	//如果上方的匿名函数（用于取出channel中的数据）写在写入数据的下方，则同样会deadlock
	intChan <- 10
}

func testCacheChannel() {
	//声明一个三个容量的缓存channel
	intChan := make(chan int, 3)

	//向这个channel中存入数据
	intChan <- 100
	intChan <- 200
	intChan <- 300
	//intChan <- 500   这里由于超出了缓冲channel的容量，会提示deadlock

	//这里从缓冲channel中取出对应的数据
	fmt.Println(<-intChan)
	fmt.Println(<-intChan)
	fmt.Println(<-intChan)

	//这里实验，如果在缓冲信道上面给一匿名函数取出数据，缓冲信道可以超出容量吗
	go func() {
		fmt.Println(<-intChan)
	}()

	intChan <- 100
	intChan <- 200
	intChan <- 300
	intChan <- 500

	//这里显式的关闭channel,不允许再写入channel,否则在下面的range会导致deadlock
	close(intChan)
	//很显然是可以的,这里intChan的值为200,300,500
	for v := range intChan {
		fmt.Println(v)
	}
}

func testAllChannel() {
	//声明一个接受一切数据类型的channel
	allChan := make(chan interface{}, 5)

	allChan <- 8
	allChan <- "aaa"
	allChan <- true
	allChan <- Cheese{Name: "miaomiao", Age: 15}
	allChan <- &Cheese{Name: "miaomiao", Age: 15}

	close(allChan)

	/*for v := range allChan {
		fmt.Println(v)
	}*/

	x := <-allChan
	<-allChan
	<-allChan
	cheese1 := <-allChan

	x1, ok := x.(int)
	if ok {
		fmt.Println(x1)
	} else {
		fmt.Println("no！")
	}

	//这个时候取出的虽然是一个结构体的对象，但是不能直接调用其中的属性和方法
	//需要先使用类型断言来判断
	cheese2, ok := cheese1.(Cheese)
	if ok {
		fmt.Println(cheese2.Name)
	}

	//这里试试看往关闭的channel中再次写入数据，会发生什么
	//会提示send on closed channel
	//allChan <- false
}

func testDemo() {
	intChan := make(chan int, 50)
	stopChan := make(chan bool)
	go func() {
		for i := 0; i <= 49; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	go func() {
		for v := range intChan {
			fmt.Println(v)
		}
		stopChan <- true
	}()

	<-stopChan
}
