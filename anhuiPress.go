package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"sort"
	"strings"
)

/**
业务参数类
*/
type apiParams struct {
	ORG_NO      string //机构编号
	PUBMS_CODE  string //省份编号
	CONS_NO     string //户号
	MON         string //月份
	FLOW_NO_OLD string //支付宝订单流水号
	REMARK      string //备注
}

//压测安徽电网
func main() {
	fmt.Println("Press Start")
	//设置最大cpu执行核数
	runtime.GOMAXPROCS(runtime.NumCPU())
	//设置一个channel
	channel := make(chan int, 20)
	for i := 0; i < 20; i++ {
		go doAction(channel, i)
	}
	for i := 0; i < 20; i++ {
		<-channel
	}
}

func doAction(channel chan int, i int) {
	//网关
	gateway := "https://test-openapi.bangdao-tech.com/gateway.do"
	//业务数据
	var params = apiParams{
		ORG_NO:      "34401",
		PUBMS_CODE:  "PWANHUI",
		CONS_NO:     "5147003923",
		MON:         "000000",
		FLOW_NO_OLD: "784766578115789",
		REMARK:      "",
	}

	apiParams := map[string]string{
		"ORG_NO":      "34401",
		"PUBMS_CODE":  "PWANHUI",
		"CONS_NO":     "5147003923",
		"MON":         "000000",
		"FLOW_NO_OLD": "784766578115789",
		"REMARK":      "",
	}
	//业务参数转json
	paramsJson, _ := json.Marshal(params)
	bizContent := string(paramsJson)

	//公共参数
	sysParams := make(map[string]interface{})
	sysParams["app_id"] = "2018051554153677"
	sysParams["bangdao_sdk"] = "1.0.0"
	sysParams["biz_content"] = bizContent
	sysParams["charset"] = "UTF-8"
	sysParams["format"] = "json"
	sysParams["method"] = "bangdao.alkaid.anhui.power.ticket.print"
	sysParams["sign_type"] = "RSA"
	sysParams["timestamp"] = "1527738428"
	sysParams["version"] = "1.0"
	sysParams["sign"] = aliPaySign(sysParams)
	sysParams["sign"] = "ZZmGLEBhZT7T%2BA9%2FZ9jqaIj%2FlV1AIpcFa%2BWrCCroQhOnqiAWZ5%2FtNF5k111XBRH5LkVsloqbzCyuTYkO52RTKCt8l69IQDZsKwn2FRFnU2VfF7nHn4NbzdeNMPLvPc%2BAEmJxVg%2F9yugB3Uo6FcvthN%2BsrZLIk63DdpNoaaClOPU%3D"

	//fmt.Println(urlencode(bizContent))

	//将参数放入url中
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range sysParams {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for i, k := range sorted_keys {
		//fmt.Printf("k=%v, v=%v\n", k, mReq[k])
		value := fmt.Sprintf("%v", sysParams[k])
		if value != "" {
			if i != (len(sorted_keys) - 1) {
				signStrings = signStrings + k + "=" + value + "&"
			} else {
				signStrings = signStrings + k + "=" + value //最后一个不加此符号
			}
		}
	}

	requestUrl := gateway + "?" + signStrings
	lparam, pErr := url.Parse(requestUrl)
	if pErr != nil {
		fmt.Println(pErr)
	}
	requestUrl = gateway + "?" + lparam.Query().Encode()
	fmt.Println(requestUrl)

	httpRequest(requestUrl, "GET", apiParams)
	//将i存入channel中
	channel <- i
}

/**
发送http请求
*/
func httpRequest(url string, method string, params map[string]string) {

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

//成功签名：M3ZHgxryIBRL+HWEACM5Ebu4d0n4SrSi//GACE6tZk5SATsPjfW3G0nFmka+zosEwkmWhI/8GC7z1KPDIHrpI4CPrqR9+349L7rNvf9ZbwPpqX5cSnXZ14QrZBSst+YvrjHZVc+2jIlkt99Sx6oIPItQKFvqQW5yiiIwJ3yjgZc=

func aliPaySign(mReq map[string]interface{}) string {
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for i, k := range sorted_keys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			if i != (len(sorted_keys) - 1) {
				signStrings = signStrings + k + "=" + value + "&"
			} else {
				signStrings = signStrings + k + "=" + value //最后一个不加此符号
			}
		}
	}
	fmt.Println("生成的待签名---->", signStrings)

	//============================= 开始签名 ==================================
	block, _ := pem.Decode([]byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAI5B6Jyyy0jTBxi7
ZjoRIIqqgdkanmkAQ6/1BJCDsmZHt9aV8wog6oaTgM5hccqSLUFtxKI1iRbzREHJ
5WowXQgbn5NzOUtclh9WB58Q0+A82Xn/qbUj1uj2rf1kfI4BYHRzqU/oqJy3zs5b
OIPDRm828zVlsOyOxoCahYK7ACt7AgMBAAECgYBpXGFHJZkwP7PC2pEVMrgyW6S+
QlvPEvYO7V7LFe9cl4OOOBMAjm6E69x7fmqECQ3hpsX7CcGWMOh9/WWf52AO1SC4
ZdViTkAmivdSrABQVk77zL76xp1JQJGIMULUIIxtHr2CIqDnUa7cE8J9h08+H3Uh
EXk8fTcfzk17Wd7reQJBAMm4EiAP7NHj8DYbQ5XThcL+qdpGiCe4pTIemHKdqXia
7HbOYJOmGz8kG1JiA9C/BXvNTJzVJRxhrMLLgz4E2J0CQQC0ia0OCk8w1q/xt65z
eFO7fp7+JjGXXdhFjNfIy7l8KIXQXAkC0AWyCaEaxFGD4naEXq6tMxDYINFYhM94
ihz3AkAs53H5ezIPwzsNAGK6Lp0fPEuZUIwss3jqwE7VbO118J7lD6oGFDsKgkIm
w7TvuiZBWgFNBFG/kKrBrkYn/zItAkEAq0KDCHOdJa4yn3DoKx91BsKZ3muq1XUv
iOKPfCpN3UqvjbUQ0zs0e+yIiFS9Iwe799u9hg2QKAYRsQ/jo+wYhwJAP+xMW1Dq
jy76DMGLEIB0UDz65jFWc/WEw6p3ZAufgUD+0tjP6+kLhypEdNI2rkqWpAvf1d9X
qj19WJGOlvi8wg==
-----END RSA PRIVATE KEY-----
`))

	/*	block, _ := pem.Decode([]byte(`-----BEGIN RSA PRIVATE KEY-----
		MIICWwIBAAKBgQCsAocILoBYZqhMYDg40AFZTUiFjcSxwrUXOF0rgw8hh98tP+Ox
		4awokuF+FJn8qN/9k9gFz9j7zM694vNv976W60k2ye6uiQdQy/gOJh+ciFME3kAH
		QoyvuItKRec+3cEhwblpuY7Gchk0mk22WXWyAHuNGXcCMGSJo9ugGnPDUwIDAQAB
		AoGARDgcZdpLfMP6K5Bdu+qDHm/QO2emgvm96J+qE/++mIXStZeJLptaNB1M4Tw6
		dkJj06Y3Htb4L6ViuVyxP875/yqKocNur4KeoTLC/t+9L7f6jey7GCvWlCcpp97A
		NiVCPILG+7Py2+xGyv0tQT+98yJqTb0yA0nIsmq1XpmjkQECQQDUmVH8L6mJvyzj
		xDA32jLKoQLWyJTvzkBxeTnsuTL6GqXEcj9HN0Q1qEyTl/DcN5YfACL4XevV1NNV
		s6CrKEojAkEAzx/5d3aAI3dg9VHfAcWdiUZ5uTbx+qpDfZ4CtoyhEfolRGTiy2SC
		y9GDrtpY3NO1mm+D5lIM3HPbTgAxtvm9EQJADmI5K8jFva4TiW1ynbTDjvYJzSJR
		AVCBB6xeAOgezNEUug/IvDa/BKpYU/wJrbyNCZfmxnqE87ise7Xlfu8A5QJBAJTP
		Bx9SLvvMMAfwm0UdonJXBOsR08Zg/35HwPFAlhRhYNcDmIHCo8olq/M7Am8dV7Mt
		/VjDiGP2hRBESXOJd9ECPxeHMhK4prFTN1N/+fHcnbB61P89mIsH5ff38D9uEwa/
		sLqen0K1ibjhgbcs0LjoVklK9fxJ5AK7SwU0Oxaq0w==
		-----END RSA PRIVATE KEY-----
		`))*/

	if block == nil {
		fmt.Println("rsaSign private_key error")
		return ""
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Printf("x509.ParsePKCS1PrivateKey-------privateKey----- error : %v\n", err)
		return ""
	} else {
		//fmt.Println("x509.ParsePKCS1PrivateKey-------privateKey-----", privateKey)
	}
	result, err := RsaSign(signStrings, privateKey)
	fmt.Println("alipay.RsaSign=========", result, err)
	return result
}

/**
 * RSA签名
 * @param $data 待签名数据
 * @param $private_key_path 商户私钥文件路径
 * return 签名结果
 */
func RsaSign(origData string, privateKey *rsa.PrivateKey) (string, error) {

	h := sha1.New()
	h.Write([]byte(origData))
	digest := h.Sum(nil)

	s, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA1, digest)
	if err != nil {
		fmt.Errorf("rsaSign SignPKCS1v15 error")
		return "", err
	}
	data := base64.StdEncoding.EncodeToString(s)
	return string(data), nil
}
