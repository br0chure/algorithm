package p0124

import "math"

// TreeNode 由 LeetCode 平台预定义；本地补一份，便于通过编译。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxPathSum 主解：闭包 + 后序聚合（B 型 · 单边贡献 vs 全局答案分离）。
//
// 三问回答：
//   - 返回值：int，「从当前节点向下延伸的**单边**最长路径和」（给父节点拼接用的零件）
//   - 空节点：返回 0（约定空树贡献 0；配合 max(0, ...) 让"不接这条腿"成为合法选择）
//   - 当前节点：
//     * 全局：sum = leftMax + rightMax + root.Val（经过当前节点的双边路径）→ 更新 maxSum
//     * 返回：max(leftMax, rightMax) + root.Val（单边贡献，给父节点）
//
// 关键 max(0, ...) 在哪能用、哪不能用：
//   - max(0, path(root.Left/Right))  ✅ 子树贡献为负就「砍掉」（不接这条腿）
//   - 但 root.Val 必须无条件加进 sum——路径经过 root，root 不能扔
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt
	var path func(*TreeNode) int
	path = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftMax := max(0, path(root.Left))
		rightMax := max(0, path(root.Right))
		sum := leftMax + rightMax + root.Val
		if maxSum < sum {
			maxSum = sum
		}
		return max(leftMax, rightMax) + root.Val
	}
	path(root)
	return maxSum
}
