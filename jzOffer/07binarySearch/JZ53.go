package binarysearch

//数字在排序数组中出现的次数
func GetNumberOfK(data []int, k int) int {
	if len(data) == 0 {
		return 0
	}
	low, high := 0, len(data)-1
	mid := (low + high) / 2
	for data[mid] != k {
		if low < high {
			if data[mid] < k {
				low = mid + 1
			} else {
				high = mid
			}
		} else {
			return 0
		}
		mid = (low + high) / 2
	}
	total := 0
	for i := mid; i >= 0; i-- {
		if data[i] == k {
			total++
		} else {
			break
		}
	}
	for i := mid + 1; i < len(data); i++ {
		if data[i] == k {
			total++
		} else {
			break
		}
	}
	return total
}
