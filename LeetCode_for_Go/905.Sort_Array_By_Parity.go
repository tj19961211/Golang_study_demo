package main

// 超级消耗内存和时间，主要在不断创建新的切片和 append 操作
func helper_1(A []int, B *[]int) {
	if len(A) == 0 {
		return
	}

	if A[0]%2 == 0 {
		// 放在数组的头部
		*B = append([]int{A[0]}, *B...)
	} else {
		// 放在数组的尾部
		*B = append(*B, A[0])
	}
	helper_1(A[1:], B)
}

func sortArrayByParity(A []int) []int {
	var B []int
	helper_1(A, &B)
	return B
}

// 第二种 快捷的方法:

// golang 位运算与准则 相同不变，不同为零
func sortArrayByParity_2(A []int) []int {
	for low, high := 0, len(A)-1; low < high; {
		//当前未判断的切片最左端 判断为偶数 不做处理
		for ; (A[low]&1 == 0) && low < high; low++ {
		}

		// 当前未判断的切片最右端 判断为奇数 不做处理
		for ; (A[high]&1 == 1) && low < high; high-- {
		}

		A[low], A[high] = A[high], A[low]
	}
	return A
}
