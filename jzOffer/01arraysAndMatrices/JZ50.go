package arraysandmatrices

import "fmt"

// 第一个只出现一次的字符位置

func FirstUniqChar(s string, n int) byte {
	switch n {
	case 1:
		return firstUniqChar1(s)
	case 2:
		return firstUniqChar2(s)
	default:
		fmt.Println("n is not 1 or 2")
	}
	return ' '
}

func firstUniqChar1(s string) byte {

	list := [26]int{}
	for _, v := range s {
		list[v-'a']++
	}
	for i, v := range s {
		if list[v-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

func firstUniqChar2(s string) byte {
	n := len(s)
	pos := [26]int{}
	for i := range pos[:] {
		pos[i] = n
	}
	for i, ch := range s {
		ch -= 'a'
		if pos[ch] == n {
			pos[ch] = i
		} else {
			pos[ch] = n + 1
		}
	}
	ans := n
	for _, p := range pos[:] {
		if p < ans {
			ans = p
		}
	}
	if ans < n {
		return s[ans]
	}
	return ' '
}
