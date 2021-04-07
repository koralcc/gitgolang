package main

import (
	"fmt"
)

func main() {
	var nums1 = []int{1, 2, 3, 0, 0, 0}
	var nums2 = []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

// 1 3 4 //45 56
func merge(nums1 []int, m int, nums2 []int, n int) {
	var merge []int
	var p1 int //s1
	var p2 int //s2
	for p1 <= m-1 || p2 <= n-1 {
		if p1 <= m-1 && p2 <= n-1 {
			if nums1[p1] < nums2[p2] {
				merge = append(merge, nums1[p1])
				p1++
			} else if nums1[p1] == nums2[p2] {
				fmt.Println(nums1[p1])
				merge = append(merge, nums1[p1], nums2[p2])
				p1++
				p2++
			} else if nums1[p1] > nums2[p2] {
				merge = append(merge, nums2[p2])
				p2++
			}
		}

		if p1 > m-1 && p2 <= n-1 {
			merge = append(merge, nums2[p2])
			p2++
		}

		if p2 > n-1 && p1 <= m-1 {
			merge = append(merge, nums1[p1])
			p1++
		}
	}
	//fmt.Println("merge", merge)
	//nums1 = merge 直接赋值会发生逃逸
	copy(nums1, merge)
}
