package main

```
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

 

Constraints:

haystack and needle consist only of lowercase English characters.
```

//直接在string中从 i 到 len(needle) 判断是否与needle相等，若相等返回 i , 以此循环 len(string) - len(needle)次
func strStr(haystack string, needle string) int {
	lengthSub := len(needle)
	if lengthSub == 0 {
		return 0
	}

	length := len(haystack)
	diffLength := length - lengthSub
	for i := 0; i <= diffLength; i++ {
		if haystack[i:i+lengthSub] == needle {
			return i
		}
	}
	return -1
}
