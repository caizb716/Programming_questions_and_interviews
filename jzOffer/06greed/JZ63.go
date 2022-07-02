package greed

//股票的最大利润
func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	s, max := 0, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] || s > 0 {
			s += (prices[i] - prices[i-1])
			if s < 0 {
				s = 0
			} else if s > max {
				max = s
			}
		}
	}
	return max
}
func MaxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minprice, max := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		}
		if prices[i]-minprice > max {
			max = prices[i] - minprice
		}
	}
	return max
}
