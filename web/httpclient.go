package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	/*params := make(map[string]string)
	params["client_id"] = "sfuykjahfkasjdhfkadsf"
	params["client_secret"] = "dsgfuiwegfjahbfjadhsfg"
	j, _ := json.Marshal(params)*/
	/*client := http.Client{}
	req, _ := http.NewRequest("POST", "http://localhost/return200", bytes.NewBuffer(j))
	req.Header.Set("Access-Token", "02d0d073be774b009180a770582905bb")
	req.Header.Set("Authorization", "Bearer 02d0d073be774b009180a770582905bb")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")*/
	resp, err := http.Post("http://localhost/return200",
		"application/x-www-form-urlencoded",
		strings.NewReader("client_id=sfuykjahfkasjdhfkadsf"))
	//resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v", resp)
	}
}
