package stackqueueandheap

//实现一个栈，包括入栈push，出栈pop，查看栈顶top，查看最小值min
type MinStack struct {
	nums []int
	min  []int
}

/** initialize your data structure here. */
func NewMinStack() MinStack {
	return MinStack{
		nums: make([]int, 0),
		min:  make([]int, 0),
	}

}

func (minStack *MinStack) Push(x int) {
	minStack.nums = append(minStack.nums, x)
	if len(minStack.min) == 0 {
		minStack.min = append(minStack.min, x)
	} else if minStack.min[len(minStack.min)-1] < x {
		minStack.min = append(minStack.min, minStack.min[len(minStack.min)-1])
	} else {
		minStack.min = append(minStack.min, x)
	}

}

func (minStack *MinStack) Pop() {
	minStack.nums = minStack.nums[:len(minStack.nums)-1]
	minStack.min = minStack.min[:len(minStack.min)-1]
}

func (minStack *MinStack) Top() int {
	return minStack.nums[len(minStack.nums)-1]

}

func (minStack *MinStack) Min() int {
	return minStack.min[len(minStack.min)-1]

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
