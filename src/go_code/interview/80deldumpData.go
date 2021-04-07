package main

import "fmt"

func main() {
	var nums = []int{1, 1, 1, 2, 2, 3}
	removeDuplicates(nums)
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
