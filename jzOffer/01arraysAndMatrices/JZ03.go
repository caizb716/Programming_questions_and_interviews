package arraysandmatrices

//数组中重复数字
import (
	"sort"
)

func FindRepeatNumber(nums []int) (func() int, func() int, func() int) {

	func1 := func() int {
		m := make(map[int]int)
		for _, v := range nums {
			if _, ok := m[v]; ok {
				return v
			} else {
				m[v] = 1
			}
		}
		return -1
	}
	func2 := func() int {
		sort.Ints(nums)
		for i, numsSize := 0, len(nums); i < numsSize-1; i++ {
			if nums[i] == nums[i+1] {
				return nums[i]
			}
		}
		return -1
	}
	func3 := func() int {
		for i := 0; i < len(nums); i++ {
			for i != nums[i] {
				if nums[i] == nums[nums[i]] {
					return nums[i]
				}
				nums[i], nums[nums[i]] = nums[nums[i]], nums[i]

			}
		}
		return -1
	}
	return func1, func2, func3
}
