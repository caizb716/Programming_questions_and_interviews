package linkedlist

//合并两个排序的链表
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1}
	curr := dummyHead
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				curr.Next = l1
				curr = curr.Next
				l1 = l1.Next
			} else {
				curr.Next = l2
				curr = curr.Next
				l2 = l2.Next
			}
		} else if l1 != nil {
			curr.Next = l1
			break
		} else {
			curr.Next = l2
			break
		}
	}
	return dummyHead.Next
}
