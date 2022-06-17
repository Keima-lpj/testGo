package main

import (
	"fmt"
	"github.com/axgle/mahonia"
)

func main() {
	bytes := []byte{3, 101, 68, 68, 84, 180, 122, 8, 40, 66, 134, 185, 255, 108}
	fmt.Println(bytes)
	str111 := mahonia.NewDecoder("utf8").ConvertString(string(bytes))
	fmt.Println(str111)

	str := "我爱中国"
	bytes1 := []byte(str)
	fmt.Println(bytes1)
	fmt.Println(string(bytes1))
}
