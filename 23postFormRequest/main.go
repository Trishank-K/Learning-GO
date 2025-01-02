package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	PerformFormRequest()
}
func PerformFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstName", "Trishank")
	data.Add("lastName", "Kaushik")
	data.Add("Studying", "GO")

	response, err := http.PostForm(myUrl, data)
	checkNilErr(err)
	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)
	checkNilErr(err)
	fmt.Println(string(content))
}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
