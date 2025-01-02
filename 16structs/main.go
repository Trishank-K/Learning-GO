package main

import "fmt"

func main() {
	Trishank := User{"Trishank", "mail@gmail.com", true, 21}
	fmt.Println(Trishank)
	fmt.Printf("Details are: %+v\n", Trishank)
	fmt.Printf("Name is: %v and email is: %v \n", Trishank.Name, Trishank.Email)
	Trishank.GetStatus()
	Trishank.NewMail()
	fmt.Printf("Modified email is: %v \n", Trishank.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user Active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@gmail.com"
	fmt.Println("Email of this user is: ", u.Email)
}
