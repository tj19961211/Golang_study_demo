package main

import "fmt"

// 累加法
func getRow(rowIndex int) []int {
	ReArray := make([]int, rowIndex+1)
	// 创建的 ReArray 的初始值为 0
	// 为初始化 ReArray 的前两列的值，也就是 rowIndex 为 0时的初始值
	ReArray[0] = 1
	// i = 1 ReArray[1, 1, 0, 0]
	// i = 2 ReArray[1, 2, 1, 0]
	// i = 3 ReArray[1, 3, 3, 1]
	for i := 0; i < rowIndex+1; i++ {
		for j := i; j >= 1; j-- {
			fmt.Println(ReArray[j])
			ReArray[j] += ReArray[j-1]
		}
	}
	return ReArray
}

func main() {
	fmt.Println(getRow(4))
}
