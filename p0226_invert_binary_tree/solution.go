package p0226

// TreeNode 由 LeetCode 平台预定义；本地补一份，便于通过编译。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree 主解：后序 DFS，原地翻转。
//
// 用 void helper 表达「这是 in-place 修改」语义——LC 签名要求外层返回 *TreeNode，
// 但实际所有翻转都在原树上完成，return root 只是把同一个指针传回去。
func invertTree(root *TreeNode) *TreeNode {
	invert(root)
	return root
}

func invert(root *TreeNode) {
	if root == nil {
		return
	}
	invert(root.Left)
	invert(root.Right)
	root.Left, root.Right = root.Right, root.Left
}
