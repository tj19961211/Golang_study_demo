package main

import "fmt"

func countPrimes(n int) int {
	var count int
	notPrime := make([]bool, n)
	// 赋初始值
	notPrime[0], notPrime[1] = true, true

	// 质数为 true ，非质数为 false
	for i := 2; i < n; i++ {
		if notPrime[i] {
			continue
		}
		count++
		// 从 2 开始对每个数剩以 2 后都为非质数，以及后面的每隔 i 个数都不为质数
		for j := i * 2; j < n; j += i {
			notPrime[j] = true
		}
	}

	return count
}

func main() {
	fmt.Println(countPrimes(10))
}
