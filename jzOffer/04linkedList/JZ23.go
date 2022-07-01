package linkedlist

//链表中环的入口
func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	slow, fast := pHead, pHead
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = pHead
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
