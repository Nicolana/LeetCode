package main

import "fmt"

func removeDuplicates(nums []int) int {
	length := len(nums)
	count := 0
	prev := nums[0]
	j := 0
	for i := 0; i < length; i++ {
		num := nums[i]
		fmt.Println("Num =", num)
		if prev == num {
			count++
			if count <= 2 {
				// 两个值相同，计数器增加
				// 把这个值挪到j的位置，并且+1
				nums[j] = num
				j++
			}
		} else {
			count = 1
			prev = num
			nums[j] = num
			j++
		}
	}

	return j
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
