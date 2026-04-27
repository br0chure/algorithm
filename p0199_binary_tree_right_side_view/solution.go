package p0199

// TreeNode 由 LeetCode 平台预定义；本地补一份，便于通过编译。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var result = make([]int, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	var rightVal int
	for len(queue) > 0 {
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
			if i == levelSize-1 {
				rightVal = node.Val
			}
		}
		result = append(result, rightVal)
	}
	return result
}
