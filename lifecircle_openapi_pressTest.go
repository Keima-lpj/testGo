package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"runtime"
	"sort"
	"strings"
	"time"
)

var (
	//gateway = "https://youdian1.51youdian.com/Wap/StorePayment/orderSave"
	gateway = "https://shq-api-test.51fubei.com/gateway"

	authCode = []string{
		"282794526385579403",
		"",
		"",
		"",
		"",
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

/**
业务参数对象
*/

func main() {
	fmt.Println("Press Start") //压测开始

	//设置CPU最大核数
	runtime.GOMAXPROCS(runtime.NumCPU())

	//设置一个长度为20的channel
	channel := make(chan int, 20)

	for i := 0; i < 1; i++ {
		go DoPressSwipe(channel, i, authCode[i])
	}
	for i := 0; i < 1; i++ {
		<-channel
	}

}

//生活圈开放平台压测。刷卡
func DoPressSwipe(c chan int, i int, authcode string) {
	//公共参数
	pp := make(map[string]interface{})
	pp["app_id"] = "20180712160225576641"
	pp["method"] = "openapi.payment.order.swipe"
	pp["format"] = "json"
	pp["sign_type"] = "md5"
	pp["nonce"] = "1234567"
	pp["version"] = "1.0"

	key := "1664aba93c82c9e459f4ae42699d6ffc"
	//业务参数
	ap := make(map[string]interface{})
	ap["type"] = 2
	ap["merchant_order_sn"] = r.Intn(1000000000)
	ap["auth_code"] = authcode
	ap["total_fee"] = 10
	ap["store_id"] = 7643
	ap["call_back_url"] = "http://www.alikoubei.com"
	ap["device_no"] = "18602767501"
	//业务参数转json
	paramsJson, _ := json.Marshal(ap)
	bizContent := string(paramsJson)
	pp["biz_content"] = bizContent
	//生成签名
	pp["sign"] = SignSort(pp, key)

	HttpRequest(gateway, "POST", pp)
}

//生成签名并加密
func SignSort(pp map[string]interface{}, key string) string {
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range pp {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for i, k := range sorted_keys {
		//fmt.Printf("k=%v, v=%v\n", k, mReq[k])
		value := fmt.Sprintf("%v", pp[k])
		if value != "" {
			if i != (len(sorted_keys) - 1) {
				signStrings = signStrings + k + "=" + value + "&"
			} else {
				signStrings = signStrings + k + "=" + value //最后一个不加此符号
			}
		}
	}
	signStrings += key
	//MD5加密
	h := md5.New()
	h.Write([]byte(signStrings)) // 需要加密的字符串为
	cipherStr := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(cipherStr))
}

//请求网关方法
func HttpRequest(url string, method string, params map[string]interface{}) {
	paramsJson, _ := json.Marshal(params)
	jsonStringContent := string(paramsJson)

	pbyte := bytes.NewBuffer([]byte(jsonStringContent))

	res, err := http.Post(url, "application/json;charset=utf-8", pbyte)
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(result))
}
