package tree

//按照层级打印二叉树

//从上到下打印二叉树————一行
func LevelOrderOneLine(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}

//从上到下打印二叉树————多行
func LevelOrderMulitLine(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	level := 0
	for len(queue) != 0 {
		temp := []*TreeNode{}
		res = append(res, []int{})
		for _, v := range queue {
			res[level] = append(res[level], v.Val)
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
		}
		level++
		queue = temp
	}
	return res
}

//从上到下打印二叉树————按之字打印
func LevelOrderZ(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res = [][]int{}
	var queue = []*TreeNode{root}
	var level, counter, knife, cur int
	for len(queue) > 0 {
		counter = len(queue)
		knife = counter
		res = append(res, []int{})
		for 0 < counter {
			cur = knife - counter
			if queue[cur].Left != nil {
				queue = append(queue, queue[cur].Left)
			}
			if queue[cur].Right != nil {
				queue = append(queue, queue[cur].Right)
			}

			counter--
			if level%2 != 0 {
				cur = counter
			}
			res[level] = append(res[level], queue[cur].Val)
		}
		queue = queue[knife:]
		level++
	}
	return res
}
