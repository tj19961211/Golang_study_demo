# 二分查找

```go
//第一种方法
func binarySearch(nums []int, target, loIndex, hiIndex int) int {
    if hiIndex < loIndex {
        return -1
    }

    mid := int((loIndex + hiIndex) / 2)
    if nums[mid] < target {
        return brnarySearch(nums, target, mid+1, hiIndex)
    }else if nums[mid] > target {
            return brnarySearch(nums, target, loIndex, mid)
    }else {
        return mid
    }
}

//第二种方法
func iterbinarySearch(nums []int, target, loIndex, hiIndex int) int {
    startIndex := loIndex
    endIndex := hiIndex
    var mid int
    for startIndex < endIndex {
        mid = int((startIndex+endIndex)/2)
        if nums[mid] > target {
            endIndex = mid
        }else if nums[mid] < target {
            startIndex = mid
        }else {
            return mid
        }
    }
    return -1
}
```