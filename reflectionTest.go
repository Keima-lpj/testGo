package main

import (
	"fmt"
	"reflect"
)

//定义一个结构体
type MonsterRua struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score float32
	Sex   string
}

//定义两个方法给这个结构体
func (m MonsterRua) GetSum(n1, n2 int) int {
	return n1 + n2
}

//接收四个值，给结构体赋值
func (m MonsterRua) Set(name string, age int, score float32, sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}

//方法，显示结构体的值
func (m MonsterRua) Print() {
	fmt.Println("start~")
	fmt.Println(m)
	fmt.Println("end~")
}

//这个方法用于测试反射
func testReflection(m MonsterRua) {
	//打印该结构体的字段和方法，调用第一个和第二个方法
	typ := reflect.TypeOf(m)

	val := reflect.ValueOf(m)

	//获取m对应的值
	if val.Kind() != reflect.Struct {
		fmt.Println("传入的类型错误!")
		return
	}

	//获取该结构体有几个字段
	num := val.NumField()
	fmt.Println("结构体的字段数量：", num)

	//遍历获取结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Println("字段值：", val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")

		if tagVal != "" {
			fmt.Println("该字段的json tag是：", tagVal)
		}
	}

	//获取该结构体的方法
	numOfMethod := val.NumMethod()
	fmt.Println("结构体的方法数量：", numOfMethod)

	//调用结构体的第一个方法
	val.Method(1).Call(nil)

	for i := 0; i < numOfMethod; i++ {
		fmt.Println("方法：", val.Method(i))
	}

	//如果要给通过反射的结构体传参，需要传一个[]reflect.Value类型的切片
	var param []reflect.Value
	param = append(param, reflect.ValueOf(10))
	param = append(param, reflect.ValueOf(40))

	res := val.Method(0).Call(param)

	fmt.Println(res[0].Int())

}

func main() {
	//定义一个结构体
	huangshulang := MonsterRua{
		Name:  "黄鼠狼",
		Age:   8,
		Score: 50,
		Sex:   "公",
	}

	testReflection(huangshulang)
}
