package main

import (
	"fmt"
	"github.com/guaidashu/go_helper"
)

func main() {
	type responseData struct {
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
		Code int         `json:"code"`
	}

	var (
		err  error
		data responseData
	)

	url := "http://127.0.0.1:8088/task/start"

	request := go_helper.Post(url)

	// 语言
	lang := "zh"

	// 时区
	timezone := "UTC"

	// 代理
	proxy := "127.0.0.1:8011"

	// 设备id
	android := "de121sda"

	// other
	other := "other"

	body := "lang=%v&timezone=%v&proxy=%v&android=%v&other=%v"

	body = fmt.Sprintf(body, lang, timezone, proxy, android, other)

	fmt.Println(body)

	request.Body(body)

	request.Header("Content-Type", "application/x-www-form-urlencoded")

	if err = request.ToJSON(&data); err != nil {
		panic(err)
	}

	fmt.Println(data)

}
