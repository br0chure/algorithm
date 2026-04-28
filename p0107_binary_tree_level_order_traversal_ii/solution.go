package p0107

// TreeNode 由 LeetCode 平台预定义；本地补一份，便于通过编译。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	var level []int

	for len(queue) > 0 {
		levelSize := len(queue)
		level = make([]int, 0, levelSize) //每层需要一个空切片
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		temp := result[i]
		result[i] = result[j]
		result[j] = temp
	}
	return result
}
