package p0102

// TreeNode 由 LeetCode 平台预定义；本地补一份，便于通过编译。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*TreeNode{root}
	//遍历树
	for len(queue) > 0 {
		size := len(queue)
		var level = make([]int, 0, size)
		//遍历层
		for i := 0; i < size; i++ {
			//当前层的数据出队
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			//下一层的数据入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}
