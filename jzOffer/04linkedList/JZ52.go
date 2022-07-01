package linkedlist

//两个链表的第一个公共节点
//两个链表一起走，相遇的节点就是公共节点
func GetIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p, q := headA, headB
	for p != q {
		if p == nil {
			p = headB
		} else {
			p = p.Next
		}
		if q == nil {
			q = headA
		} else {
			q = q.Next
		}
	}
	return p
}

//方法二，利用map结构
func GetIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	m := make(map[ListNode]*ListNode)
	for headA != nil {
		m[*headA] = headA
		headA = headA.Next
	}
	for headB != nil {
		if m[*headB] == headB { //不能使用v,ok := m[*headB];ok方式判断，两个列表可能节点数据相同，但是地址不同
			return headB
		}
		headB = headB.Next
	}
	return nil
}
