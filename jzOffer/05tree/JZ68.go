package tree

//68.1 二叉搜索树的最近公共祖先

func getPath(root, target *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	node := root
	var path []*TreeNode

	for node != target {
		path = append(path, node)
		if target.Val > node.Val {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	path = append(path, node)
	return path

}
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	pathP := getPath(root, p)
	pathQ := getPath(root, q)
	var res *TreeNode
	for i := 0; i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i]; i++ {
		res = pathP[i]
	}
	return res
}

//68.2 普通二叉树的最近公共祖先
func getPath2(root, target *TreeNode, path []*TreeNode, res *[]*TreeNode) {
	if root == nil || *res != nil {
		return
	}
	path = append(path, root)
	if root == target {
		temp := make([]*TreeNode, len(path))
		copy(temp, path)
		*res = append(*res, temp...)
	}
	getPath2(root.Left, target, path, res)
	getPath2(root.Right, target, path, res)

}
func LowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var path []*TreeNode
	var pathP []*TreeNode
	var pathQ []*TreeNode
	getPath2(root, p, path, &pathP)
	getPath2(root, q, path, &pathQ)
	var res *TreeNode
	for i := 0; i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i]; i++ {
		res = pathP[i]
	}
	return res
}
