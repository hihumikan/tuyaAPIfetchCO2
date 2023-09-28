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
	// Step 1: アクセストークンを取得
	accessToken, err := getAccessToken()
	if err != nil {
		fmt.Println("アクセストークンを取得できませんでした:", err)
		return
	}
	fmt.Println("アクセストークン:", accessToken)

	// Step 2: デバイス情報を取得
	deviceInfo, err := getDeviceInfo(accessToken)
	if err != nil {
		fmt.Println("デバイス情報を取得できませんでした:", err)
		return
	}
	fmt.Println("デバイス情報:", deviceInfo)
}

func getAccessToken() (string, error) {
	url := "https://openapi.tuyaus.com/v1.0/token?grant_type=1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "", err
	}
	req.Header.Add("client_id", "")
	req.Header.Add("sign", "")
	req.Header.Add("t", "")
	req.Header.Add("sign_method", "HMAC-SHA256")
	req.Header.Add("nonce", "")
	req.Header.Add("stringToSign", "")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var response AccessTokenResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	return response.Result.AccessToken, nil
}

func getDeviceInfo(accessToken string) (string, error) {
	url := "https://openapi.tuyaus.com/v1.0/devices/eb4d67616c470a928fj1e4"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "", err
	}
	req.Header.Add("client_id", "")
	req.Header.Add("access_token", accessToken)
	req.Header.Add("sign", "")
	req.Header.Add("t", "")
	req.Header.Add("sign_method", "HMAC-SHA256")
	req.Header.Add("nonce", "")
	req.Header.Add("stringToSign", "")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
