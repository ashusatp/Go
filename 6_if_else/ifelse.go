package main

import "fmt"

func main() {

	age := 18

	// if age >= 18 {
	// 	fmt.Println("adult")
	// } else {
	// 	fmt.Println("not an adult")
	// }

	if age >= 18 {
		fmt.Println("adult")
	} else if age >= 12 {
		fmt.Println("teenager")
	} else {
		fmt.Println("kid")
	}

	// [+] logical operators in if else
	var role = "admin"
	var hasPermission = true

	if role == "admin" || hasPermission {
		fmt.Println("yes")
	}

	// [+] declare variablse inside condition
	if age := 15; age >= 18 {
		fmt.Println("person is an adult", age)
	} else if age >= 12 {
		fmt.Println("person is teenager", age)
	}

	// go doesnot have ternary operator, you will have to use if else
}
