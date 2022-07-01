package binarysearch

//旋转数组的最小的数字
func MinArray(numbers []int) int {
	low, high := 0, len(numbers)-1
	for low < high {
		mid := (low + high) / 2
		if numbers[mid] < numbers[high] {
			high = mid
		} else if numbers[mid] > numbers[high] {
			low = mid + 1
		} else {
			high--
		}
	}
	return numbers[high]
}
