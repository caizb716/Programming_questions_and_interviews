package doublepointer

//和为s的两个数字
func TwoSumByHash(nums []int, target int) []int {
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := m[target-v]; ok {
			return []int{v, target - v}
		} else {
			m[v] = struct{}{}
		}
	}
	return nil
}

func TwoSumByTwoPointer(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{nums[left], nums[right]}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
