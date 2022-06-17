package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	intChan = make(chan int, 2)
)

func main() {
	go writeDataToFile("1")
	go getDataSort("1")

	<-intChan
	<-intChan
	fmt.Println("我完事了")
}

//这个协程随机生成1000个数据，存放到文件中
func writeDataToFile(n string) {
	fileName := fmt.Sprintf("%v.txt", n)
	fmt.Println("writeData文件名：", fileName)
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("writeData打开文件句柄失败!")
	}

	defer file.Close()
	rand.Seed(time.Now().Unix())
	write := bufio.NewWriter(file)

	for i := 1; i <= 1000; i++ {
		str := "100\n"
		fmt.Println("要写入的字符串：", str)
		//将这个随机数写入到文件中去
		write.WriteString(str)
	}
	write.Flush()

	intChan <- 1
}

func getDataSort(n string) {
	sortSlice := []int{}
	fileName := fmt.Sprintf("%v.txt", n)
	//读取文件n
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件句柄失败！", n)
	}
	defer file.Close()
	//创建一个带缓冲的*Reader
	reader := bufio.NewReader(file)
	//循环的读取文件的内容
	for {
		str, end := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		fmt.Println("我读到了！", str)
		num, _ := strconv.Atoi(str)
		//读到str,则将其转为int,放入到slice中
		sortSlice = append(sortSlice, num)
		//如果读取到了文件的末尾，则跳出循环
		if end == io.EOF {
			break
		}
	}
	//对这个slice进行排序，排序完成之后，循环放入到另一个文件下去
	maoPao1(&sortSlice)
	//循环这个数组,放到新的文件中去
	newFileName := fmt.Sprintf("%va.txt", n)
	for _, v := range sortSlice {
		str := string(v) + "\n"
		ioutil.WriteFile(newFileName, []byte(str), 0666)
	}
	intChan <- 1
}

func maoPao1(arr *[]int) {
	fmt.Println("排序前的数组：", *arr)
	var temp int

	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j+1]
				(*arr)[j+1] = (*arr)[j]
				(*arr)[j] = temp
			}
		}
	}
}
