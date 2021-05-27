package main

// Write an algorithm to determine if a number n is happy.

// A happy number is a number defined by the following process:

// Starting with any positive integer, replace the number by the sum of the squares of its digits.
// Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy.
// Return true if n is a happy number, and false if not.

func isHappy(n int) bool {
	m := make(map[int]bool)
	tmp := n
	for tmp != 1 {
		// 倘若 map 中有一样的数值，则会形成无限循坏，需要手动退出
		if _, ok := m[tmp]; ok {
			return false
		}
		m[tmp] = true
		tmp = helper(tmp)
	}
	return true
}

func helper(n int) int {
	rst := 0
	for n != 0 {
		// 取个位数相乘(计算平方)
		rst += (n % 10) * (n % 10)

		// 取整数计算下一个循环
		n /= 10
	}
	return rst
}
