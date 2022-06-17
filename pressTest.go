package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"runtime"
	"runtime"
	"strings"
)

//设置一个缓存channel
var cx chan int

//主函数
func main() {
	fmt.Println("压测开始")
	//设置最大cpu执行核数
	runtime.GOMAXPROCS(runtime.NumCPU())
	cx = make(chan int, 10)
	//设置10个并发数
	for i := 0; i < 10; i++ {
		go Go(i)
	}
}

//压测函数
func Go(i int) {
	//设置参数
	url := "www.baidu.com"

	data := [][]string{
		[]string{"name", "LPJ"},
		[]string{"age", "26"},
		[]string{"sex", "男"},
	}
	//发送http请求
	HttpDo(url, "POST", data)
	//存入channel中
	cx <- i
}

//发送http请求
func HttpDo(url string, method string, params map[string]string) {

	client := &http.Client{}

	req, err := http.NewRequest(method, url, strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
		fmt.Println(err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//设置cookie
	//req.Header.Set("Cookie", "name=anny")
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(body))
}

func httpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}

//使用这个方法的话，第二个参数要设置成”application/x-www-form-urlencoded”，否则post参数无法传递。
func httpPost(url string) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}
