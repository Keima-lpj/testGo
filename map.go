package main

import (
	"fmt"
)

func main() {
	var m map[int]string
	m = make(map[int]string)
	m[1] = "ok"
	m[2] = "yes"
	delete(m, 1)
	fmt.Println(m)

	m2 := make(map[int]map[int]int)
	a, ok := m2[2][1]
	if !ok {
		m2[2] = make(map[int]int)
	}
	m2[2][1] = 10
	a, ok = m2[2][1]
	fmt.Println(a, ok)

	sm := make([]map[int]string, 10)
	for k := range sm {
		sm[k] = make(map[int]string)
		sm[k][1] = "x"
		fmt.Println(sm[k])
	}
	fmt.Println(sm)

	/*	mmm := map[int]string{1:"a", 2:"b", 3:"c", 4:"d", 5:"e"}
		smm := make([]int, len(mmm))
		i := 0
		for k,_ := range mmm{
			smm[i] = k
			i++
		}
		sort.Ints(smm)
		fmt.Println(smm)*/

	/*	m1 := map[int]string{1:"a", 2:"b", 3:"c", 4:"d", 5:"e"}
		m2 := make(map[string]int)

		for k,v := range m1 {
			m2[v] = k
		}
		fmt.Println(m1)
		fmt.Println(m2)*/
}

/**
编写一个函数，使用map[string]map[string]string的map类型
key表示用户名，是唯一的，不可以重复
如果某个用户名存在，就将其密码修改为888888，如果不存在就增加这个用户信息（包括昵称和密码）
*/

func modifyUser(users map[string]map[string]string, name string) {
	v, ok := users[name]
	if ok {
		//有，改密码
		v["pwd"] = "888888"
	} else {
		//没有，新增
		newUser := make(map[string]string)
		newUser["pwd"] = "888888"
		newUser["nickname"] = "111"
		users[name] = newUser
	}

}
