package main

import (
	"fmt"
	"regexp"
)

const text = `my email is liaopujian@qq.co
m@abc.com`

/**
正则表达式的基本语法。
.表示任意字符
+表示大于0个字符
*表示大于等于0个字符
.+表示匹配任意字符，但是数量必须大于0
.*表示匹配任意字符，数量可以为0
*/

func main() {
	//这个语句返回一个正则表达式的匹配器
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9.]+)`)
	//这里的语句的意思是，用这个睁着表达式的匹配器去匹配这个字符串，返回符合条件的子串串
	//注意，一个换行符只会匹配一个字符串
	match := re.FindString(text)
	fmt.Println(match)
	matchAll := re.FindAllString(text, -1)
	fmt.Println(matchAll)

	//这个语句可以提取匹配到的字串，要在表达式中对要提取的字串加上()
	subMatch := re.FindAllStringSubmatch(text, -1)
	fmt.Println(subMatch)
}
