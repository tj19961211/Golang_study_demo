# 冒泡排序

冒泡排序的原理是，对给定的数组进行多次遍历，每次均比较相邻的两个数，如果前一个比后一个大，则交换这两个数。经过第一次遍历之后，最大的数就在最右侧了；第二次遍历之后，第二大的数就在右数第二个位置了；以此类推
```go
func bobbleSort(nums []int) {
    for i := 0; i < len(nums); i++ {
        for j := 1; j < len(nums) - i; j++ {
            if nums[j] < nums[j-1] {
                 nums[j], nums[j-1] = nums[j-1], nums[j]
            }
        }
    }
}
```