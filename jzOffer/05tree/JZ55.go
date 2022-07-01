package tree

import "math"

//55.1 二叉树的深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//55.2 判断平衡二叉树
func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if math.Abs(float64(left-right)) > 1 {
		return false
	}
	return IsBalanced(root.Left) && IsBalanced(root.Right)
}
