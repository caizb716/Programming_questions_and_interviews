package tree

//二叉查找树的第K大节点
func KthNode(root *TreeNode, k int) *TreeNode {
	count := 0
	var p *TreeNode
	getNum(root, &count, &p, k)
	return p

}
func getNum(root *TreeNode, count *int, p **TreeNode, k int) {
	if root == nil {
		return
	}
	getNum(root.Right, count, p, k)
	*count++
	if *count == k {
		*p = root
	}
	getNum(root.Left, count, p, k)
}
