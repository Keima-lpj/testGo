package main

import (
	"fmt"
	"github.com/LiaoPuJian/keima/spider"
)

func main() {

	//录入信息
	getType := 0
	fmt.Println("请输入您想爬取的标题（数字，以回车结束） 0自拍偷拍 1亚洲色图 2欧美色图 3美腿丝袜 4清纯唯美 5乱伦熟女 6卡通动漫 7极品美女")
	fmt.Scanln(&getType)
	start := 0
	end := 0
	fmt.Println("请输入您想爬取的起始页（数字，以回车结束）:")
	fmt.Scanln(&start)
	fmt.Println("请输入您想爬取的结束页（数字，以回车结束）:")
	fmt.Scanln(&end)

	spider.CatRun(getType, start, end)
	//spider.CatRun(3, 1, 1)

}
