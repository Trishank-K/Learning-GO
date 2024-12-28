package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	// 01-02-2006 is the standard used for formatting and vice versa

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}

/*
	If I wish to create a executable file for the current program:

	- For Windows
		$env:GOOS="windows"; $env:GOARCH="amd64"; go build
	- For mac
		$env:GOOS="darwin"; $env:GOARCH="amd64"; go build
	- For linux
		$env:GOOS="linux"; $env:GOARCH="amd64"; go build
*/
