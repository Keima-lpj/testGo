package client_server

import (
	"fmt"
	"io"
	"net"
	"strings"
)

var connList map[string]net.Conn

//服务端的监听程序
func ServerRun() {

	fmt.Println("我是服务端，我开始监听了!")

	listen, err := net.Listen("tcp", "0.0.0.0:8888")

	if err != nil {
		fmt.Println("listen err = ", err)
	}

	connList = make(map[string]net.Conn)

	for {
		fmt.Println("等待客户端连接...")

		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("conn err = ", err)
			continue
		} else {
			ip := conn.RemoteAddr().String()
			fmt.Println(fmt.Sprintf("客户端连接成功ip=%s", ip))
			//将这个conn放入我们的一个map中
			connList[ip] = conn
		}
		//这里开一个goroutine来处理这个链接
		go dealConn(conn)
	}
}

//这里是处理连接请求的方法
func dealConn(conn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("连接报错了! err=", err)
			conn.Close()
		}
	}()

	fromIp := conn.RemoteAddr().String()

	for {
		fmt.Println("服务器在等待客户端发送信息...", fromIp)
		//声明一个1024长度的字节切片
		message := make([]byte, 1024)
		n, err := conn.Read(message)
		if err == io.EOF {
			//如果错误等于io.EOF，则证明客户端也退出了
			fmt.Println("客户端：", fromIp, "读取ip错误， err=", err)
			return
		}
		//用字符串分隔信息
		messageString := string(message[:n-1])
		//用分号截取
		messageSlice := strings.Split(messageString, ";")

		fmt.Println(connList)

		//如果这个ip在connlist存在，则将接下来的这个消息发送给对应的客户端连接
		if _, ok := connList[messageSlice[0]]; ok {
			//将信息发送给客户端想要发送的连接
			targetConn := connList[messageSlice[0]]
			targetIp := targetConn.RemoteAddr().String()

			_, targetErr := targetConn.Write([]byte(messageSlice[1]))
			if targetErr != nil {
				fmt.Println("给客户端：", targetIp, "发送信息错误， err=", err)
				continue
			}
			fmt.Println("客户端A：", fromIp, "给客户端B：", targetIp, "发送信息成功！ 信息:", messageSlice[1])

		} else {
			errMessage := []byte("您发送的ip不存在或者已下线")
			_, err := conn.Write(errMessage)
			if err != nil {
				fmt.Println("客户端：", fromIp, "断线了!")
				return
			}
			fmt.Println("给客户端回复了信息！")
		}

	}
}
