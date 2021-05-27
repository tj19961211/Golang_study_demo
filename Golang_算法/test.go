package main

import "fmt"

func partition(nums []int, lo, hi int) int {
	p := nums[hi]
	for i := lo; i < hi; i++ {
		if nums[i] < p {
			nums[i], nums[lo] = nums[lo], nums[i]
			lo++
		}
	}
	nums[hi], nums[lo] = nums[lo], nums[hi]
	return lo
}

func quickSort(nums []int, lo, hi int) {
	if lo > hi {
		return
	}

	p := partition(nums, lo, hi)
	quickSort(nums, lo, p-1)
	quickSort(nums, p+1, hi)
}

func topk(nums []int, k, lo, hi int) int {
	if lo > hi {
		return 0
	}
	p := partition(nums, lo, hi)
	if p == k {
		return nums[p]
	}
	if p < k {
		quickSort(nums, p+1, hi)
	}
	if p > k {
		quickSort(nums, lo, p-1)
	}
	return nums[k]
}

func main() {
	list := []int{55, 90, 74, 20, 16, 46, 43, 59, 2, 99, 79, 10, 73, 1, 68, 56, 3, 87, 40, 78, 14, 18, 51, 24, 57, 89, 4, 62, 53, 23, 93, 41, 95, 84, 88, 19}

	// quickSort(list, 0, len(list)-1)
	// fmt.Println(list)
	fmt.Println(topk(list, 1, 0, len(list)-1))
	fmt.Println(list)
}
