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

### 3. 临时 slice 的初始化位置 = 语义

```go
var level []int                                  // A. 永不重置
for len(queue) > 0 {
    level = make([]int, 0, levelSize)            // B. 每层重置 ← 正确写法
    for i := 0; i < levelSize; i++ {
        // level := make([]int, 0)               // C. 每节点重置
        ...
    }
    result = append(result, level)
}
```

| 位置 | 重置频率 | 语义 |
|---|---|---|
| A. 外层 for **之外** | 永不重置 | 所有层共享同一底层数组 → 结果错乱 |
| B. 外层 for **里**、内层 for **前** | 每层一次 | **正确**，每层独立切片 |
| C. 内层 for **里** | 每节点一次 | 每层只保留最后一个 Val → 等价于「右视图」(199) |

> **记忆口诀**：初始化位置 = 重置频率。要"每层一份"就放每层那个循环的开头。把它挪进内层循环 = 你想要每层最后一个元素 = 右视图。

### 4. 聚合型问题：在标量上滚动维护，不要先收集再处理

求「每层最大 / 最小 / 求和 / 平均」时，**不需要**先把整层 Val 收成切片再算。
直接维护一个标量（`levelMax / levelSum / ...`）随节点滚动更新，层尾把它追加进 `result` 即可。

```go
for len(queue) > 0 {
    levelSize := len(queue)
    levelMax := queue[0].Val          // ← 每层重置；用「集合中已有元素」做初值，免哨兵
    for i := 0; i < levelSize; i++ {
        node := queue[0]; queue = queue[1:]
        if node.Val > levelMax { levelMax = node.Val }
        // ... 入队左右孩子
    }
    result = append(result, levelMax)
}
```

**关键自问：「这个变量属于哪一层的状态？」** 答案是「层」→ 在外层 for 开头、内层 for 之前初始化（位置 B）。
心法 #3 的「初始化位置 = 重置频率」对**切片和标量同样适用**——不要把它当成只是切片的规律。

> **初值技巧**：能用「集合中已有元素」就别用哨兵。`queue[0].Val` 比 `math.MinInt` 更稳——节点值范围再变也不出错（注：`math.MinInt` 在 Go 里**是负数**，等于 `-2^63`，可以做哨兵；但能避免就避免）。

### 5. 题面方向 ≠ 解法方向

题目说"自底向上"不代表你必须**自底向上 BFS**。BFS 的天然方向是 FIFO + 从根开始，强行反向更难。**用熟模板 + 后处理**（如 reverse）几乎永远比改造模板简单。这条对所有模板题都通用。

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
- [0107 二叉树的层序遍历 II](../p0107_binary_tree_level_order_traversal_ii/) 🟡
- [0515 在每个树行中找最大值](../p0515_find_largest_value_in_each_tree_row/) 🟡
