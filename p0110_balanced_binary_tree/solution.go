package p0110

// TreeNode 由 LeetCode 平台预定义；本地补一份，便于通过编译。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isBalanced 主解：后序 DFS + -1 哨兵。
//
// 核心思路（详见 patterns/tree_dfs.md 心法 #4「用哨兵值携带失败标记」）：
//   - height 函数返回当前子树的高度（节点数），高度 ≥ 0
//   - 一旦发现「不平衡」就返回 -1，借助高度永非负的特性，让 -1 不会和真实高度撞车
//   - 父节点拿到孩子返回值时先检查 -1，是的话立刻短路，无需计算另一侧子树
func isBalanced(root *TreeNode) bool {
	var height func(*TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := height(node.Left)
		if left == -1 {
			return -1
		}
		right := height(node.Right)
		if right == -1 {
			return -1
		}
		if left-right > 1 || right-left > 1 {
			return -1
		}
		return max(left, right) + 1
	}
	return height(root) != -1
}
