package main

import "fmt"

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	var arr4 [4][5]int

	for i, v := range arr3 {
		fmt.Println(arr3[i])
		fmt.Println(i, v)
	}

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(arr4)
	printArray(arr1)
}
