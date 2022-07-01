package search

//机器人的运动范围
func MovingCount(m int, n int, k int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	return dfs(m, n, 0, 0, k, dp)
}

func dfs(m, n, i, j, k int, dp [][]int) int {
	if i < 0 || j < 0 || i >= m || j >= n || dp[i][j] != 0 || (sumPos(i)+sumPos(j)) > k {
		return 0
	}
	dp[i][j] = 1
	sum := 1
	sum += dfs(m, n, i+1, j, k, dp)
	sum += dfs(m, n, i-1, j, k, dp)
	sum += dfs(m, n, i, j+1, k, dp)
	sum += dfs(m, n, i, j-1, k, dp)
	return sum
}

func sumPos(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
