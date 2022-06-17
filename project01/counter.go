package project1

import "fmt"

type Account struct {
	AccountNo string  //账号
	Pwd       string  //密码
	Balance   float64 //余额
}

/**
账户存款
*/
func (a *Account) Deposit(amount float64, pwd string) {
	//判断密码是否正确
	if pwd != a.Pwd {
		fmt.Println("你输入的密码有误")
		return
	}
	//判断输入的金额是否正确
	if amount <= 0 {
		fmt.Println("你输入的金额有误")
		return
	}
	//加余额
	a.Balance += amount
	fmt.Println("存款成功")
}

/**
账户取款
*/
func (a *Account) Withdrawal(amount float64, pwd string) {
	//判断密码是否正确
	if pwd != a.Pwd {
		fmt.Println("你输入的密码有误")
		return
	}
	//判断输入的金额是否正确
	if amount <= 0 || amount > a.Balance {
		fmt.Println("你输入的金额有误")
		return
	}
	//加余额
	a.Balance -= amount
	fmt.Println("取款成功")
}

/**
查询信息
*/
func (a *Account) Query(pwd string) {
	//判断密码是否正确
	if pwd != a.Pwd {
		fmt.Println("你输入的密码有误")
		return
	}
	//打印用户信息和余额
	fmt.Printf("您的账号为：%v，余额为：%v\n", a.AccountNo, a.Balance)
}
