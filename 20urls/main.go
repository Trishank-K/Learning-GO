package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://jsonplaceholder.typicode.com/posts?id=1&purchase=gsaiwiogvrg&isUser=true"

func main() {
	fmt.Println(myUrl)
	result, err := url.Parse(myUrl)
	checkNilErr(err)
	// fmt.Println("Scheme: ", result.Scheme)
	// fmt.Println("Host: ", result.Host)
	// fmt.Println("Path: ", result.Path)
	// fmt.Println("RawQuery: ", result.RawQuery)
	// fmt.Println("Port: ", result.Port())

	qParams := result.Query()
	fmt.Printf("The Type of query Params are: %T\n", qParams)

	fmt.Println(qParams["id"])
	for key, val := range qParams {
		fmt.Println("Value of", key, "is :", val)
	}
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "jsonplaceholder.typicode.com",
		Path:     "/posts",
		RawQuery: "id=1&purchase=eijiqap",
		RawPath:  "user=Trishank",
	}
	// Whenever Creating a URL, pass as reference
	newURL := partsOfUrl.String()
	fmt.Println("New URL:", newURL)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
