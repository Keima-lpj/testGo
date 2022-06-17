package main

import (
	"fmt"
	"net"
)

type Username struct {
	username string
	password string
}

var (
	userList = make([]Username, 2)
	connList = make([]net.Conn, 100)
)

func main() {
	user_1 := Username{
		username: "geziwang", password: "123456",
	}
	user_2 := Username{
		username: "maomi", password: "111111",
	}
	userList[0] = user_1
	userList[1] = user_2

	fmt.Println("服务端开始监听端口...")

	listen, err := net.Listen("tcp", "0.0.0.0:8080")

	defer listen.Close()

	if err != nil {
		fmt.Println("监听端口出错，", err)
		return
	}

	id := 0

	//这里启动goroutine，将连接进来的客户端放入goroutine中处理
	for {
		fmt.Println("等待客户端连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("接受客户端连接失败，", err)
			continue
		} else {
			fmt.Printf("接受客户端成功，con=%v, 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		//connList = append(connList, conn)
		connList[id] = conn
		id++
		go login(conn)
	}
}

func login(conn net.Conn) {

	defer conn.Close()

	is_login := 0
	var user Username

	for {
		//创建一个新的字节切片
		buf := make([]byte, 1024)

		conn.Write([]byte("请输入你的用户名\n"))
		//fmt.Printf("等待客户端%s发送信息\n", conn.RemoteAddr().String())
		n, readUsernameErr := conn.Read(buf)
		if readUsernameErr != nil {
			fmt.Printf("读取用户名error，客户端退出: %v", readUsernameErr)
			return
		}
		username := string(buf[:n-2])

		conn.Write([]byte("请输入你的密码\n"))
		//fmt.Printf("等待客户端%s发送信息\n", conn.RemoteAddr().String())
		n, readPwdErr := conn.Read(buf)
		if readPwdErr != nil {
			fmt.Printf("读取密码error，客户端退出: %v", readPwdErr)
			return
		}
		password := string(buf[:n-2])

		//判断用户名和密码是否在数据库中
		for _, v := range userList {
			if v.username == username && v.password == password {
				//登录成功，显示欢迎字符，同时给其他的用户打印该用户登录成功的信息
				_, replyErr := conn.Write([]byte("欢迎你，" + v.username + "\n"))

				if replyErr != nil {
					fmt.Printf("登录成功写入欢迎error，客户端退出: %v", readPwdErr)
					return
				}
				is_login = 1
				user = v
				break
			}
		}

		if is_login == 1 {
			break
		}

	}

	if is_login == 1 {
		//开始聊天
		chat(conn, user)
	} else {
		conn.Write([]byte("账号或者密码输错了，请重新输入！\n"))
	}

}

func chat(conn net.Conn, user Username) {
	defer conn.Close()

	for {
		//读取此链接的用户发送的信息并广播给其他用户
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("接受用户发送的信息报错了！", err)
			return
		}
		for _, v := range connList {
			if v == nil {
				break
			}
			_, err := v.Write(buf)
			if err != nil {
				fmt.Printf("广播给用户%s报错了,%v", v.RemoteAddr().String(), err)
				continue
			}
		}

	}
}
