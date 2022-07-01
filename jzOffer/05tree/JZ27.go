package tree

//二叉树的镜像
//方法一，递归，后续遍历
func MirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	MirrorTree(root.Left)
	MirrorTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

//方法二 非递归，层次遍历
func MirrorTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return root
}
