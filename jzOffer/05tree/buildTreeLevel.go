package tree

//根据层级遍历构建一颗二叉树
func BuildTreeLevel(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	i := 1
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if i < len(nums) {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(nums) {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}
