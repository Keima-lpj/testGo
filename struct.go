package main

import (
	"errors"
	"fmt"
)

type Point struct {
	px float32
	py float32
}

type MethodUtils struct {
}

type Calculator struct {
}

type Dog struct {
	name   string
	age    int
	weight float64
}

type cat struct {
	name string
	sex  int
}

func (dog *Dog) say() {
	fmt.Printf("dog的信息：name=%v, age=%v, weight=%v\n", dog.name, dog.age, dog.weight)
}

func (p *Point) setXY(px, py float32) {
	p.px = px
	p.py = py
}

func (p *Point) getXY() (float32, float32) {
	return p.px, p.py
}

/**
这个方法里打印10*8的矩形
*/
func (m *MethodUtils) print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/**
这里接收参数打印矩形
*/
func (m *MethodUtils) printX(x, y int) {
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/**
这里计算面积
*/
func (m *MethodUtils) getArea(len, width float64) float64 {
	return len * width
}

/**
判断奇数还是偶数
*/
func (m *MethodUtils) judgeNum(x int) {
	if x%2 == 0 {
		fmt.Println("是偶数")
	} else {
		fmt.Println("是奇数")
	}
}

func (m *MethodUtils) chengfakoujue() {
	var x int
	fmt.Scanln(&x)

	for i := 1; i <= x; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d ", j, i, i*j)
		}
		fmt.Println()
	}
}

func (m *MethodUtils) daozhishuzu() {
	var arr = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < i; j++ {
			arr[i][j], arr[j][i] = arr[j][i], arr[i][j]
		}
	}
	fmt.Println(arr)
}

func (c *Calculator) cal(x, y float64, method string) (float64, error) {
	switch method {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		return x / y, nil
	default:
		return 0, errors.New("输入的方法有误")
	}
}

/*func main() {
	//var mu MethodUtils
	var mu = MethodUtils{}
	//mu.print()
	//mu.printX(5,6)
	area := mu.getArea(5, 6)
	fmt.Println(area)

	var c Calculator
	res, _ := c.cal(6, 3, "+")
	fmt.Println(res)

	res2, err := c.cal(6, 3, "a")
	fmt.Println(res2)
	fmt.Println(err)

	//mu.chengfakoujue()
	mu.daozhishuzu()

	dog := Dog{
		name:   "feifei",
		age:    10,
		weight: 15.6,
	}

	dog.say()
	//声明一个账户结构体
	account := project1.Account{
		AccountNo: "123456",
		Pwd:       "111111",
		Balance:   100,
	}

	//存一百块
	account.Deposit(100, "111111")
	//取一百五
	account.Withdrawal(150, "111111")
	//查询余额
	account.Query("111111")

}*/
