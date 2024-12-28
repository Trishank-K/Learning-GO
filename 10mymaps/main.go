package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["Js"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS is short for : ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	for key, value := range languages {
		fmt.Printf("For key %v, values is %v\n", key, value)
	}

}
