package mian

import "fmt"

// numbered sequence of specific length
func main() {
	var nums [4]int // default assign default values

	// default values / zeroed values
	// int = 0
	// bool = false
	// string = ""

	// arary legth
	fmt.Println(len(nums))

	var index = 0
	nums[index] = 1

	fmt.Println(nums[0])

	// print entire array
	fmt.Println(nums)

	// declare and add values
	newNum := [3]int{1, 2, 3}
	fmt.Println(newNum)

	// 2d array
	twoDnums := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(twoDnums)

	// [+] fixes size, that is predictable
	// [+] memory optimization

}
