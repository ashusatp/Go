package main

import "fmt"

// for -> only for looping there is no while loop

func main() {

	// while loop
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// infinite loop
	// for{
	// 	println("1")
	// }

	// classic for loop
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	// we can use break, continue keywords

	// Range loop
	//  0 to n-1
	for i := range 11 {
		fmt.Println(i)
	}
}
