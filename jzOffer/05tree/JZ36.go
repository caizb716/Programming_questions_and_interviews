package tree

//二叉搜索树与双向链表
func TreeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var pre *TreeNode
	helper(root, &pre) //使用多重指针即可
	head, tail := root, root
	for head.Left != nil {
		head = head.Left
	}
	for tail.Right != nil {
		tail = tail.Right
	}
	head.Left = tail
	tail.Right = head
	return head
}

func helper(root *TreeNode, pre **TreeNode) {
	if root == nil {
		return
	}
	helper(root.Left, pre)
	if (*pre) != nil {
		root.Left = (*pre)
		(*pre).Right = root
	}
	(*pre) = root
	helper(root.Right, pre)
}
