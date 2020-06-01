# 快排的Golang实现

原理：快速排序的原理是，首先找到一个数pivot把数组‘平均’分成两组，使其中一组的所有数字均大于另一组中的数字，此时pivot在数组中的位置就是它正确的位置。然后，对这两组数组再次进行这种操作。

这里的思路：
    在原切片上进行数值的排序，选取两个指针，分别指向未排序数组的前端和末端(lo和hi表示)。遍历数组，如果`nums[i]`小于`nums[hi]`则`nums[i]`与`nums[lo]`换位，当遍历完数组时，将比较的`nums[hi]`与`nums[lo]`换位，这样在`nums[lo]`前方的则是比`nums[lo]`小的而后方则是比它大的数，而`nums[lo]`的位置就是已经排好序后的位置。
    最后再将`nums[lo]`左右两边再进行快排

```go
func partition(nums []int, lo, hi int) int{
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

func main() {
	list := []int{55, 90, 74, 20, 16, 46, 43, 59, 2, 99, 79, 10, 73, 1, 68, 56, 3, 87, 40, 78, 14, 18, 51, 24, 57, 89, 4, 62, 53, 23, 93, 41, 95, 84, 88}

	quickSort(list, 0, len(list)-1)
	fmt.Println(list)
}
```