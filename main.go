package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://openapi.tuyaus.com/v1.0/token?grant_type=1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("client_id", "")
	req.Header.Add("sign", "")
	req.Header.Add("t", "")
	req.Header.Add("sign_method", "HMAC-SHA256")
	req.Header.Add("nonce", "")
	req.Header.Add("stringToSign", "")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
