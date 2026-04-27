# Tree BFS（树的广度优先遍历）

## 何时使用

- 需要按层处理（每层独立结果、第 K 层节点等）
- 求最短路径 / 最少步数（无权图、树）
- 水平视图（左视图、右视图）
- 层间统计（最值、求和、平均）

## Go 通用模板

```go
func bfs(root *TreeNode) {
    if root == nil {
        return
    }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        levelSize := len(queue)        // ⚠️ 锁定本层节点数
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            // —— 节点级动作：处理 node ——
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        // —— 层级动作：本层结束收尾 ——
    }
}
```

## 心法

### 1. 节点级 vs 层级动作要分离

- **节点级**（每个 node 都做）：弹出 + 把左右孩子入队。**和"是否本层最后一个"无关**。
- **层级**（每层做一次）：开 level 切片、收 rightVal、累加 sum 等。

混淆这两类是新手最常见 bug。把"塞孩子"放进 `if i == levelSize-1` 分支里，会死循环或漏节点。

### 2. `levelSize := len(queue)` 是分层的唯一钩子

循环一开始就冻结这个数，下面 for 循环只处理本层；新入队的孩子自动归到"下一层"。整个 BFS 模板就靠这一根线把"层"和"节点"分开。

## 常见变体（在通用模板上的改动）

| 题号 | 名称 | 改动 |
|---|---|---|
| 102 | 二叉树的层序遍历 | 收每层全部 Val |
| 199 | 二叉树的右视图 | `i == levelSize-1` 时收 |
| 107 | 层序遍历 II | 同 102，最后 reverse |
| 515 | 每行最大值 | 每层维护 max |
| 637 | 层平均值 | 每层 sum / size |
| 103 | 锯齿形层序遍历 | 奇偶层方向交替 |
| 116 / 117 | 填充右侧节点指针 | BFS + 层内连接 |

## 已刷题目

- [0102 二叉树的层序遍历](../p0102_binary_tree_level_order_traversal/) 🟡
- [0199 二叉树的右视图](../p0199_binary_tree_right_side_view/) 🟡
