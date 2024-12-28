package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println("Initial Value: ", ptr)
	num := 24
	ptr = &num
	fmt.Println("Value of Pointer is: ", *ptr)
	fmt.Println("Address of Pointer is: ", ptr)

}
