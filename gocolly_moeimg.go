package main

import (
	"fmt"
	"github.com/LiaoPuJian/keima/spider"
)

func main() {
	start := 0
	end := 0
	fmt.Println("请输入您想爬取的起始页（数字）:")
	fmt.Scanln(&start)
	fmt.Println("请输入您想爬取的结束页（数字）:")
	fmt.Scanln(&end)

	if start <= 0 || end <= 0 || start > end {
		fmt.Println("输入的页码有错！")
		return
	}
	spider.MoeimgRun(start, end)
	fmt.Println("结束")
}
