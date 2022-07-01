package tree

//二叉树的下一个节点
func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	p := pNode
	if p.Right != nil {
		p = p.Right
		for p.Left != nil {
			p = p.Left
		}
		return p
	} else {
		for p.Next != nil && p.Next.Right == p {
			p = p.Next
		}
		return p.Next
	}
}
