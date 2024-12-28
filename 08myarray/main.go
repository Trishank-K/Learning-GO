package main

import "fmt"

func main() {

	var alphabets [4]string
	alphabets[0] = "A"
	alphabets[1] = "B"
	alphabets[3] = "D"

	fmt.Println("Alphabets: ", alphabets)
	fmt.Println("Length of Alphabets is: ", len(alphabets))

	var nums = [3]string{"1", "2", "3"}
	fmt.Println("Numbers: ", nums)
	fmt.Println("Length of Nums is: ", len(nums))

}
