package doublepointer

import "strings"

//反转单词顺序列
func ReverseWords(s string) string {
	stringList := strings.Split(s, " ")
	var res []string
	for i := len(stringList) - 1; i >= 0; i-- {
		if len(stringList[i]) > 0 {
			res = append(res, stringList[i])
		}

	}
	return strings.Join(res, " ")

}
