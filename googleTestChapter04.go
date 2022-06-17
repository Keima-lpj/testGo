package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type Get interface {
	Get(url string) string
}

type Post interface {
	Post(url string) string
}

type GetPost interface {
	Get
	Post
}

type MRetriever struct {
	contents string
}

func (r MRetriever) String() string {
	return fmt.Sprintf("我重写了打印这个struct时的值：contents:%s", r.contents)
}

func (r MRetriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}

func (r MRetriever) Post(url string) string {
	return url
}

func main() {
	var r GetPost
	r = MRetriever{"青鸟不传云外信"}
	fmt.Println(r)

	/*fmt.Println(r.Get("https://www.baidu.com"))
	fmt.Println(r.Post("www.bilibili.com"))*/
}
