package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8085"
	DELIMITER      = '\t'
)

//打印日志
func printLog(role string, sn int, format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf("%s[%d]: %s", role, sn, fmt.Sprintf(format, args))
}

//打印服务端日志
func printServerLog(format string, args ...interface{}) {
	printLog("Server", 0, format, args)
}

//打印客户端日志
func printClientLog(sn int, format string, args ...interface{}) {
	printLog("Client", sn, format, args)
}

func main() {

}

//启动服务端
func serverGo() {
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		printServerLog("Listen Error: %s", err)
	}
	defer listener.Close()
	printServerLog("Got listener for the server. (local address: %s)", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			printServerLog("Accept Error: %s", err)
		}
		printServerLog("Established a connection with a client application. (remote address: %s)", conn.RemoteAddr())
		go handleConn(conn)
	}
}

//处理连接
func handleConn(conn net.Conn) {
	for {
		//设置读取超时时间为10秒
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		strReq, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printServerLog("The connection is closed by another side.")
			} else {
				printServerLog("Read Error: %s", err)
			}
			break
		}
		printServerLog("Received request: %s", strReq)
	}
}

//从连接中读取数据
func read(conn net.Conn) (string, error) {
	//设置一次只读取一个字节的原因是，方式从连接中读出多余的数据对后续的读取操作造成影响，也可以判断分解符
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == DELIMITER {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}

func clientGo() {

}
