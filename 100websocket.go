package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"golang.org/x/sys/unix"
	"log"
	"net"
	"net/http"
	"reflect"
	"sync"
	"syscall"
)

/*
如果我们的服务需要面向100w个用户时，会发生什么情况？
1.每创建一个websocket连接，按照以上的实现方式，我们就需要创建一个goroutine来接收客户端的信息。一个goroutine大概需要2~8kb的内存
2.如果是同时有100万个连接，假设每个goroutine占用4kb内存，那么内存消耗大概在：4kb*1000000=4G。
光是保持连接，不做任何处理就已经消耗了4G的内存，还是挺恐怖的，所以下面开始介绍用epoll模型来解决这个问题。

Epoll
epoll是Linux内核为处理大批量文件描述符而作了改进的poll，是Linux下多路复用IO接口select/poll的增强版本，它能显著提高程序在大量并发连接中只有少量活跃的情况下的系统CPU利用率。
Epoll提供了3个方法：Create、Ctl、Wait

Create: 创建epoll句柄，返回文件标识符(fd)。
Ctl: 根据epoll的fd，完成注册事件、删除事件、更新事件。
Wait: 返回就绪事件。
我们可以通过epoll模型，来管理websocket连接，用来替代通过goroutine去监听的方案。
实现如下：
*/

func main() {
	// 创建epoll
	epollFd, err := unix.EpollCreate(1)
	if err != nil {
		log.Fatalf("unable to create epoll:%v\n", err)
	}
	//创建一个存储100万链接的map
	connections := make(map[int32]*websocket.Conn, 1000000)
	var mu sync.Mutex
	ch := make(chan int32, 50)
	// 开启协程，处理接受到的信息
	go func() {
		for {
			select {
			case fd := <-ch:
				conn := connections[fd]
				if conn == nil {
					continue
				}
				// 接收消息
				_, message, err := conn.ReadMessage()
				if err != nil {
					log.Println("unable to read message:", err.Error())
					_ = conn.Close()
					// 删除epoll事件
					if err := unix.EpollCtl(epollFd, syscall.EPOLL_CTL_DEL, int(fd), nil); err != nil {
						log.Println("unable to remove event")
					}
				}
				// 给客户端回消息
				if err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("receive:%s", string(message)))); err != nil {
					log.Println("unable to send message:", err.Error())
				}
			}
		}
	}()
	// 开启一个协程监听epoll事件
	go func() {
		for {
			// 声明50个events，表明每次最多获取50个事件
			events := make([]unix.EpollEvent, 50)
			// 每100ms执行一次
			n, err := unix.EpollWait(epollFd, events, 100)
			if err != nil {
				log.Println("epoll wait:", err.Error())
				continue
			}
			// 取出来的是就绪的websocket连接的fd
			for i := 0; i < n; i++ {
				if events[i].Fd == 0 {
					continue
				}
				ch <- events[i].Fd // 通过channel传递到另一个goroutine处理
			}
		}

	}()
	// 绑定http服务
	http.ListenAndServe(":8085", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 升级为websocket连接
		conn, err := NewWebsocketConnection(w, r)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		// 获取文件标识符
		fd := GetSocketFD(conn.UnderlyingConn())
		// 注册事件
		if err := unix.EpollCtl(epollFd,
			unix.EPOLL_CTL_ADD,
			int(fd),
			&unix.EpollEvent{Events: unix.POLLIN | unix.POLLHUP, Fd: fd}); err != nil {
			log.Println("unable to add event:%v", err.Error())
			return
		}
		// 保存到map里
		mu.Lock()
		connections[fd] = conn
		mu.Unlock()
	}))
}

var u = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }} // use default options

// NewWebsocketConnection return web socket connection
func NewWebsocketConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return u.Upgrade(w, r, nil)

}

// GetSocketFD get socket connection fd
func GetSocketFD(conn net.Conn) int32 {
	tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")
	return int32(pfdVal.FieldByName("Sysfd").Int())
}
