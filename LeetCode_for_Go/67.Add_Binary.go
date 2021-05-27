package main

import (
	"fmt"
	"strconv"
)

// Given two binary strings, return their sum (also a binary string).

// The input strings are both non-empty and contains only characters 1 or 0.

// Example 1:

// Input: a = "11", b = "1"
// Output: "100"
// Example 2:

// Input: a = "1010", b = "1011"
// Output: "10101"

// Constraints:

// Each string consists only of '0' or '1' characters.
// 1 <= a.length, b.length <= 10^4
// Each string is either "0" or doesn't contain any leading zero.

// 方法一：使用字节码进行计算
func addBinary_1(a string, b string) string {
	i, j, res := len(a)-1, len(b)-1, []byte{}
	// 二进制值 (例：1010 )
	var carry byte
	for i >= 0 || j >= 0 || carry == 1 {
		if i >= 0 {
			// 获取 string 中的最后一个值(通过下标取出的是 byte 类型)
			// 因为传入的是二进制字符串，所以获取的下标基本为固定的字节码数值
			// 获取的字节码值减去 '0' 的字节码所剩下的即是 0 与 1 对应了二进制
			carry += a[i] - '0'
			i--
		}
		if j >= 0 {
			carry += b[j] - '0'
			j--
		}
		// 二进制值为 2 时，该位置为 0， 并向 res 中添加字节码
		res = append(res, carry%2+'0')
		// 若 carry 内的值为 2 的倍数，则除以 2 取整
		carry = carry / 2
	}
	return string(reverse(res))
}

// 反转 slice
func reverse(s []byte) []byte {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
	return s
}

func main() {
	fmt.Println(addBinary_1("100", "110"))
	fmt.Println(addBinary_2("100", "110"))
}

// 方法二： 使用 string 与 int 进行计算
func addBinary_2(a, b string) string {
	i, j, res, carry := len(a)-1, len(b)-1, "", 0
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		res = strconv.Itoa(sum%2) + res
		carry = sum / 2
	}
	if carry == 1 {
		return "1" + res
	}
	return res
}
