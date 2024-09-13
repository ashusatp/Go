package main

import "fmt"

func main() {

	// -> for slice
	nums := []int{1, 2, 3, 4, 5}
	for index, num := range nums {
		fmt.Println(index, num)
	}

	// -> for map
	m := map[string]int{"phone": 2, "price": 50}
	for key, value := range m {
		fmt.Println(key, value)
	}

	// -> for string
	// unicode point rune
	//  here i is for starting byte of rune
	// 255 1 byte
	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}

}
