package p0116

// connectV2 是首次 AC 版本，保留作对比。
// 思路上做了「重复抽象」：拉一对 (left, right) 到外层维护，并显式复位 right = nil。
// 实际上 BFS 的 queue + levelSize 钩子已经管好了「同层」语义，
// 不需要再造一层指针对——简化成主解 connect 即可。
//
// 反思流程见 patterns/tree_bfs.md 心法 #8（AC 后两问）：
//   - var left/right 跨迭代没用到 → 应挪进循环
//   - right 从写到读只 1 次 → 应直接内联 node.Next = queue[0]
func connectV2(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		var left *Node
		var right *Node
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			left = queue[0]
			queue = queue[1:]

			if levelSize-i > 1 { // 本层除了 left 还有节点
				right = queue[0]
			}
			if levelSize-i == 1 { // 本层只剩 left 一个节点
				right = nil
			}

			if left.Left != nil {
				queue = append(queue, left.Left)
			}
			if left.Right != nil {
				queue = append(queue, left.Right)
			}
			left.Next = right
		}
	}
	return root
}
