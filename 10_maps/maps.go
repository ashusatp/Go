package main

import (
	"fmt"
	"maps"
)

func main() {

	//Declaring a Map -> You can declare maps in different ways:

	//-> a. Using a Literal
	//-> b. Using make()

	// ages := make(map[string]int)

	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}

	// Add a new key-value pair
	ages["Charlie"] = 35

	// Modify an existing key-value pair
	ages["Alice"] = 26

	// Delete a key-value pair
	delete(ages, "Bob")

	// Check if a key exists
	age, exists := ages["Alice"]
	if exists {
		fmt.Println("Alice's age is", age) // Output: Alice's age is 26
	}

	// Iterate over the map
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	m1 := map[string]int{"price": 20, "phone": 1}
	m2 := map[string]int{"price": 20, "phone": 1}

	if maps.Equal(m1, m2) {
		fmt.Println("m1 and m2, both are equal")
	}
}
