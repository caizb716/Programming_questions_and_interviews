package tree

//对称的二叉树
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}
func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}
