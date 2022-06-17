package main

import (
	"fmt"
	"strings"
)

/**
闭包累加器
该函数返回了一个匿名函数，但这个匿名函数引用到了其之外的变量n
该匿名函数和n就形成了一个整体，构成了闭包

可以将闭包理解成一个类，变量是类里面的属性，匿名函数是类里面的方法
外部调用AddUpper这个函数之后，相当于获得了一个类
 */
func AddUpper() func(int) (int, string) {
	var n int = 10
	var str string = "hello"
	return func (x int) (int, string) {
		n = n + x
		str += string(x)
		return n, str
	}
}

func main (){
	f := AddUpper()

	fmt.Println(f(1))
	fmt.Println(f(2))

	f1 := makeSuffix(".jpg")

	filename := f1("myPic.jpg")
	fmt.Println(filename)
	filename2 := f1("myPic111")
	fmt.Println(filename2)
}


func makeSuffix(suffix string) func (string) string {

	return func(filename string) string {
		//判断，如果该文件名包含suffix，则直接返回，否则加上suffix再返回
		//可以用Contains函数，也可以用strings.HasSuffix()函数
		if strings.Contains(filename, suffix) {
			return filename
		}else{
			return filename + suffix
		}
	}
}
