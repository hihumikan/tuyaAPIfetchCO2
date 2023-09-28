package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AccessTokenResponse struct {
	Result struct {
		AccessToken  string `json:"access_token"`
		ExpireTime   int    `json:"expire_time"`
		RefreshToken string `json:"refresh_token"`
		UID          string `json:"uid"`
	} `json:"result"`
	Success bool   `json:"success"`
	T       int    `json:"t"`
	TID     string `json:"tid"`
}

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

	var response AccessTokenResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	accessToken := response.Result.AccessToken
	fmt.Println("アクセストークン:", accessToken)
}
