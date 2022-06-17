package main

import (
	"github.com/LiaoPuJian/keima/spider"
)

//这个是一个爬虫程序，用于爬取一拳超人的漫画。

func main() {
	chapter := make([]int, 2)
	chapter[0] = 1
	chapter[1] = 25
	spider.OnePunchRun(0)
}
