package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	gogogo()
}

func gogogo() {
	syncNum := 200
	u := "http://0.0.0.0:9090"
	method := "PUT"
	route := "/api/v0/app/meta"
	data := make(map[string]interface{})
	data["bin_path"] = "/usr/bin/python3.7"
	data["key"] = "/home/uos/Desktop/language_info.json"
	b, _ := json.Marshal(data)

	client := http.Client{
		Transport: &http.Transport{
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   100,
			IdleConnTimeout:       60 * time.Second,
			ResponseHeaderTimeout: 60 * time.Second,
		},
	}

	var wg sync.WaitGroup
	wg.Add(syncNum)

	run := func() {
		defer wg.Done()
		req, err := http.NewRequest(method, u+route, bytes.NewReader(b))
		if err != nil {
			fmt.Println("req报错了！", err)
			return
		}
		req.Header.Add("cookie", "token=40c207f994c52234a790828659e5260b")
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("res报错了！", err)
			return
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println("结果：", res)
		fmt.Println("body:", string(body))
	}

	for i := 1; i <= syncNum; i++ {
		go run()
	}
	wg.Wait()
}
