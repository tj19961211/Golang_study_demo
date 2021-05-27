pakcage main

// Given a string, find the first non-repeating character in it and return its index. If it doesn't exist, return -1.

// Examples:

// s = "leetcode"
// return 0.

// s = "loveleetcode"
// return 2.


func firstUniqChar(s string) int {
	if s == ""   {
		return -1
	}

	if len(s) == 1 {
		return s
	}

	list := [26]int{}
	for i := 0; i < len(s); i++ {
		list[s[i] - 'a']++
	}
	for i := 0; i < len(s); i++ {
		if list[s[i] - 'a'] == 1 {
			return i
		}
	}

	return -1
}
