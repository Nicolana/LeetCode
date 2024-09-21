package main

import "fmt"

func majorityElement(nums []int) int {
	candidate := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		value := nums[i]
		if count == 0 {
			candidate = value
			count++
		} else if candidate == value {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func main() {
	// case 1
	num := []int{3, 2, 3}
	fmt.Println("case 1 =", majorityElement(num))

	// case 2
	nums2 := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println("case 2 =", majorityElement(nums2))
}
