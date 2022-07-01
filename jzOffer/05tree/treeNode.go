package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type TreeLinkNode struct {
	Val   int
	Left  *TreeLinkNode
	Right *TreeLinkNode
	Next  *TreeLinkNode //指向父节点的指针
}
