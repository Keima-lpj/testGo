package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var test = []struct {
	h       shellFunc
	code    int
	message string
}{
	{testSysError, 500, "Internal Server Error"},
}

func testSysError(w http.ResponseWriter, r *http.Request) error {
	return nil
}

//这个函数来测试handleShellFunc
func TestHandleShellFunc(t *testing.T) {
	for _, v := range test {
		f := handleShellFunc(v.h)
		request := httptest.NewRequest("GET", "https://www.baidu.com", nil)
		response := httptest.NewRecorder()
		f(response, request)

		b, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(b), "\n")
		if response.Code != v.code {
			t.Errorf("code出错了，expect【%s】, got【%s】", v.message, body)
		}
		if body != v.message {
			t.Errorf("body出错了，expect【%s】, got【%s】", v.message, body)
		}

	}
}

func TestHandleShellFuncInServer(t *testing.T) {
	for _, v := range test {
		f := handleShellFunc(v.h)

		//启动一个真实的测试服务器
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)

		b, _ := ioutil.ReadAll(resp.Body)
		body := strings.Trim(string(b), "\n")
		if resp.StatusCode != v.code {
			t.Errorf("code出错了，expect【%s】, got【%s】", v.message, body)
		}
		if body != v.message {
			t.Errorf("body出错了，expect【%s】, got【%s】", v.message, body)
		}

	}
}
