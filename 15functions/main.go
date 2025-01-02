package main

import "fmt"

func main() {
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	proRes, msg := proAdder(2, 5, 8, 7, 5, 4, 3, 32, 1)
	fmt.Println("Pro Result is:", proRes)
	fmt.Println("Msg:", msg)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Yo mama"
}

func greeter() {
	fmt.Println("Hello World")
}
