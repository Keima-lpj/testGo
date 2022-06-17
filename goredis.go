package main

import (
	"fmt"
	. "github.com/xuyu/goredis"
)

func main() {
	client, err := DialURL("tcp://auth:@127.0.0.1:6379/0?timeout=10s&maxidle=1")
	if err != nil {
		fmt.Println("链接redis错误")
	}

	err = client.Set("test_go_redis", "111", 0, 0, false, false)

	if err != nil {
		fmt.Println("设置redis值错误")
	}

	value, err := client.Get("test_go_redis")

	if err != nil {
		fmt.Println("获取redis值错误")
	}

	fmt.Println("获取redis test_go_redis的值：", string(value))

}
