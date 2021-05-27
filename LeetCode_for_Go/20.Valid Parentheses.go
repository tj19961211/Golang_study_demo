package main

```
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
 

Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
Example 4:

Input: s = "([)]"
Output: false
Example 5:

Input: s = "{[]}"
Output: true
 

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
```

func isValid(s string) int {
	tack := make([]rune, len(s))
	top := 0

	for _, v := range s {
		switch (v) {
		case '(':
			tack[top] = ')'
			top += 1
			break
		case '[':
			tack[top] = ']'
			break
		case '{':
			tack[top] = '}'
			break
		default:
			if top == 0 || tack[top-1] != v {
				return false
			}

			top -= 1
			break
		}
	}
	return top == 0
}