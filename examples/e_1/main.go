package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	var reqUrl string

	var reqMethod string

	flag.StringVar(&reqUrl, "u", "", "request url")

	flag.StringVar(&reqMethod, "m", "GET", "request method")

	flag.Parse()

	if reqUrl == "" {
		fmt.Println(errors.New("please input a correct url"))
		return
	}

	client := http.Client{}

	request, err := http.NewRequest(reqMethod, reqUrl, nil)

	if err != nil {
		panic(err)
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.87 Safari/537.36")

	res, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	result, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))

}
