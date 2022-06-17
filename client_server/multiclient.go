package client_server

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func MultiClientRun() {
	syncNum := 1000
	for i := 1; i <= syncNum; i++ {
		go func() {
			conn, err := net.Dial("tcp", "127.0.0.1:8888")
			if err != nil {
				fmt.Println("连接失败，err=", err, conn.LocalAddr())
			}
			defer conn.Close()
			//连接成功，从终端获取一个信息，发送给服务端
			reader := bufio.NewReader(os.Stdin) //os.Stdin代表从标准终端读取
			for {
				_, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("输入错误，err=", err)
				}
			}
		}()
	}
}
