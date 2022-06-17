package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("Hello World.")
}

//结构体User的绑定方法
func (u User) CallReflect(name string, id, age int) {
	fmt.Printf("My name is %s, I am %d, My id is %d, ni hao %s\n", name, age, id, u.Name)
}

func main() {
	u := User{
		Id:   1,
		Name: "LPJ",
		Age:  2,
	}
	info(u)

	/*m := 123
	vOfValue := reflect.ValueOf(&m)
	vOfValue.Elem().SetInt(999)
	fmt.Println(m)*/

	user := User{2, "lpj", 18}
	user.CallReflect("jjy", 10, 20)
	//获取user的反射值类型
	v := reflect.ValueOf(user)
	//获取方法
	mehoodU := v.MethodByName("CallReflect")
	//设置参数
	args := []reflect.Value{reflect.ValueOf("YJJ"), reflect.ValueOf(3), reflect.ValueOf(5)}
	mehoodU.Call(args)

	a := 123
	reflectTest01(a)
}

//反射方法   传递一个空的接口
func info(o interface{}) {
	//使用反射获取这个空接口的类型集合
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t)

	//判断传入的类型是否正确
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("传入了错误的类型")
		return
	}

	//使用反射获取这个空接口的值集合
	v := reflect.ValueOf(o)
	fmt.Println("Value:", v)
	//Value可以通过Type()方法获取对应的Type对象
	fmt.Println("Second Type:", v.Type())

	fmt.Println("num:", t.NumField(), "Field:", t.Field(0))

	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("%6s:%v => %v\n", t.Field(i).Name, t.Field(i).Type, v.Field(i).Interface())
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s:%v\n", m.Name, m.Type)
	}
}

//这个方法对基本数据类型进行反射的基本操作
func reflectTest01(b interface{}) {

	//获取传递的这个空接口的类型
	rType := reflect.TypeOf(b)
	fmt.Println("type:", rType)
	//获取值
	rVal := reflect.ValueOf(b)
	fmt.Printf("type:%T, Val:%d\n", rVal, rVal)

	//将这个值转换为int
	n1 := rVal.Int()
	fmt.Println(n1)

	//将这个值转换为空接口
	iV := rVal.Interface()
	fmt.Printf("type:%T, val:%v\n", iV, iV)

	//将空接口通过断言的方式转换成需要的类型
	n2 := iV.(int)
	fmt.Println("n2:", n2)

}
