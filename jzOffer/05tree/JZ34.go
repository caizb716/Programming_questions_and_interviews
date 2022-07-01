package tree

//二叉树中和为某一值的路径
func PathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	dfs(root, sum, []int{}, &res)
	return res
}
func dfs(root *TreeNode, sum int, path []int, res *[][]int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	if root.Val == sum && root.Left == nil && root.Right == nil {
		//防止底层数组被改变
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}
	dfs(root.Left, sum-root.Val, path, res)
	dfs(root.Right, sum-root.Val, path, res)
}
