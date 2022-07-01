package linkedlist

//聪尾到头打印链表

func ReversePrint1(head *ListNode) []int {
	if head == nil {
		return nil
	}
	res := ReversePrint1(head.Next)
	res = append(res, head.Val)
	return res
}

func ReversePrint2(head *ListNode) []int {
	if head == nil {
		return nil
	}
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}
