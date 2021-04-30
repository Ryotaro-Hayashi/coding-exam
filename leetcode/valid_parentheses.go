package leetcode

import (
	"strings"
)

/*
   Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

	An input string is valid if:

	Open brackets must be closed by the same type of brackets.
	Open brackets must be closed in the correct order.
*/

func isValid(s string) bool {
	sslice := strings.Split(s, "") // 文字列をスライスに格納
	charSet := map[string]string{"(": ")", "{": "}", "[": "]"} // 括弧の組

	stack := make([]string, 0, len(sslice))
	for _, figure := range sslice {
		if _, ok := charSet[figure]; ok { // 右括弧のとき
			stack = append(stack, figure) // push
		} else { // 左括弧のとき
			if len(stack) == 0 { // stackが空で左括弧
				return false
			}
			if charSet[stack[len(stack)-1]] != figure { // ペアが不一致
				return false
			} else { // ペアが一致したらpop
				stack = stack[:len(stack)-1] // pop
			}
		}
	}
	return len(stack) == 0
}
