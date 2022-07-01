package doublepointer

//和为s的连续正数序列
func FindContinuousSequence(target int) [][]int {
	var res [][]int
	left, right, mid := 1, 2, (target+1)/2
	sum := left + right
	for left < mid {
		if sum == target {
			var tmp []int
			for i := left; i <= right; i++ {
				tmp = append(tmp, i)
			}
			res = append(res, tmp)
			right++
			sum += right
		} else if sum < target {
			right++
			sum += right
		} else {
			sum -= left
			left++
		}
	}
	return res
}
