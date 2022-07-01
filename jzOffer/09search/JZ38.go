package search

func Permutation(s string) []string {
	res := make(map[string]bool)
	c := []byte(s)
	dfs2(res, c, 0)
	RES := []string{}
	for key := range res {
		RES = append(RES, key)
	}
	return RES
}
func dfs2(res map[string]bool, c []byte, i int) {
	if i == len(c)-1 {
		res[string(c)] = true //完成了剪纸操作
		return
	} else {
		for j := i; j < len(c); j++ {
			c[i], c[j] = c[j], c[i]
			dfs2(res, c, i+1)
			c[i], c[j] = c[j], c[i]
		}
	}
}
