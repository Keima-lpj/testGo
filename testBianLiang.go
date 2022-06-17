package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	var a int = 10

	b := 10

	c := 3.1415

	fmt.Printf("a的类型是%T， b的类型是%T， c类型是%T\n", a, b, c)

	var c1 int = '北'

	fmt.Printf("c1的类型是%T，值是%d\n", c1, c1)

	c2 := '北'

	c3 := "北"

	fmt.Printf("c2的类型是%T，值是%d，%c, c3的类型是%T，值是%d\n", c2, c2, c2, c3, c3)

	var c4 int = 97
	var c5 int = 22269
	c6 := 'a'
	//%c输出对应的unicode值
	fmt.Printf("c4对应的unicode是%c c5对应的unicode是%c c6对应的码值是%d\n", c4, c5, c6)

	var c7 bool = false
	fmt.Println(c7, "c7占用的字节数：", unsafe.Sizeof(c7))

	var s1 string = "my"

	s1 = "111"

	fmt.Println(s1)

	//这里开始是字符串和基本类型互相转换
	var x1 int = 8
	var x2 float32 = 123.2
	var x3 bool = true
	var str string

	//第一种方式，使用Sprintf
	str = fmt.Sprintf("%d", x1)
	fmt.Println(str)
	str = fmt.Sprintf("%f", x2)
	fmt.Println(str)
	str = fmt.Sprintf("%t", x3)
	fmt.Println(str)

	//第二种方式，使用strconv包提供的方法 ,base代表的是进制，10表示10进制
	str = strconv.FormatInt(int64(x1), 2)
	fmt.Println(str)
	//'f'代表格式 prec代表精度
	str = strconv.FormatFloat(float64(x2), 'f', 10, 64)
	fmt.Println(str)
	str = strconv.FormatBool(x3)
	fmt.Println(str)

	str = "123"

	//string转基本类型
	z1, _ := strconv.ParseInt(str, 10, 8)
	fmt.Println(z1)
	z2, _ := strconv.ParseFloat(str, 64)
	fmt.Println(z2)
	z3, _ := strconv.ParseBool(str)
	fmt.Println(z3)

	//Itoa 和 AtoI
	var zz int = 100
	str1 := strconv.Itoa(int(zz))
	fmt.Printf("str1的类型是%T， 值是%v", str1, str1)
	z4, _ := strconv.Atoi(str1)
	fmt.Printf("z4的类型是%T， 值是%d", z4, z4)

}
