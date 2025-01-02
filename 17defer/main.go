package main

import "fmt"

func main() {
	defer fmt.Println("World")
	// When we use the defer keyword, the line just next to it gets executed at the end of program
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
