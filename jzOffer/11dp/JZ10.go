package dp

//斐波那契数列
func Fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	res := make([]int, n+1)
	res[0] = 0
	res[1] = 1
	for i := 2; i <= n; i++ {
		res[i] = (res[i-1] + res[i-2]) % 1000000007
	}
	return res[n]
}

//矩阵覆盖
//使用n个2*1的小矩阵覆盖2*n的大矩阵有多少种方法，输入个数，输出种类
func RectCover(number int) int {
	if number <= 2 {
		return number
	}
	dp := make([]int, number+1)
	dp[0], dp[1], dp[2] = 0, 1, 2
	for i := 3; i <= number; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[number]
}

//跳台阶
func NumWays(n int) int {
	if n < 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}
	return dp[n]
}

//变态跳台阶
func JumpFloorII(n int) int {
	if n < 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = 2 * dp[i-1]
	}
	return dp[n]
}
