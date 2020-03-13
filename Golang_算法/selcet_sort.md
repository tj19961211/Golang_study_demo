## 选择排序

选择排序的原理是，对给定的数组进行多次遍历，每次均找出最大的一个值的索引。

```go
func selectSort(nums []int) {
	k := len(nums)
	for j := 0; j < k; j++ {
		maxIndex := k - 1 - j
    	for i := 0; i < k-j; i++ {
			if nums[i] > nums[maxIndex] {
				nums[i], nums[maxIndex] = nums[maxIndex], nums[i] 
			}	
		}
	}
}
```