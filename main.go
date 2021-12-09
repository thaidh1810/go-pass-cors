package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"syscall/js"
	"time"
)

const jsCallHTTPError = `
wasmCallHTTP accepts 2 variable as the input
The first one is a object like this: {
  Method: "Method type"
  Url: "https://example.com"
  Headers: {
    "Header 1": "Value"
    "Header 2": "Value"
  }
}
The second is the body of the request
`

type jsHTTPInput struct {
	Method, Url string
	Headers map[string]string
}
/*
func addFunction(this js.Value, p []js.Value) interface{} {
	sum := p[0].Int() + p[1].Int()
	return js.ValueOf(sum)
}*/

var quit = make(chan struct{})

func main() {
	fmt.Println("Go wasm")

	//js.Global().Set("addWASM", js.FuncOf(addFunction))
	js.Global().Set("wasmCallHTTP", js.FuncOf(jsCallHTTP))
	time.Sleep(time.Second * 5)
	<-quit
}

func jsCallHTTP(this js.Value, p []js.Value) interface{} {
	if len(p) == 0 {
		fmt.Println(jsCallHTTPError)
		return js.ValueOf(jsCallHTTPError)
	}
	input := p[0].String()
	fmt.Println(input)
	var inputStruct jsHTTPInput
	err := json.Unmarshal([]byte(input), &inputStruct)
	if err != nil {
		return js.ValueOf(err.Error())
	}
	var reqBody io.Reader
	if len(p) >= 2 {
		reqBody = strings.NewReader(p[1].String())
	}
	res, err := httpCall(inputStruct.Method, inputStruct.Url, inputStruct.Headers, reqBody)
	if err != nil {
		return js.ValueOf(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return js.ValueOf(err)
	}
	return js.ValueOf(data)
}

func httpCall(method, url string, headers map[string]string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if len(headers) > 0 {
		for k,v := range headers {
			req.Header.Set(k, v)
		}
	}
	res, err := client.Do(req)
	fmt.Println()
	return res, err
}