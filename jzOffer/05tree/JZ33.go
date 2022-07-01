package tree

//判断一个数组是不是某个二叉搜索树的后序遍历
func VerifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	f := len(postorder) - 1
	for i := 0; i < len(postorder); i++ {
		if postorder[i] >= postorder[len(postorder)-1] {
			f = i
			for ; i < len(postorder); i++ {
				if postorder[i] < postorder[len(postorder)-1] {
					return false
				}
			}
			break
		}

	}
	return VerifyPostorder(postorder[:f]) && VerifyPostorder(postorder[f:len(postorder)-1])
}
