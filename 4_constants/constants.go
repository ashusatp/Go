package main

import "fmt"

const age = 30

// but cant use sorthand outside the function
//  age := 30

func main() {
	const name string = "golang"
	// now we can't reassign the name variable

	// const age int = 12

	// [+] multible constant grouping

	const (
		port = 5000
		host = "Localhost"
	)

	fmt.Println(age)
	fmt.Println(port)
}
