package stackqueueandheap

//滑动窗口的最大值
func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k < 0 || k > len(nums) {
		return nil
	}
	maxValue := []int{}
	max := -1
	for i := 0; i <= len(nums)-k; i++ {
		l := i
		r := i + k - 1
		if max == -1 || max == l-1 {
			max = getMax(nums, l, r)
		} else {
			if nums[max] < nums[r] {
				max = r
			}
		}
		maxValue = append(maxValue, nums[max])
	}
	return maxValue
}
func getMax(nums []int, left, right int) int {
	for left < right {
		if nums[left] < nums[right] {
			left++
		} else {
			right--
		}
	}
	return left
}
