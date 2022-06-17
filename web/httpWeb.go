package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//定义一个函数类型
type shellFunc func(w http.ResponseWriter, r *http.Request) error

//接收这个函数类型，并且返回一个可以被HandleFunc接受的func
func handleShellFunc(shell shellFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := shell(w, r)
		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(code), code)
		}
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("hello world rua!"))
	go func() {
		fmt.Println(222)
		time.Sleep(time.Second * 10)
		fmt.Println(111)
	}()

	go func() {
		fmt.Println(333)
		time.Sleep(time.Second * 10)
		fmt.Println(444)
	}()
	return nil
}

var i = 1

func Return308TenTime(w http.ResponseWriter, r *http.Request) {
	if i == 10 {
		fmt.Println("达到了10次，返回200")
		i = 1
		w.WriteHeader(200)
	}
	if r.Host == "localhost:9090" {
		fmt.Println("没达到10次，返回308", i)
		i++
		w.Header().Set("Location", "http://localhost:9090/return308")
		w.WriteHeader(308)
	}
}

func Return308(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Return308 收到一个308，请跳转到200去")
	w.Header().Set("Location", "http://localhost:9090/Return200")
	w.WriteHeader(308)
}

func Return200(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Return200 请求成功")
	fmt.Println("header:", r.Header)
	fmt.Println("err：", r.ParseForm())
	fmt.Println("Form:", r.Form)
	fmt.Println("PostForm:", r.PostForm)
	/*a, _ := ioutil.ReadAll(r.Body)
	fmt.Println("body:", string(a))*/
	w.WriteHeader(200)
}

func main() {

	//http.HandleFunc("/", handleShellFunc(sayhelloName)) //设置访问的路由
	//http.HandleFunc("/Return308TenTime", Return308TenTime) //设置访问的路由
	http.HandleFunc("/Return308", Return308)
	http.HandleFunc("/return200", Return200)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
