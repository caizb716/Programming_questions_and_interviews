package tree

import (
	"fmt"
	"strconv"
	"strings"
)

//层次遍历序列化
func Serialize(root *TreeNode) string {
	ans := "["
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]
		if temp == nil {
			ans += "null,"
			continue
		} else {
			ans += fmt.Sprintf("%d,", temp.Val)
		}
		queue = append(queue, temp.Left)
		queue = append(queue, temp.Right)
	}
	ans = ans[:len(ans)-1] + "]"
	return ans
}
func Deserizlize(data string) *TreeNode {
	data = data[1 : len(data)-1]
	dataSplit := strings.Split(data, ",")
	if len(dataSplit) == 1 {
		return nil
	}
	nodeQueue := make([]*TreeNode, 0)
	dataSplitCount := 0
	rootVal, _ := strconv.Atoi(dataSplit[dataSplitCount])
	root := &TreeNode{Val: rootVal}
	nodeQueue = append(nodeQueue, root)
	dataSplitCount++
	for len(nodeQueue) > 0 {
		temp := nodeQueue[0]
		nodeQueue = nodeQueue[1:]
		LeftNodeContext := dataSplit[dataSplitCount]
		dataSplitCount++
		if LeftNodeContext != "null" {
			LeftNodeVal, _ := strconv.Atoi(LeftNodeContext)
			temp.Left = &TreeNode{Val: LeftNodeVal}
			nodeQueue = append(nodeQueue, temp.Left)
		} else {
			temp.Left = nil
		}
		RightNodeContext := dataSplit[dataSplitCount]
		dataSplitCount++
		if RightNodeContext != "null" {
			RightNodeVal, _ := strconv.Atoi(RightNodeContext)
			temp.Right = &TreeNode{Val: RightNodeVal}
			nodeQueue = append(nodeQueue, temp.Right)
		} else {
			temp.Right = nil
		}
	}
	return root
}
