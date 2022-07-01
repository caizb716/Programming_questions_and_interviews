package stackqueueandheap

//最小的K个数
func GetLeastNumbersByMaxHeap(arr []int, k int) []int {
	length := len(arr)
	if k > length {
		return arr
	}
	//创建大顶堆
	maxHeap := make([]int, 0)
	maxHeap = append(maxHeap, arr[:k]...)
	createMaxHeap(maxHeap)
	//遍历数组剩余数据，小于堆顶数据时，替换堆顶，重新维护小顶堆
	for i := k; i < length && len(maxHeap) > 0; i++ {
		if arr[i] < maxHeap[0] {
			maxHeap[0] = arr[i]
			maxHeapify(maxHeap, 0, k)
		}
	}
	return maxHeap
}

//自底向上创建大顶堆
func createMaxHeap(nums []int) {
	length := len(nums)
	//从最后一个非叶子节点开始调整，最后一个非叶子节点为第(length-1)/2个节点
	for i := length/2 - 1; i >= 0; i-- {
		maxHeapify(nums, i, length)
	}
}

func maxHeapify(nums []int, posIndex, length int) {
	//左孩子节点index
	leftIndex := posIndex*2 + 1
	//右孩子节点index
	rightnIndex := posIndex*2 + 2
	//当前节点以及左右孩子节点中最大值的节点，初始化为当前节点index
	maxIndex := posIndex
	//左孩子存在并且大于当前节点值时，最大值index替换为左孩子index
	if leftIndex < length && nums[leftIndex] > nums[maxIndex] {
		maxIndex = leftIndex
	}
	//右孩子存在并且大于当前节点值时，最大值index替换为右孩子index
	if rightnIndex < length && nums[rightnIndex] > nums[maxIndex] {
		maxIndex = rightnIndex
	}
	// 最小值节点index不等于当前节点时,替换当前节点和其中较小孩子节点值
	if maxIndex != posIndex {
		nums[posIndex], nums[maxIndex] = nums[maxIndex], nums[posIndex]
		//递归调用，把当前节点的子节点调整为大顶堆
		maxHeapify(nums, maxIndex, length)
	}
}

func GetLeastNumbersByQuickSort(arr []int, k int) []int {
	if len(arr) == 0 || k >= len(arr) {
		return arr
	}
	left := 0
	right := len(arr) - 1
	index := partation(arr, left, right)
	for index != k-1 {
		if index > k-1 {
			right = index - 1
			index = partation(arr, left, right)
		} else {
			left = index + 1
			index = partation(arr, left, right)
		}
	}
	result := make([]int, 0)
	result = append(result, arr[:k]...)
	return result
}
func partation(num []int, left, right int) int {
	if left == right {
		return right
	}
	if left < right {
		// 1.找到基准元素
		base := num[left]
		l := left
		r := right
		for {
			// 2. 如果右侧的值大于base ，尾指针 向前移动
			for num[r] >= base && l < r {
				r--
			}
			// 3.如果左侧的值小于base , 头指针向后移动
			for num[l] <= base && l < r {
				l++
			}
			// 交换一下
			if l >= r {
				//fmt.Println("break")
				break

			}
			num[l], num[r] = num[r], num[l]
		}
		// 基准值为找到的那个值
		num[left] = num[l]
		num[l] = base
		return l
	}
	return -1
}
