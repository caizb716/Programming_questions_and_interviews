package cider

//斐波那契数列
func Fibonacci(n int) []int {
	res := []int{1, 1}
	if n == 1 || n == 2 {
		return res[:n]
	}
	for i := 2; i < n; i++ {
		res = append(res, res[i-1]+res[i-2])
	}
	return res
}
