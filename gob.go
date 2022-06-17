package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func main() {
	/*str := "ekwhlashf;kajhsdf"
	target, err := Serialize(str)
	if err != nil {
		fmt.Println("err1:", err)
		return
	}
	fmt.Println("加密后：", string(target))

	str1 := ""
	err = Deserialize(target, &str1)
	if err != nil {
		fmt.Println("err2:", err)
		return
	}
	fmt.Println("解码后：", str1)*/

	str := "\\x0e\\xff\\x81\\x04\\x01\\x02\\xff\\x82\\x00\\x01\\x10\\x01\\x10\\x00\\x00v\\xff\\x82\\x00\\x01\\x06string\\x0c\\x0f\\x00\\radminuserinfo\\x16*audituserapi.UserInfo\\xff\\x83\\x03\\x01\\x01\\bUserInfo\\x01\\xff\\x84\\x00\\x01\\x04\\x01\\bClientID\\x01\\x0c\\x00\\x01\\x05Roles\\x01\\xff\\x86\\x00\\x01\\x05Scope\\x01\\x0c\\x00\\x01\\x06UserID\\x01\\x0c\\x00\\x00\\x00\\x16\\xff\\x85\\x02\\x01\\x01\\b[]string\\x01\\xff\\x86\\x00\\x01\\x0c\\x00\\x00K\\xff\\x84H\\x01(4b81c5c0c9b944f64beff7777eac2b9bbd0fdcaa\\x01\\x01\\x05admin\\x01\\x05admin\\x01\\x0czhaolunxiang\\x00"
	target := []byte(str)
	var dst interface{}
	err := Deserialize(target, dst)
	if err != nil {
		fmt.Println("err2:", err)
		return
	}
	fmt.Println("解码后：", dst)
}

func Serialize(a string) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(a)
	if err == nil {
		return buf.Bytes(), nil
	}
	return nil, err
}

func Deserialize(d []byte, dst interface{}) error {
	dec := gob.NewDecoder(bytes.NewBuffer(d))
	return dec.Decode(&dst)
}
