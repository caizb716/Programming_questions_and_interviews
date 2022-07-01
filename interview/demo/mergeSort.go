package demo

func MergeSort(num []int, start, end int) {
	if start >= end {
		return
	}
	mid := start + (end-start)/2
	MergeSort(num, start, mid)
	MergeSort(num, mid+1, end)
	temp := make([]int, 0)
	left, right := start, mid+1
	for left <= mid && right <= end {
		if num[left] <= num[right] {
			temp = append(temp, num[left])
			left++
		} else {
			temp = append(temp, num[right])
			right++
		}
	}
	for ; left <= mid; left++ {
		temp = append(temp, num[left])
	}
	for ; right <= end; right++ {
		temp = append(temp, num[right])
	}
	for i := start; i <= end; i++ {
		num[i] = temp[i-start]
	}
}
