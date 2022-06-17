package project1

import "fmt"

type Control struct {
}

/**
客户结构体
*/
type Customer struct {
	Id    int
	Name  string
	Sex   string
	Age   int
	Phone string
	Email string
}

/**
客户结构体的切片
*/
var customerSlice = make([]Customer, 10)

var incrementId int = 1

/**
显示主菜单
*/
func (con *Control) MainMenu() {

	var method int

	for {
		fmt.Println("-------------客户信息管理软件------------")
		fmt.Println()
		fmt.Println("               1 添加客户")
		fmt.Println("               2 修改客户")
		fmt.Println("               3 删除客户")
		fmt.Println("               4 客户列表")
		fmt.Println("               5 退   出")
		fmt.Print("               请选择(1-5):")
		fmt.Scanln(&method)

		//根据传递的method调用不同的方法
		switch method {
		case 1:
			con.AddCustomer()
		case 2:
			con.UpdateCustomer()
		case 3:
			con.DelCustomer()
		case 4:
			con.CustomerList()
		case 5:
			break
		}
	}
}

/**
添加用户
*/
func (con *Control) AddCustomer() {

	var customer = Customer{Id: incrementId}

	fmt.Println("-------------添加客户------------")
	fmt.Print("姓名: ")
	fmt.Scanln(&customer.Name)
	fmt.Print("性别: ")
	fmt.Scanln(&customer.Sex)
	fmt.Print("年龄: ")
	fmt.Scanln(&customer.Age)
	fmt.Print("电话: ")
	fmt.Scanln(&customer.Phone)
	fmt.Print("邮箱: ")
	fmt.Scanln(&customer.Email)
	fmt.Println("-------------添加完成------------")

	//将这个客户放入客户切片中
	customerSlice[incrementId] = customer
	//自增id加1
	incrementId += 1
}

/**
更新客户
*/
func (con *Control) UpdateCustomer() {

	var updateId, age int
	var name, sex, phone, email string

	fmt.Println("-------------修改客户------------")
	fmt.Print("请选择待修改客户编号(-1退出):")
	fmt.Scanln(&updateId)

	if updateId == -1 {
		return
	}
	//判断当前输入的id是否有效
	if updateId >= len(customerSlice) {
		fmt.Printf("输入的客户编号有误")
		return
	}

	fmt.Printf("姓名(%v):", customerSlice[updateId].Name)
	fmt.Scanln(&name)
	fmt.Printf("性别(%v):", customerSlice[updateId].Sex)
	fmt.Scanln(&sex)
	fmt.Printf("年龄(%v):", customerSlice[updateId].Age)
	fmt.Scanln(&age)
	fmt.Printf("电话(%v):", customerSlice[updateId].Phone)
	fmt.Scanln(&phone)
	fmt.Printf("邮箱(%v):", customerSlice[updateId].Email)
	fmt.Scanln(&email)

	if name != "" {
		customerSlice[updateId].Name = name
	}
	if sex != "" {
		customerSlice[updateId].Sex = sex
	}
	if age != 0 {
		customerSlice[updateId].Age = age
	}
	if phone != "" {
		customerSlice[updateId].Phone = phone
	}
	if email != "" {
		customerSlice[updateId].Email = email
	}

	fmt.Println("-------------修改完成------------")
}

/**
删除用户
*/
func (con *Control) DelCustomer() {
	var delId int
	var check string

	fmt.Println("-------------删除客户------------")
	fmt.Print("请选择待删除客户编号(-1退出):")
	fmt.Scanln(&delId)

	if delId == -1 {
		return
	}
	//判断当前输入的id是否有效
	if delId >= len(customerSlice) {
		fmt.Printf("输入的客户编号有误")
		return
	}
	fmt.Print("确认是否删除(Y/N):")
	fmt.Scanln(&check)

	//删除
	if check == "Y" || check == "y" {
		customerSlice = append(customerSlice[:delId], customerSlice[delId+1:]...)
		fmt.Println("-------------删除完成------------")
	}
	return
}

/**
客户列表
*/
func (con *Control) CustomerList() {
	fmt.Println("-------------客户列表------------")

	if len(customerSlice) == 0 {
		fmt.Println("当前暂无客户信息")
	}

	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")

	for _, v := range customerSlice {
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v \n",
			v.Id, v.Name, v.Sex, v.Age, v.Phone, v.Email)
	}

	fmt.Println("-------------客户列表完成------------")
}
