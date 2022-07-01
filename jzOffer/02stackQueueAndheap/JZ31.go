package stackqueueandheap

//栈的压入、弹出序列

func ValidateStackSequences(pushed []int, popped []int) bool {

	stack := make([]int, 0)
	i := 0
	for _, value := range pushed {
		stack = append(stack, value)
		for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}
	if len(stack) == 0 && i == len(popped) {
		return true
	}
	return false

}
