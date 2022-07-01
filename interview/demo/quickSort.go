package demo

func QuickSort(num []int, start, end int) {
	if start >= end {
		return
	}
	base, index := num[start], start
	left, right := start, end
	for left < right {
		for left < right {
			if num[right] < base {
				num[index] = num[right]
				index = right
				right--
				break
			}
			right--
		}
		for left < right {
			if num[left] > base {
				num[index] = num[left]
				index = left
				left++
				break
			}
			left++
		}
	}
	num[index] = base
	QuickSort(num, start, index-1)
	QuickSort(num, index+1, end)
}
