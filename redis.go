package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	//跟redis建立连接
	conn, err := redis.Dial("tcp", "192.168.1.135:6379")
	errorHandle(err)
	defer conn.Close()
	//redis验证密码
	authPassword(conn)
	//操作键值对
	operaKeyValue(conn)
	//操作hash
	operaHash(conn)
	//操作链表
	operaList(conn)
}

func errorHandle(err error) {
	if err != nil {
		panic(fmt.Sprintf("catch a err :%v", err))
	}
}

/**
验证redis密码
*/
func authPassword(conn redis.Conn) {
	_, err := conn.Do("auth", "redis123456")
	errorHandle(err)
}

/**
操作键值对
*/
func operaKeyValue(conn redis.Conn) {
	//设置键值对
	_, err := conn.Do("Set", "name", "Keima!")
	errorHandle(err)

	//通过键获取值
	value, err := redis.String(conn.Do("Get", "name"))
	errorHandle(err)
	fmt.Println(value)

	//给键设置过期时间
	_, err = conn.Do("expire", "name", 10)
}

//操作hash
func operaHash(conn redis.Conn) {
	_, err := conn.Do("HSet", "user01", "name", "Joe")
	errorHandle(err)

	_, err = conn.Do("HMSet", "user01", "sex", "男", "age", "18")
	errorHandle(err)

	name, err := redis.String(conn.Do("HGet", "user01", "name"))
	errorHandle(err)
	fmt.Println(name)

	age, err := redis.String(conn.Do("HGet", "user01", "age"))
	errorHandle(err)
	fmt.Println(age)

	nas, err := conn.Do("HMGet", "user01", "name", "age", "sex")
	errorHandle(err)
	if nasSlice, ok := nas.([]interface{}); ok {
		for _, v := range nasSlice {
			fmt.Printf("%s\n", v)
		}
	}
}

func operaList(conn redis.Conn) {
	_, err := conn.Do("lpush", "hero_list", "Iron Man", "American Captain")
	errorHandle(err)

	hero, err := redis.String(conn.Do("rpop", "hero_list"))
	errorHandle(err)
	fmt.Println(hero)

}
