package divideandconquer

//数值的整数次方

func help(x float64, n int) float64 {
	if n == 1 {
		return x
	} else if n&1 == 0 {
		//此时n一定是个偶数，减少迭代次数
		return help(x*x, n/2)
	} else {
		return x * help(x, n-1)
	}
}
func MyPow(x float64, n int) float64 {
	if n < 0 {
		return help(1.0/x, -n)
	} else if n == 0 {
		return 1
	} else {
		return help(x, n)
	}
}
