package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

}

func testGet() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic("baocuole!")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}

func testPost() {
	resp, err := http.Post("http://www.baidu.com", "application/x-www-form-urlencoded", strings.NewReader("id=1"))
	if err != nil {
		panic("baocuole!")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}
