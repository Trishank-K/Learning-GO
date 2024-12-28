package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	num := 1

	for num < 10 {
		if num == 2 {
			goto test
		}
		if num == 5 {
			break
		}
		fmt.Println("Value is: ", num)
		num++
	}

test:
	fmt.Println("Loop ends Here")

}
