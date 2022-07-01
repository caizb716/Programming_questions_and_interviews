package stackqueueandheap

//数据流中的中位数
type MedianFinder struct {
	minHeap []int
	maxHeap []int
	length  int
}

/** initialize your data structure here. */
func NewMedianFinder() MedianFinder {
	return MedianFinder{
		minHeap: []int{},
		maxHeap: []int{},
		length:  0,
	}
}

//用于插入元素调整堆
func upAdjust(arr []int, cmp func(x, y int) bool) {
	if len(arr) == 0 {
		return
	}
	length := len(arr)
	child := length - 1
	parent := (child - 1) / 2
	newValue := arr[child]
	for child > 0 {
		if cmp(newValue, arr[parent]) {
			arr[child] = arr[parent]
			child = parent
			parent = (child - 1) / 2
		} else {
			break
		}
	}
	arr[child] = newValue
}
func dowmAdjust(arr []int, i int, cmp func(x, y int) bool) {
	if len(arr) == 0 {
		return
	}
	parent := i
	child := 2*i + 1
	rootValue := arr[parent]
	for child < len(arr) {
		if child+1 < len(arr) && cmp(arr[child+1], arr[child]) {
			child++
		}
		if cmp(arr[child], rootValue) {
			arr[parent] = arr[child]
			parent = child
			child = 2*parent + 1
		} else {
			break
		}
	}
	arr[parent] = rootValue
}
func pop(arr []int) []int {
	length := len(arr)
	if length == 1 {
		return arr[:length-1]
	}

	arr[0] = arr[length-1]
	return arr[:length-1]

}

//维护大顶堆
func cmpMax(x, y int) bool {
	return x > y
}

//维护小顶堆
func cmpMin(x, y int) bool {
	return x < y
}
func (mFinder *MedianFinder) AddNum(num int) {
	mFinder.length++
	if mFinder.length%2 == 0 {
		//存入小顶堆并调整
		mFinder.minHeap = append(mFinder.minHeap, num)
		upAdjust(mFinder.minHeap, cmpMin)
		//弹出小顶堆堆顶元素，并存入大顶堆
		mFinder.maxHeap = append(mFinder.maxHeap, mFinder.minHeap[0])
		mFinder.minHeap = pop(mFinder.minHeap)
		//调整弹出元素的小顶堆，和新加元素的大顶堆
		upAdjust(mFinder.maxHeap, cmpMax)
		dowmAdjust(mFinder.minHeap, 0, cmpMin)
	} else if mFinder.length%2 == 1 {
		//存入大顶堆并调整
		mFinder.maxHeap = append(mFinder.maxHeap, num)
		upAdjust(mFinder.maxHeap, cmpMax)
		//弹出大顶堆堆顶元素，并存入小顶堆
		mFinder.minHeap = append(mFinder.minHeap, mFinder.maxHeap[0])
		mFinder.maxHeap = pop(mFinder.maxHeap)
		//调整弹出元素的大顶堆，和新加元素的小顶堆
		upAdjust(mFinder.minHeap, cmpMin)
		dowmAdjust(mFinder.maxHeap, 0, cmpMax)
	}
}

func (mFinder *MedianFinder) FindMedian() float64 {
	if mFinder.length%2 == 1 {
		return float64(mFinder.minHeap[0])
	}
	if mFinder.length%2 == 0 {
		return (float64(mFinder.maxHeap[0]) + float64(mFinder.minHeap[0])) / 2
	}
	return 0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
