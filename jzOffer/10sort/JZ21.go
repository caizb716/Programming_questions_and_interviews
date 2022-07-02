package sort

//调整数组顺序使奇数位于偶数前面
func Exchange(num []int) []int {
	for i, j := 0, 0; i < len(num); i++ {
		if num[i]&1 == 1 {
			num[i], num[j] = num[j], num[i]
			j++
		}
	}
	return num

}

//相对位置不改变
func Exchange2(num []int) []int {
	temp := make([]int, len(num))
	left := 0
	copy(temp, num)
	for i := 0; i < len(num); i++ {
		if num[i]&1 == 1 {
			num[i], num[left] = num[left], num[i]
			left++
		}
	}
	for m, n := len(temp)-1, len(temp)-1; m >= 0; m-- {
		if temp[m]&1 == 0 {
			temp[m], temp[n] = temp[n], temp[m]
			n--
		}
	}
	num = append(num[:left], temp[left:]...)
	return num

}
