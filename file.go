package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	BufReader()

	//NoBufReader()

	//WriteString()

	//putFileToAnotherFile()

	/*filename1 := "e:/111.jpg"
	filename2 := "e:/1111.jpg"
	copyFileToAnotherFile(filename1, filename2)

	filename1 = "e:/111.jpg"
	filename2 = "e:/1111.txt"
	copyFileToAnotherFile(filename1, filename2)*/
}

/**
带缓冲区的读取文件，适用于很大文件的情况
*/
func BufReader() {
	//打开一个文件
	file, err := os.Open("e:/test.txt")
	if err != nil {
		fmt.Println("打开文件错误：", err)
		return
	}

	//当函数退出时自动关闭文件。如果不及时关闭句柄，会存在内存泄漏
	defer file.Close()

	//输出file的类型  这里打印出是一个 *os.File的指针类型
	fmt.Printf("打开的文file类型%T\n", file)

	//创建一个带缓冲的*Reader
	reader := bufio.NewReader(file)

	//循环的读取文件的内容
	for {
		str, end := reader.ReadString('\n')
		//输出内容
		fmt.Println(str)
		//如果读取到了文件的末尾，则跳出循环
		if end == io.EOF {
			break
		}
	}
	fmt.Println("文件读取结束")
}

/**
使用ioutil来操作，适用于文件不大的情况
*/
func NoBufReader() {
	content, err := ioutil.ReadFile("E:/test.txt")
	if err != nil {
		fmt.Println("读取文件失败:", err)
	}
	fmt.Println(string(content))

	//下面这种方式也可以
	/*file, err := os.Open("e:/test.txt")
	if err != nil {
		fmt.Println("打开文件错误：", err)
		return
	}
	defer file.Close()
	content1, _ := ioutil.ReadAll(file)
	fmt.Println(string(content1))*/
}

func WriteString() {
	//如果第二个参数使用O_APPEND，则会在原先的内容上追加
	file, err := os.OpenFile("E:/test.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close()
	str := "hello L!\n"
	//使用带缓存的写入
	write := bufio.NewWriter(file)
	//循环写入五遍
	for i := 0; i < 5; i++ {
		write.WriteString(str)
	}
	write.Flush()
}

/**
将一个文件的内容写入到另一个文件
*/
func putFileToAnotherFile() {
	//将一个文件的内容导入到另一个文件的内容
	file1 := "e:/test.txt"
	file2 := "e:/test1.txt"

	data, err := ioutil.ReadFile(file1)

	if err != nil {
		fmt.Println("打开失败：", err)
		return
	}

	err = ioutil.WriteFile(file2, data, 0666)

	if err != nil {
		fmt.Println("写入失败", err)
	}
	return
}

/**
将一个文件路径copy到另一个文件
*/
func copyFileToAnotherFile(filename1, filename2 string) {
	//首先打开文件1的句柄
	file1, err := os.Open(filename1)
	if err != nil {
		fmt.Println("打开文件1报错", err)
		return
	}
	defer file1.Close()

	//创建一个文件1的读取流
	reader := bufio.NewReader(file1)

	//然后打开文件2的句柄
	file2, err := os.OpenFile(filename2, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("打开文件2报错", err)
		return
	}
	defer file2.Close()

	//创建一个文件2的写入流
	write := bufio.NewWriter(file2)

	//将文件1的内容写入到文件2中
	io.Copy(write, reader)

	//刷新 这里如果是写入txt文件，则一定要加，如果是图片可以不用加
	write.Flush()
}
