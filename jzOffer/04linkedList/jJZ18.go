package linkedlist

//删除链表节点

//单指针
func DeleteNode1(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	pre := head
	for pre.Next != nil && pre.Next.Val != val {
		pre = pre.Next
	}
	if pre.Next != nil {
		pre.Next = pre.Next.Next
	}
	return head
}

//双指针
func DeleteNode2(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	pre, cur := head, head.Next
	for cur != nil && cur.Val != val {
		pre, cur = cur, cur.Next
	}
	if cur != nil {
		pre.Next = cur.Next
	}
	return head
}

//递归
func DeleteNode3(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	head.Next = DeleteNode3(head.Next, val)
	return head
}
