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

> **初值技巧**：能用「集合中已有元素」就别用哨兵。`queue[0].Val` 比 `math.MinInt` 更稳——节点值范围再变也不出错（注：`math.MinInt` 在 Go 里**是负数**，等于 `-2^63`，可以做哨兵；但能避免就避免）。

### 5. 题面方向 ≠ 解法方向

题目说"自底向上"不代表你必须**自底向上 BFS**。BFS 的天然方向是 FIFO + 从根开始，强行反向更难。**用熟模板 + 后处理**（如 reverse）几乎永远比改造模板简单。这条对所有模板题都通用。

### 6. 输出大小已知 → 预分配定长切片，按 index 写入

BFS 一进入新一层就有 `levelSize := len(queue)`——本层输出的**大小已经确定**。这种情况下两种切片用法的差别变得关键：

```go
level := make([]int, 0, levelSize)   // 长度 0、容量 levelSize ← 只能 append
level := make([]int, levelSize)      // 长度就是 levelSize、全是零值 ← 可以随机写 level[i] = x
```

第二种用法的好处是**随机写**——你可以决定第 i 个节点的 Val 写到 `level` 的哪个位置，不必非得追加到尾。这对「事后想再调整顺序」的题（zigzag、按某种规则重排）有奇效：

```go
// 103 锯齿形：偶数层倒着写，省掉事后 reverse
level := make([]int, levelSize)
for i := 0; i < levelSize; i++ {
    node := queue[0]; queue = queue[1:]
    pos := i
    if reverseLevel {
        pos = levelSize - 1 - i
    }
    level[pos] = node.Val
    // ... 入队孩子
}
```

**为什么不用「append 到队头」实现倒序？** 头插每次都要把已有元素整体后移，本层 k 个节点 → O(k²)，**比事后 reverse（O(k)）更慢**，反而是负优化（同样的坑见 107 「为什么不在头部 insert」）。

> **通用思维**：
>
> - **输出大小已知** → `make([]T, k)` 预分配定长 + 按 index 写
> - **输出大小未知** → `make([]T, 0[, cap])` + append + 必要时后处理（reverse / sort / ...）
>
> 这个区分不只 BFS 适用，写任何 Go 代码遇到「我要构造一个切片」时都该先问一句「大小知不知道」。

### 7. 「同层下一个节点」 = 弹出当前后的 `queue[0]`

凡是需要「同层相邻节点」的题（116/117 填 Next 指针、层内对比相邻节点等），不要再造一对持久指针 `(left, right)` 跨迭代维护——**信息已经在 queue 里了**：

```go
for i := 0; i < levelSize; i++ {
    node := queue[0]
    queue = queue[1:]
    if i < levelSize - 1 {
        node.Next = queue[0]   // ← 此刻 queue[0] 就是同层下一个节点
    }
    // ... 入队孩子（追加到尾，不影响 queue[0]）
}
```

**前提**：`levelSize` 钩子保证当前层节点都还在 queue 头部连续排列；孩子追加到尾部，不会污染头部。

**反例（过度设计）**：拉一对 `var left, right *Node` 到外层、用 `levelSize - i > 1` 判断本层剩余、还要显式 `right = nil` 复位——这是把「已经在 queue 里的信息」复制到外层变量再用一次，**重复抽象**。BFS 的 `queue + levelSize` 钩子已经替你管好了「同层」语义，不要再造一层。

### 8. AC 后的「两问反思」（删冗余的固定流程）

代码 AC 不代表写得好。每次 AC 后强制对每个变量问两遍：

> **Q1：这个变量跨迭代用过吗？**
> 答「没有」（每轮一开始就被覆盖）→ **挪进循环里声明**。
>
> **Q2：它从赋值到使用之间被读了几次？**
> 答「只 1 次」（写完立刻读、中间没别的事）→ **删掉变量，直接把右值贴到使用处**。

这两问是机械操作、不需要灵感。跑完一遍代码会自动从「能跑」收敛到「最简」。
配合「先写朴素版 AC，再跑两问简化」的节奏——比想清楚再动手快得多。

> **典型案例（116）**：
>
> - `var left *Node` → Q1 没跨迭代 → 挪进循环 → 改 `left := queue[0]`
> - `var right *Node` → Q1 没跨迭代 → 挪进循环 → 顺带删掉 `right = nil` 复位
> - `right` → Q2 写完立刻读 1 次 → 删变量 → 直接 `node.Next = queue[0]`

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

- [0102 二叉树的层序遍历](../../problems/二叉树/0102%20二叉树的层序遍历.md) 🟡
- [0199 二叉树的右视图](../../problems/二叉树/0199%20二叉树的右视图.md) 🟡
- [0107 二叉树的层序遍历 II](../../problems/二叉树/0107%20二叉树的层序遍历%20II.md) 🟡
- [0515 在每个树行中找最大值](../../problems/二叉树/0515%20在每个树行中找最大值.md) 🟡
- [0103 二叉树的锯齿形层序遍历](../../problems/二叉树/0103%20二叉树的锯齿形层序遍历.md) 🟡
- [0116 填充每个节点的下一个右侧节点指针](../../problems/二叉树/0116%20填充每个节点的下一个右侧节点指针.md) 🟡

## 🎓 BFS 毕业（按 CLAUDE.md「Pattern 推进策略」）

5 种「质」覆盖完成：

| 质 | 题 |
|---|---|
| 收集全部 | 0102 |
| 选择性收集 | 0199 |
| 聚合（标量滚动） | 0515 |
| 顺序变换（reverse / 预分配） | 0107 / 0103 |
| 横向连接（同层相邻） | 0116 |

模板默写无压力 → **BFS 毕业**，下一步 `tree_dfs`。
（637 层平均值与 515 同质，不必再刷；117 与 116 思路一致，仅常数空间约束略有不同，可后续按需补。）
