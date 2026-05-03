package p0111

// TreeNode 由 LeetCode 平台预定义；本地补一份，便于通过编译。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// minDepth 后序 DFS（Type A 含边界）。
//
// 三问回答：
//   - 返回值：int，「以当前节点为根的子树最小深度（节点数）」
//   - 空节点：返回 0
//   - 合并：如果某一侧子树为 nil（其返回值是 0），则不能用 0 参与 min——
//     因为 nil 不是「叶子」，那一侧没有合法的"到叶子路径"。
//     必须用非 nil 那一侧的深度 + 1。
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if left == 0 && right != 0 {
		return right + 1
	}
	if left != 0 && right == 0 {
		return left + 1
	}
	if left == 0 && right == 0 {
		return 1
	}
	return min(left, right) + 1
}
