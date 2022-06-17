package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	poll *redis.Pool
)

func init() {
	poll = &redis.Pool{
		MaxActive:   0,
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (conn redis.Conn, e error) {
			conn, err := redis.Dial("tcp", "192.168.1.135:6379")
			if err != nil {
				return nil, err
			}
			//验证密码
			if _, err = conn.Do("AUTH", "redis123456"); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, nil
		},
	}
}

func main() {
	conn1 := poll.Get()
	defer conn1.Close()

	_, err := conn1.Do("SET", "name2", "L!")
	errorHandle(err)
	if err != nil {
		fmt.Println(err)
		return
	}

	name2, err := redis.String(conn1.Do("GET", "name2"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name2)

}
