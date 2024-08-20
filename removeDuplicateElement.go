package main

import "fmt"

func removeDuplicates(nums []int) int {
}

func main() {
	// case 1
	nums := []int{1, 1, 1, 2, 2, 3}
	k := removeDuplicates(nums)
	fmt.Println("k =", k, "nums =", nums[0:k])

	// case 2
	nums2 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	k2 := removeDuplicates(nums2)
	fmt.Println("k2 =", k2, "nums2 =", nums2[0:k2])
}
