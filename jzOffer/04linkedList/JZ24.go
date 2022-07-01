package linkedlist

//反转列表

//迭代
func ReverseList1(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}

//递归
func ReverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := ReverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
