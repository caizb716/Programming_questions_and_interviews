package cider

import (
	"fmt"
)

//字符串中的最长回文子字符串
func LongestPalindrome(str string) string {
	newStr := "#"
	for i := 0; i < len(str); i++ {
		newStr = fmt.Sprintf("%s%c#", newStr, str[i])
	}
	max, start, slen := 1, 0, len(newStr)
	for i, _ := range newStr {
		left, right := i-1, i+1
		step := 0
		for left >= 0 && right < slen && newStr[left] == newStr[right] {
			left--
			right++
			step++
		}
		if step > max {
			max = step
			start = (i - max) / 2
		}
	}
	return str[start : start+max]

}
