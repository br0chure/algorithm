package p0104

// TreeNode 由 LeetCode 平台预定义；本地补一份，便于通过编译。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth 后序 DFS：先算左右子树的最大深度，再 +1 当前节点。
// 三问回答：
//   - 返回值：int，含义是「以当前节点为根的子树最大深度（节点数）」
//   - 空节点（base case）：返回 0
//   - 合并：max(left, right) + 1
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return max(leftDepth, rightDepth) + 1
}
