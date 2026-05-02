package p0543

// TreeNode 由 LeetCode 平台预定义；本地补一份，便于通过编译。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// diameterOfBinaryTree 主解：闭包 + 后序聚合。
//
// 核心思路（详见 patterns/tree_dfs.md 心法 #2「返回值 vs 全局答案分离」）：
//   - 函数 depth 返回的是「子树向下到最深叶子的节点数（含自己）」——给父节点拼接用的「单边零件」。
//   - 全局变量 maxDiameter 记录「经过当前节点的最长路径边数」的最大值——这才是题目要的答案。
//   - 关键恒等式：路径边数 = 路径节点数 - 1 = depth(left) + depth(right)。
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := depth(node.Left)
		right := depth(node.Right)
		maxDiameter = max(maxDiameter, left+right) // 经过当前节点的路径边数
		return max(left, right) + 1                // 向父节点贡献的单边深度（节点数）
	}
	depth(root)
	return maxDiameter
}
