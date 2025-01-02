package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://google.com"

func main() {
	response, err := http.Get(url)
	checkNilErr(err)
	fmt.Printf("Type of Response : %T\n", response)
	defer response.Body.Close() //Caller's responsibilty to close the connection

	data, err := io.ReadAll(response.Body)
	checkNilErr(err)
	fmt.Println(string(data))
}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
