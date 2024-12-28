package main

import "fmt"

// variable := 30000
/*
	This is not allowed since the walrus operator is allowed only inside functions
*/

/*
	var var2 = 30000
	This Works
*/

const LoginToken string = "audfhaefuhadufhea"

// LoginToken has a capital "L" which implies it is a public variable

func main() {
	var username string = "Trishank Kaushik"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.4556466886416878948
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var num int
	fmt.Println(num)
	fmt.Printf("Variable is of type: %T \n", num)

	// implicit type
	var website = "www.google.com"
	fmt.Println(website)
	// website = 3 This is wrong since website is already declared string
	fmt.Printf("Variable is of type: %T \n", website)

	// no var style

	numberOfUser := 30000
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
