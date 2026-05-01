package p0116

// Node 由 LeetCode 平台预定义；本地补一份，便于通过编译。
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// connect 主解：标准 BFS 模板 + 1 行节点级动作。
// 关键 insight：弹出当前节点后，queue[0] 就是「同层的下一个节点」。
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if i < levelSize-1 {
				node.Next = queue[0]
			}
			// 最后一个节点 Next 保持 Go 零值 nil，无需显式赋值。
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}
