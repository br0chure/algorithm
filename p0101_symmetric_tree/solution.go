package p0101

// TreeNode 由 LeetCode 平台预定义；本地补一份，便于通过编译。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSymmetric 双树同步递归（详见 patterns/tree_subtree.md）。
//
// 核心：把 root 的左右子树看成两棵独立的树，递归判断它们是否互为镜像。
// 镜像配对：左树的左 vs 右树的右、左树的右 vs 右树的左。
//
// 用闭包是因为递归必须带两个 node 参数（dfs(p, q)），但 LC 入口签名只有一个 root，
// 闭包内嵌一个 dfs(p, q) 桥接两者。
func isSymmetric(root *TreeNode) bool {
	var dfs func(p, q *TreeNode) bool
	dfs = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		return dfs(p.Left, q.Right) && dfs(p.Right, q.Left)
	}
	return dfs(root.Left, root.Right)
}
