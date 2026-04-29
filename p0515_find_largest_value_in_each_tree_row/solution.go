package p0515

// TreeNode 由 LeetCode 平台预定义；本地补一份，便于通过编译。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelMax := queue[0].Val // 用本层第一个节点的 Val 做初值，免去哨兵
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Val > levelMax {
				levelMax = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, levelMax)
	}
	return result
}
