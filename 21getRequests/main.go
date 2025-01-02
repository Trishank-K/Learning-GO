package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"
	response, err := http.Get(myUrl)
	checkNilErr(err)
	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)

	var responseString strings.Builder
	content, err := io.ReadAll(response.Body)
	checkNilErr(err)
	byteCount, err := responseString.Write(content) // String Builder Method
	checkNilErr(err)
	fmt.Println("Byte Count:", byteCount)
	fmt.Println(responseString.String())
	fmt.Println(string(content)) //String Method
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
