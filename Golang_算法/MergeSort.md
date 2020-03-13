# 归并排序

归并的思想是先拆分，后组合，所以再mergeSort中运用递归进行数组的拆分，然后再调用merge进行数组的组合

```go
func merge(a, b []int) []int {
    var r = make([]int, len(a)+len(b))
    var i = 0
    var j = 0

    //对两个拆分数组进行值的比较，然后把较小的放到 r 切片里，然后指针往后移
    for i < len(a) && j < len(b) {
        if a[i] <= b[j] {
            r[i+j] = a[i]
            i++
        }else{
            r[i+j] = b[j]
            j++
        }
    }

    //对还有剩下的数组的值放进 r 切片内
    for i < len(a) {
        r[i+j] = a[i]
        i++
    }

    for j < len(b) {
        r[i+j] = b[j]
        j++
    }
    return r
}

func mergeSort(nums []int) []int {
    if len(nums) < 2 {
        return nums
    }

    var middle = len(nums) / 2
    var a = mergeSort(nums[:middle])
    var b = mergeSort(nums[middle:])
    return merge(a, b)
}
```