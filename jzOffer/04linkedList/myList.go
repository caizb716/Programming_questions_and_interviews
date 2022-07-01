package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func newNode(val int) *ListNode {
	return &ListNode{Val: val}
}

//不带头结点的链表
func NewList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := newNode(nums[0])
	p := head
	for i := 1; i < len(nums); i++ {
		p.Next = newNode(nums[i])
		p = p.Next
	}
	return head
}
