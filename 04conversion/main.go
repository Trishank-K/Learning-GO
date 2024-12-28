package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Program")
	fmt.Println("Provide a number between 1 and 5")
	reader := bufio.NewReader((os.Stdin))
	input, _ := reader.ReadString('\n')
	fmt.Println("The Number You provided is: ", input)
	/*
		input = input + 1
		This is wrong since input is string
		We'll use strconv package
	*/
	// number, err := strconv.ParseFloat(input, 64)
	/*
		Causes Error: strconv.ParseFloat: parsing "4\r\n": invalid syntax
		-  Since while taking input, we pressed enter which caused the input to be appended with the new line character.
	*/
	number, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Final Value: ", number+1)
	}
}
