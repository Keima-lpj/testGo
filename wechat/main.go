package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"log"
	"net/http"
	"net/url"
)

const (
	APPID     = "wxfa7b4164fcdca3b3"
	APPSECRET = "b32bd2449b6110906baf97484cd00e62"
	TOKEN     = "akgfakgfiuekejwkjfadf"
	STATE     = "ajdfghaigfeifbadf"
)

func main() {
	http.HandleFunc("/", WechatToken)              //设置访问的路由
	http.HandleFunc("/getQrcode", GetWechatQrCode) //设置访问的路由
	http.HandleFunc("/getWebQrcode", GetWebWechatQrCode)
	http.HandleFunc("/codeToAccesstoken", CodeToAccessToken) //设置访问的路由
	http.HandleFunc("/getUserInfo", GetUserInfo)             //设置访问的路由
	err := http.ListenAndServe(":8990", nil)                 //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//配置token
func WechatToken(resp http.ResponseWriter, req *http.Request) {
	//获取请求参数获取参数
	val := req.URL.Query()
	signature := val["signature"]
	timestamp := val["timestamp"]
	nonce := val["nonce"]
	echostr := val["echostr"]
	fmt.Println(signature, timestamp, nonce, echostr)
	signStr := nonce[0] + timestamp[0] + TOKEN
	sign := Sha1(signStr)
	if sign == signature[0] {
		resp.Write([]byte(echostr[0]))
	} else {
		resp.Write([]byte("请求参数错误"))
	}
}

//生成微信快捷登录的链接
func GetWechatQrCode(resp http.ResponseWriter, req *http.Request) {
	wechatUrl := "https://open.weixin.qq.com/connect/oauth2/authorize"
	redirect_uri := url.QueryEscape("http://9w3x3i.natappfree.cc/codeToAccesstoken")
	url := fmt.Sprintf("%s?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect",
		wechatUrl, APPID, redirect_uri, "snsapi_login", STATE,
	)
	resp.Write([]byte(url))
}

func GetWebWechatQrCode(resp http.ResponseWriter, req *http.Request) {
	wechatUrl := "https://open.weixin.qq.com/connect/qrconnect"
	redirect_uri := url.QueryEscape("http://9w3x3i.natappfree.cc/codeToAccesstoken")
	url := fmt.Sprintf("%s?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect",
		wechatUrl, APPID, redirect_uri, "snsapi_login", STATE,
	)
	resp.Write([]byte(url))
}

//用Code换区AccessToken
func CodeToAccessToken(resp http.ResponseWriter, req *http.Request) {
	wechatUrl := "https://api.weixin.qq.com/sns/oauth2/access_token"
	val := req.URL.Query()
	code := val["code"][0]
	url := fmt.Sprintf("%s?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		wechatUrl, APPID, APPSECRET, code,
	)
	res := make(map[string]interface{})
	err := httplib.Get(url).ToJSON(&res)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	resp.Write([]byte(res["access_token"].(string) + "--" + res["openid"].(string)))
}

func GetUserInfo(resp http.ResponseWriter, req *http.Request) {
	wechatUrl := "https://api.weixin.qq.com/sns/userinfo"
	val := req.URL.Query()
	openId := val["open_id"][0]
	accessToken := val["access_token"][0]
	u := fmt.Sprintf("%s?access_token=%s&openid=%s&lang=zh_CN",
		wechatUrl, accessToken, openId,
	)
	fmt.Println(u)
	res := make(map[string]interface{})
	err := httplib.Get(u).ToJSON(&res)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	for k, v := range res {
		resp.Write([]byte(fmt.Sprintf("%s:%s", k, v)))
	}
}

func Sha1(data string) string {
	sha := sha1.New()
	sha.Write([]byte(data))
	return hex.EncodeToString(sha.Sum([]byte("")))
}
