package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

type Response struct {
	Code    int
	Message string
	Success bool
	Data    map[string]string
}

func main() {
	yaTest()
}

//waitgroup
var wait = sync.WaitGroup{}

func yaTest() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wait.Add(1000)

	var data = make(map[int]map[string]string)
	var d = make(map[string]string)

	d["orderSn"] = "20180425104312907640"
	d["openid"] = "oo_bBwm8ZvgrW_Tc-oWfT_kzcMKc"
	data[0] = d

	d["orderSn"] = "20180424200702423680"
	d["openid"] = "oDBdo1nAHyucc482qho4-KbPb4nQ"
	data[1] = d

	d["orderSn"] = "2018042419571388733s"
	d["openid"] = "oo_bBwo3W8dT8jkgZbFhqTuyKO1s"
	data[2] = d

	d["orderSn"] = "20180424180128826836"
	d["openid"] = "oo_bBwm8ZvgrW_Tc-oWfT_kzcMKc"
	data[3] = d

	d["orderSn"] = "20180424162715764289"
	d["openid"] = "oo_bBwjvdH9Db0uVdwS331YWz9eM"
	data[4] = d

	d["orderSn"] = "20180424162229264686"
	d["openid"] = "oo_bBwv4dsw8OJZewQLE1Z2q7j-8"
	data[5] = d

	d["orderSn"] = "20180424160838462185"
	d["openid"] = "oo_bBwo3W8dT8jkgZbFhqTuyKO1s"
	data[6] = d

	d["orderSn"] = "20180424160220811535"
	d["openid"] = "oDBdo1nAHyucc482qho4-KbPb4nQ"
	data[7] = d

	d["orderSn"] = "20180424102801841304"
	d["openid"] = "oo_bBwn987Q4OkRuI1RqAvd2s3BE"
	data[8] = d

	d["orderSn"] = "20180424101003435693"
	d["openid"] = "oo_bBwi9mR_8VPUPs58QeiKmOEdY"
	data[9] = d

	var n int
	for i := 0; i < 1000; i++ {

		if i > 100 {
			n = i % 100
		} else {
			n = i % 100

		}
		go doAction(data[n])
	}
	wait.Wait()
}

func doAction(test map[string]string) {
	var req_url = "http://lifecircle-openapi-test.51youdian.com:9080/getAudience"

	client := &http.Client{}
	req, err := http.NewRequest("POST", req_url, strings.NewReader("orderSn="+test["orderSn"]+"&openid="+test["openid"]))

	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)

	res := Response{}
	json.Unmarshal([]byte(string(body)), &res)
	fmt.Println(res)
	resp.Body.Close()
	wait.Done()
}
