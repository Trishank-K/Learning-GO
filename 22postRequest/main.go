package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	PerformPostJSONRequest()
}

func PerformPostJSONRequest() {
	const myUrl = "http://localhost:8000/post"
	requestBody := strings.NewReader(`
		{
			"courseName":"GO",
			"price":0
		}
	`)
	response, err := http.Post(myUrl, "application/json", requestBody)
	checkNilErr(err)
	var res strings.Builder
	content, err := io.ReadAll(response.Body)
	res.Write(content)
	checkNilErr(err)
	fmt.Println(res.String())
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
