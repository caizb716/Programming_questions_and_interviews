package linkedlist

//复杂链表的复制
//方法一：两次遍历
func CopyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	newHead := Node{
		Val:    head.Val,
		Next:   nil,
		Random: nil,
	}
	pre, p := &newHead, head.Next
	for p != nil {
		newNode := &Node{
			Val:    p.Val,
			Next:   nil,
			Random: nil,
		}
		pre.Next = newNode
		p = p.Next
		pre = pre.Next
	}
	p = head
	newP := &newHead
	for p != nil {
		if p.Random != nil {
			step := findStep(head, p.Random)
			newP.Random = findTarget(&newHead, step)
		}
		p = p.Next
		newP = newP.Next
	}
	return &newHead
}
func findStep(head, target *Node) int {
	p := head
	step := 0
	for p != target {
		p = p.Next
		step++
	}
	return step
}
func findTarget(head *Node, step int) *Node {
	p := head
	for i := step; i > 0; i-- {
		p = p.Next
	}
	return p
}

//方法二：使用map结构，保存原节点与新节点的对应关系
func CopyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	newHead := Node{
		Val:    head.Val,
		Next:   nil,
		Random: nil,
	}
	p := head.Next
	pre := &newHead
	m := make(map[*Node]*Node)
	m[head] = pre
	for p != nil {
		newNode := &Node{
			Val:    p.Val,
			Next:   nil,
			Random: nil,
		}
		pre.Next = newNode
		m[p] = newNode
		p = p.Next
		pre = pre.Next
	}
	p = head
	newP := &newHead
	for p != nil {
		if p.Random != nil {
			newP.Random = m[p.Random]
		}
		p = p.Next
		newP = newP.Next
	}
	return &newHead
}

//方法三，新节点插入原对应节点后面，然后差分新节点
func CopyRandomList3(head *Node) *Node {
	if head == nil {
		return nil
	}
	p := head
	for p != nil {
		newNode := &Node{
			Val:    p.Val,
			Next:   p.Next,
			Random: nil,
		}
		p.Next = newNode
		p = p.Next.Next
	}
	p = head
	for p.Next != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}
	newHead := head.Next
	oldp := head
	p = newHead
	for p.Next != nil {
		oldp.Next = oldp.Next.Next
		p.Next = p.Next.Next
		oldp = oldp.Next
		p = p.Next
	}
	oldp.Next = nil
	return newHead
}
