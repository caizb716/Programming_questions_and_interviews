package greed

//减绳子，将绳子剪成m段，乘积最大
func CuttingRope(n int) int {
	dp := make(map[int]int)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < (i/2 + 1); j++ {
			dp[i] = max(dp[i], max(dp[j], j)*max(dp[i-j], i-j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
