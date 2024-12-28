package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello World"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your Name: ")

	// comma ok || err err

	input, _ := reader.ReadString('\n')
	// input, err := reader.ReadString('\n')
	// The error gets stored in the err variable
	fmt.Print("Hello ", input)
	fmt.Printf("Type of Name: %T", input)

}
