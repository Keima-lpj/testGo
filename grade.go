package main

import (
	"fmt"
	"math/rand"
)

func main(){
	var a int
	var b string
	fmt.Println("请输入你的成绩:")
	fmt.Scan(&a)

	switch a{
		case 90: b = "小伙子不错"
		case 80: b = "小伙子一般般"
		case 70: b = "革命尚未成功，同志仍需努力"
		case 60: b = "你是智障吧?"
		default:
			b = "GGWP"
	}
	fmt.Println("对你的评价是:", b)

	fmt.Println("预测你下次的分数是:", rand.Intn(10))
}


