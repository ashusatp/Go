package mian

import "fmt"

// pass by value (coppy)
func changenum(num int) {
	num = 5
	fmt.Println(num)
}

// by refrence
func refChange(num *int) {
	*num = 5
}

func changeSlices(nums *[]int) {
	*nums = append(*nums, 5)
}

func main() {

	num := 1
	changenum(num)

	fmt.Println("number in main function", num)

	// pointers(pass by refrence)
	refChange(&num)
	fmt.Println("number address", &num)

	// -> refrence in for slices
	nums := []int{1, 2, 3}
	changeSlices(&nums)
	fmt.Println(nums)
}
