package arraysandmatrices

// 替换空格

import (
	"strings"
)

func ReplaceSpace(s string) (func() string, func() string) {
	func1 := func() string {
		var strs []string = strings.Split(s, " ")
		var res string = ""
		for i, v := range strs {
			if i == len(strs)-1 {
				res += v
				break
			}
			res += v + "%20"
		}
		return res
	}
	func2 := func() string {
		res := ""
		for _, v := range s {
			if v == ' ' {
				res += "%20"
			} else {
				res += string(v)
			}
		}
		return res
	}
	return func1, func2
}
