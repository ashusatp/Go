package main

import (
	"fmt"
	"slices"
)

// slice -> dinamic array
func main() {

	//[+] uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums == nil) //-> true

	// not want to make it nil

	// [+] make
	var nums = make([]int, 0) // fill the initial size that is 2
	fmt.Println(cap(nums))    // print the capacity that is 2
	fmt.Println(nums == nil)  //-> false

	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	fmt.Println(nums)

	var nums2 = make([]int, 2)
	nums2[1] = 1
	nums2[0] = 0
	nums2 = append(nums2, 10)
	fmt.Println(nums2) // [0,1,10]

	// [+] sub slice
	subSlice := nums[1:4] // 0 to 3
	fmt.Println(subSlice)
	//[:5] from 0 to 4
	//[1:] from 1 to end

	//[+] copy
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)   //copy into dst from src
	fmt.Println(dst) // Output: [1, 2, 3]

	// [+] Slices package

	//-> slices.Contains
	numbers := []int{1, 2, 3, 4, 5}
	found := slices.Contains(numbers, 3)
	fmt.Println(found) // Output: true

	//-> slices.Sort
	slices.Sort(numbers)
	fmt.Println(numbers) // Output: [1, 2, 3, 4, 5]

	//-> slices.Equal
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	equal := slices.Equal(slice1, slice2)
	fmt.Println(equal) // Output: true

	//-> slices.Reverse
	slices.Reverse(numbers)
	fmt.Println(numbers) // Output: [5, 4, 3, 2, 1]

	//-> slices.BinarySearch
	index, found := slices.BinarySearch(numbers, 3)
	fmt.Println(index, found) // Output: 2 true

	//-> slices.Insert //Inserts elements into a slice at the specified index.
	temp := []int{1, 2, 4, 5}
	temp = slices.Insert(temp, 2, 3)
	fmt.Println(temp) // Output: [1, 2, 3, 4, 5]

	//-> slices.Delete //Removes elements from a slice at the specified range.
	temp = slices.Delete(temp, 1, 3) // Removes elements at index 1 and 2
	fmt.Println(temp)                // Output: [1, 4, 5]

	//-> slices.Clone //Creates a copy of the slice.

	original := []int{1, 2, 3}
	clone := slices.Clone(original)
	fmt.Println(clone) // Output: [1, 2, 3]

	//-> slices.Compact //Removes consecutive duplicate elements in a slice.

	digits := []int{1, 2, 2, 3, 3, 3, 4}
	result := slices.Compact(digits)
	fmt.Println(result) // Output: [1, 2, 3, 4]

}
