package client_server

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func ClientRun() {

	fmt.Println("我是客户端，我要开始连接服务端了...")

	conn, err := net.Dial("tcp", "127.0.0.1:8888")

	if err != nil {
		fmt.Println("连接失败，err=", err)
	}

	defer conn.Close()

	//连接成功，从终端获取一个信息，发送给服务端
	reader := bufio.NewReader(os.Stdin) //os.Stdin代表从标准终端读取

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("输入错误，err=", err)
		}

		if line == "exit" {
			fmt.Println("退出终端")
			return
		}

		//将读取的行发给服务端
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("发送数据错误, err=", err)
			break
		}
		fmt.Println(fmt.Sprintf("发送数据成功，成功的字节数：%d", n))

		//读取从服务端获取的信息
		message := make([]byte, 1024)
		getN, getErr := conn.Read(message)
		if getErr != nil {
			fmt.Println("从服务端读取数据报错了！！")
		}
		fmt.Println(string(message[:getN]))

	}

}
