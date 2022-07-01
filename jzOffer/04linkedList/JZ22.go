package linkedlist

//链表中倒数第K个节点

//双指针
func GetKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
