package main

import (
	"fmt"
	project1 "testGo/project01"
)

func main() {

	//声明一个账户实例
	account := project1.Account{
		AccountNo: "111111",
		Pwd:       "123456",
		Balance:   100,
	}

	var accountNo string
	//从控制台读取操作
	fmt.Println("登录请输入账户名:")
	fmt.Scanln(&accountNo)

	if accountNo != account.AccountNo {
		fmt.Println("您输入的账户名有误")
		return
	}

	var method string

	var amount float64

	var pwd string

	for {
		fmt.Println("请输入您需要的操作：1存款，2取款，3查询，4退出")
		fmt.Scanln(&method)

		switch method {
		case "1":
			fmt.Println("请输入存款金额:")
			fmt.Scanln(&amount)
			fmt.Println("请输入密码:")
			fmt.Scanln(&pwd)
			account.Deposit(amount, pwd)
		case "2":
			fmt.Println("请输入取款金额:")
			fmt.Scanln(&amount)
			fmt.Println("请输入密码:")
			fmt.Scanln(&pwd)
			account.Withdrawal(amount, pwd)
		case "3":
			fmt.Println("请输入密码:")
			fmt.Scanln(&pwd)
			account.Query(pwd)
		case "4":
			fmt.Println("再见")
			return
		}
	}
}
