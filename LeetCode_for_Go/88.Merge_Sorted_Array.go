package main

import "fmt"

// 2020-10-7
// Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

// Note:

// The number of elements initialized in nums1 and nums2 are m and n respectively.
// You may assume that nums1 has enough space (size that is equal to m + n) to hold additional elements from nums2.
// Example:

// Input:
// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3

// Output: [1,2,2,3,5,6]

// Constraints:

// -10^9 <= nums1[i], nums2[i] <= 10^9
// nums1.length == m + n
// nums2.length == n

func merge(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		// 从两个 slice 的尾部开始对比
		// 因为 slice 的最后长度为 m + n
		// 所以比较后较大的一个数会放到最后的位置，即 m+n-1
		if m == 0 || nums2[n-1] > nums1[m-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
}

func main() {
	var i = []int{1, 2, 3, 0, 0, 0}
	merge(i, 3, []int{4, 5, 6}, 3)
	fmt.Println(i)
}
