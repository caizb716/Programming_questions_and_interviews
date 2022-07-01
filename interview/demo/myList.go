package demo

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListWithHead(num []int) *ListNode {
	head := &ListNode{-1, nil}
	p := head
	for _, v := range num {
		p.Next = &ListNode{v, nil}
		p = p.Next
	}
	return head
}
func NewListWithoutHead(num []int) *ListNode {
	head := &ListNode{num[0], nil}
	p := head
	for i := 1; i < len(num); i++ {
		p.Next = &ListNode{num[i], nil}
		p = p.Next
	}
	return head
}
func (l *ListNode) Insert(num int) {
	for l.Next != nil {
		l = l.Next
	}
	l.Next = &ListNode{num, nil}
}
func (l *ListNode) Delete(num int) *ListNode {
	if l.Val == num {
		return l.Next
	}
	p := l
	for p.Next.Val != num {
		p = p.Next
	}
	tem := p.Next
	p.Next = tem.Next
	tem.Next = nil
	return l
}
