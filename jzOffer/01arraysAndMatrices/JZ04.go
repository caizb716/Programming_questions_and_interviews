package arraysandmatrices

//二维数组中的查找
import (
	"sort"
)

func FindNumberIn2DArray(matrix [][]int, target int) (func() bool, func() bool) {

	func1 := func() bool {
		i := len(matrix) - 1
		j := 0
		for i > -1 {
			if j < len(matrix[i]) {
				if matrix[i][j] == target {
					return true
				} else if matrix[i][j] > target {
					i--
				} else {
					j++
				}
			} else {
				return false
			}
		}
		return false
	}
	func2 := func() bool {
		for _, nums := range matrix {
			i := sort.SearchInts(nums, target)
			if i < len(nums) && nums[i] == target {
				return true
			}

		}
		return false
	}
	return func1, func2

}
