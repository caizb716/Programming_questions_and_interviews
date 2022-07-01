package tree

//树的子结构
func IsSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}
	return helper1(A, B) || IsSubStructure(A.Left, B) || IsSubStructure(A.Right, B)

}
func helper1(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return helper1(A.Left, B.Left) && helper1(A.Right, B.Right)
}
