package main

import "fmt"

func main() {
	// There is no inheritance in golang. No super, parent or child
	Trishank := User{"Trishank", "mail@gmail.com", true, 21}
	fmt.Println(Trishank)
	fmt.Printf("Details are: %+v\n", Trishank)
	fmt.Printf("Name is: %v and email is: %v \n", Trishank.Name, Trishank.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
