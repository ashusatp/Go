package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

//[+] multiple returns
func getlanguage() (string, string, bool) {
	return "Js", "Go", true
}

//[+] Passing Functions as Arguments
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

//if parameters have same type
func addHavingSameparamTypes(a, b int) int {
	return a + b
}

//[+] Returning Functions
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// [+]  Named Return Values
func calculate(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return // This will return the named values `sum` and `product`
}

// [+] Variadic Functions - Variadic functions can take a variable number of arguments. This is done by using an ellipsis (...) before the type.
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	//  multi return
	lang1, lang2, langBool := getlanguage()
	fmt.Println(lang1, lang2, langBool)

	// Passing Functions as Arguments
	result := applyOperation(5, 3, addHavingSameparamTypes)
	fmt.Println(result) // Output: 8

	// Returning Functions
	double := multiplier(2)
	fmt.Println(double(5)) // Output: 10

	// Anonymous Functions
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(2, 3)) // Output: 5

	// Variadic Functions
	nums := []int{4, 5, 6, 7, 8}
	fmt.Println(sum(1, 2, 3)) // Output: 6
	fmt.Println(sum(nums...)) // Output: 30
}
