package main

import (
	"encoding/json"
	"fmt"
)

//定义一个结构体
type Monster struct {
	//这里如果希望转json之后使用别名，则直接在对应的key后面添加tag标签
	Name     string `json:"monster_name"`
	Age      int    `json:"monster_age"`
	Birthday string
}

func main() {
	/**
	struct转json
	*/
	monster := Monster{
		Name:     "牛牛",
		Age:      100,
		Birthday: "aaa",
	}
	//这里不论传指针还是实例都可以
	structToJson(&monster)

	//map转json
	map1 := make(map[string]string)
	map1["name"] = "111"
	map1["hobby"] = "basketball"

	mapToJson(map1)

	//slice转json
	slice1 := []map[string]string{}
	slice1 = append(slice1, map1)
	mapToJson(slice1)

	//json转struct
	jsonStr := "{\"monster_name\":\"牛牛\",\"monster_age\":100,\"Birthday\":\"aaa\"}"
	monster2 := Monster{}
	jsonToStruct(jsonStr, &monster2)

	//json转struct
	jsonStr2 := "{\"monster\":\"牛牛\",\"monster_age\":100,\"Birthday\":\"aaa\"}"
	map2 := make(map[string]interface{})

	jsonToStruct(jsonStr2, &map2)
}

/**
结构体转json
*/
func structToJson(monster *Monster) {
	data, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("转换失败：", err)
	}
	fmt.Println(string(data))
}

/**
map或者slice转json
*/
func mapToJson(m interface{}) {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("转换失败：", err)
	}
	fmt.Println(string(data))
}

/**
json转strcut或者map或者slice
*/
func jsonToStruct(jsonStr string, a interface{}) {
	err := json.Unmarshal([]byte(jsonStr), a)
	if err != nil {
		fmt.Println("转换失败：", err)
	}
	fmt.Println(a)
}
