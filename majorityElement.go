package main

import "fmt"

func majorityElement(nums []int) int {
	n := len(nums)
	num := n / 2
	m := make(map[int]int) // 用于统计数字出现的次数
	for i := 0; i < n; i++ {
		m[nums[i]]++
	}

	// 遍历map，找到出现次数最多的数字
	for key, value := range m {
		if value > num {
			return key
		}
	}
	return -1
}

func main() {
	// case 1
	num := []int{3, 2, 3}
	fmt.Println("case 1 =", majorityElement(num))

	// case 2
	nums2 := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println("case 2 =", majorityElement(nums2))
}
