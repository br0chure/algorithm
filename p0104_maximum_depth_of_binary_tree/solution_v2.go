package p0104

// maxDepthBFS 备用解法：BFS 数层数。
// 思路：标准 BFS 模板，每进入一层 depth++。
// 时间 O(n)；空间 O(w)，w = 树最大宽度（最坏满二叉树最后一层 ≈ n/2 → O(n)）。
//
// 与主解 maxDepth（DFS 后序）的对比见 README「BFS vs DFS 空间权衡」。
func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 0
	for len(queue) > 0 {
		depth++
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return depth
}
