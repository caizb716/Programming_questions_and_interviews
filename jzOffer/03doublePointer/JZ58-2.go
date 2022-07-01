package doublepointer

//左旋转字符串
//字符串第i个位置变为（i+n）%len(s)，原字符串第i个位置变为（i-n+len(s)）%len(s)
func ReverseLeftWords1(s string, n int) string {
	l := len(s)
	str := []rune(s)
	counter := 0
	for i := 0; counter != l; i++ {
		t := str[i]
		for j := (i + n) % l; j != i; j = (j + n) % l {
			str[(j-n+l)%l] = str[j]
			counter++
		}
		str[(i-n+l)%l] = t
		counter++
	}
	return string(str)
}

//解法二，太麻烦，不推荐
func ReverseLeftWords2(s string, n int) string {
	str := []rune(s)
	recv(str, n)
	return string(str)
}

func swap(str []rune, i, j, n int) {
	for k := n; k != 0; k-- {
		str[i], str[j] = str[j], str[i]
		i++
		j++
	}
}

func recv(str []rune, n int) {
	if n == 0 {
		return
	}
	var rst, i, j int
	if n > len(str)/2 || (len(str)%2 == 0 && n == len(str)/2) {
		rst = len(str) - n
		i, j = len(str)-2*rst, len(str)-rst
		n -= rst
	} else {
		rst = n
		i, j = 0, len(str)-rst
	}
	swap(str, i, j, rst)
	recv(str[:j], n)
}

//解法三，反转反转再反转
//执行最快
func ReverseLeftWords3(s string, n int) string {
	str := []rune(s)
	reverse(str[:n])
	reverse(str[n:])
	reverse(str)
	return string(str)
}
func reverse(str []rune) {
	i, j := 0, len(str)-1
	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
}
