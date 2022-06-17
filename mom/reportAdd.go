package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/reportAdd", reportAdd)
	http.HandleFunc("/amountAdd", amountAdd)
	http.ListenAndServe(":8080", nil)
}

//报表
func reportAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var sum float64
	//从post中读取数据
	r.ParseForm()
	data := r.PostForm["data"]
	stringSlice := strings.Split(data[0], "\n")
	for _, v := range stringSlice {
		if v == "" {
			continue
		} else {
			vA := strings.Split(v, ";")
			if len(vA) == 2 {
				x := vA[1]
				num := x[:len(x)-1]
				tnum, _ := strconv.ParseFloat(num, 64)
				sum += tnum
			}
		}
	}
	m := make(map[string]string)
	m["error"] = "0"
	m["data"] = strconv.FormatFloat(sum, 'E', -1, 64)
	m["data"] = fmt.Sprintf("%.3f", sum)
	v, err := json.Marshal(m)
	if err != nil {
		fmt.Println("转换失败：", err)
	}
	w.Write(v)
}

//金额
func amountAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var sum float64
	//从post中读取数据
	r.ParseForm()
	data := r.PostForm["data"]
	stringSlice := strings.Split(data[0], "\n")
	for _, v := range stringSlice {
		if v == "" {
			continue
		} else {
			vA := v
			index := strings.Index(v, "元")
			if index != -1 {
				vA = v[:index]
			}
			tnum, _ := strconv.ParseFloat(vA, 64)
			sum += tnum
		}
	}
	fmt.Println(sum)
	m := make(map[string]string)
	m["error"] = "0"
	m["data"] = strconv.FormatFloat(sum, 'E', -1, 64)
	m["data"] = fmt.Sprintf("%.2f", sum)
	fmt.Println(m)
	v, err := json.Marshal(m)
	if err != nil {
		fmt.Println("转换失败：", err)
	}
	w.Write(v)
}
