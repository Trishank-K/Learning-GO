package main

import (
	"fmt"
	"sort"
)

func main() {
	var alphabets = []string{"A", "B", "C"}
	fmt.Printf("Type of alphabets is: %T \n", alphabets)
	fmt.Println("Alphabets is: ", alphabets)
	alphabets = append(alphabets, "D", "E")
	fmt.Println("Alphabets is: ", alphabets)
	alphabets = alphabets[1:3]
	// colon syntax(To slice a slice):  array[start index: last index(non inclusive)]
	fmt.Println(alphabets)

	nums := make([]int, 4)

	nums[0] = 234
	nums[1] = 945
	nums[2] = 465
	nums[3] = 867
	// nums[4] = 777 This is wrong
	nums = append(nums, 123, 345, 112)
	// This works ^
	fmt.Println(nums)
	sort.Ints(nums)
	fmt.Println(nums)

	// How to Remove a value from slices based on index

	var courses = []string{"react", "js", "rust", "go"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
