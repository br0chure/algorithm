# Tree DFS（树的深度优先遍历）

## 何时使用

- 需要沿路径深入处理（路径和、最长路径、根到叶的某属性等）
- 需要左右子树**先算好**才能算当前节点（树的高度、直径、是否平衡等聚合型）
- 需要按某种顺序访问节点（中序遍历 BST 得到有序序列）
- 需要在递归中携带"上下文"（祖先信息、深度、累计值等）

## Go 通用模板：三种顺序

DFS 在树上落地为递归，按"处理当前节点 vs 处理子树"的相对顺序分三种：

### 前序（Preorder）：根 → 左 → 右

适合「自顶向下」传递信息，或在子树展开前先做记录。

```go
func preorder(node *TreeNode) {
    if node == nil {
        return
    }
    // —— 处理 node（在递归子树之前）——
    preorder(node.Left)
    preorder(node.Right)
}
```

### 中序（Inorder）：左 → 根 → 右

BST 中序 = 升序序列；二叉树中序 = 左到右的「投影」。

```go
func inorder(node *TreeNode) {
    if node == nil {
        return
    }
    inorder(node.Left)
    // —— 处理 node（左子树访问完之后）——
    inorder(node.Right)
}
```

### 后序（Postorder）：左 → 右 → 根 ⭐ 最常用

**先把左右子树的结果算出来**，再合并成当前节点的结果。聚合型 DFS（树高、直径、平衡判断、最大路径和……）几乎都是后序。

```go
func postorder(node *TreeNode) ResultT {
    if node == nil {
        return zeroValue   // 空树的边界值，按题目语义定
    }
    left := postorder(node.Left)
    right := postorder(node.Right)
    // —— 合并 left、right 与 node 自身，算出当前节点的 ResultT ——
    return merged
}
```

## 三个核心问题（写任何 DFS 前先问）

每写一个 DFS 函数，先在脑子里把这三个填好：

1. **递归返回什么？**（返回值类型 + 含义。如「以当前节点为根的子树最大深度」）
2. **空节点（base case）返回什么？**（直接决定边界正确性。最大深度 → 0；存在性 → false；最大路径和 → 看题目允许不允许走空路径）
3. **当前节点用左右子树的返回值做什么？**（这就是后序模板里"合并"那一行）

写不出来这三条就别动手——等想清楚再写。这比 BFS 的"两个动作分离"更需要预先设计。

## 心法

### 1. 递归的空间复杂度 = 栈深 × 单帧大小

每次函数调用，运行时在「调用栈」上分配一块**栈帧**，装：

- 参数（传进来的）
- 局部变量（函数体里 `:=` / `var` 声明的）
- 返回地址 + 寄存器保存等 metadata

函数 `return` 时这块栈帧**立刻释放**。所以：

> **递归总空间 = 同时活着的栈帧数 × 单帧大小 = 栈深 × 单帧大小**

#### 单帧是 O(1) 还是 O(k)？看本地分配

| 函数本地内容 | 单帧大小 |
|---|---|
| 几个标量 / 指针 / 固定大小数组 | **O(1)** |
| `make([]T, k)`、`var [k]T`、字符串拼接产生大对象 | O(k) |

⚠️ **返回值类型不影响单帧大小**——返回值走寄存器或调用者预留位，不占被调用者栈帧。返回 `int` 还是 `*TreeNode` 还是 `bool` 都一样。决定单帧的是**函数体里声明了多少东西**。

#### 栈深 = 当前路径长度，最大 = 树高 h

- 平衡树：h = O(log n)
- 链状树：h = O(n)
- 一般写 O(h)，最坏 O(n)

#### 写 DFS 前心里过一遍三步

1. **栈深多深？** → 通常是 h（树高）
2. **每一帧本地分配多少？** → 看函数体（一般是 O(1)）
3. **总空间 = 栈深 × 单帧** → 一般 DFS 是 **O(h) × O(1) = O(h)**

#### 反例：单帧带 O(k) 分配的递归

```go
func badDfs(root *TreeNode) int {
    if root == nil { return 0 }
    arr := make([]int, 10000)        // ← 每帧多 8 万字节
    _ = arr
    return max(badDfs(root.Left), badDfs(root.Right)) + 1
}
```

栈深 100 → 总空间 ~8MB（不是想象中的 100 个 int）。看到本地大型分配要警觉。

---

### 2. 函数返回值与最终答案的三种关系

DFS 题最重要的**设计决定**：函数返回的东西和题目要的答案是什么关系？想清楚这个，结构就有了。

| 结构 | 函数返回 | 最终答案在哪 | 何时是这种 |
|---|---|---|---|
| **A. 重合** | 答案本身 | `dfs(root)` 直接就是 | 答案是整棵树的某个聚合数值（深度、节点数、和） |
| **B. 分离** | 给父节点拼接的「单边零件」 | 闭包变量在递归过程中累计 | 答案藏在「子树之间」或「不必从根开始」 |
| **C. 哨兵** | 合法值 / 哨兵值（如 `-1`） | `dfs(root) != 哨兵` | 答案是 bool，且单点违反即整树违反 |

#### 识别信号（按题面措辞）

| 题面关键词 | 大概率是 |
|---|---|
| 「整棵树的 X」「树的 X」（X 是数值） | A |
| 「任意两点」「不必经过根」「不必到叶子」「子树间最优」 | B |
| 「是否平衡 / 对称 / 是 BST」「每个节点都满足 …」 | C |

#### 写法骨架

```go
// A：返回值即答案
func solve(root *TreeNode) int {
    if root == nil { return 0 }
    return combine(solve(root.Left), solve(root.Right), root)
}

// B：闭包变量记答案，返回值是单边零件
func solve(root *TreeNode) int {
    answer := initialValue
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil { return 0 }
        l, r := dfs(node.Left), dfs(node.Right)
        answer = update(answer, mergeAcross(l, r, node))   // 经过当前节点的最优
        return contributeUp(l, r, node)                    // 单边贡献给父节点
    }
    dfs(root)
    return answer
}

// C：哨兵 + 短路
func solve(root *TreeNode) bool {
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil { return baseValue }
        l := dfs(node.Left);  if l == sentinel { return sentinel }
        r := dfs(node.Right); if r == sentinel { return sentinel }
        if !valid(l, r, node) { return sentinel }
        return aggregate(l, r, node)
    }
    return dfs(root) != sentinel
}
```

#### B 型的关键约束：「跨节点合并」 ≠ 「向上贡献」

`mergeAcross`（左右两边拼起来经过当前节点）和 `contributeUp`（向上贡献给父节点）是**两个不同的合并方式**——绝不能写成同一个表达式。

原因：路径是**简单路径**，不能 Y 形分叉上去。所以拼成「经过当前节点的最优」后这条路径就**不能再延伸**，必须把它存到 `answer` 里；而向父节点贡献的，只能选 left/right **一边**接上 node。

#### C 型的哨兵设计：两个条件

1. **哨兵值不撞车合法返回值**——返回值的合法范围里没有这个值。`-1` 适合「高度永非负」类；不适合「路径和可负」类（用 `MinInt` 或封装在 struct 里）。
2. **一旦发现就立刻 return**——哨兵的全部价值在短路；拿到孩子返回值的下一行就检查，是哨兵就 return，跳过另一侧子树。

### 3. 闭包是 B 型 / C 型递归的 Go 写法首选

B 型需要在递归中维护一个 `answer`，C 型偶尔也用辅助变量。这个变量怎么放？

| 写法 | 评价 |
|---|---|
| 包级 `var` | ❌ 全局污染、多次调用需重置、并发不安全 |
| **函数内局部变量 + 闭包捕获** | ⭐ 推荐，干净 |
| 指针参数 `*int` | 能用，但每层多传一个参数 |

#### 自递归闭包的两步声明套路

```go
var dfs func(*TreeNode) int          // ① 声明：让 dfs 进入作用域
dfs = func(node *TreeNode) int {     // ② 赋值：函数体可以引用 dfs(...)
    ...
    dfs(node.Left)
    dfs(node.Right)
    ...
}
dfs(root)                            // ③ 调用
```

**为什么不能 `dfs := func(...)` 一行？**
Go `:=` 是「先求右值再绑定左值」——执行右边的 `func` 字面量时，左边的 `dfs` 还没在作用域里，函数体里 `dfs(...)` 找不到标识符。两步声明把「让名字进入作用域」和「赋值」拆开。

非自递归闭包不存在这个问题，可以一行：

```go
double := func(x int) int { return x * 2 }
```

#### 闭包 ≠ goroutine

- **闭包** = 函数字面量 + 捕获外层变量（**同步**执行）
- **`go func(){...}()`** = `go` 关键字启动 goroutine（**异步并发**）

刷题里用的全是闭包，没有并发。

---

## 常见变体

| 题号 | 名称 | 关键动作 |
|---|---|---|
| 104 | 二叉树的最大深度 | 后序聚合：max(left, right) + 1 |
| 226 | 翻转二叉树 | 前 / 后序皆可，左右指针交换 |
| 100 | 相同的树 | 同步 DFS 两棵树 |
| 110 | 平衡二叉树 | 后序 + 剪枝（用「-1」表示已不平衡） |
| 543 | 二叉树的直径 | 后序聚合 + 用全局变量记录最大值 |
| 124 | 二叉树中的最大路径和 | 后序 + 「向上贡献」与「全局答案」分离 |
| 236 | 最近公共祖先 | 后序 + 「子树是否包含目标」 |
| 111 | 二叉树的最小深度 | 后序，但要小心「单边为 nil」的边界 |

## 已刷题目

- [0104 二叉树的最大深度](../0104%20二叉树的最大深度.md) 🟢
- [0543 二叉树的直径](../0543%20二叉树的直径.md) 🟢
- [0110 平衡二叉树](../0110%20平衡二叉树.md) 🟢
- [0111 二叉树的最小深度](../0111%20二叉树的最小深度.md) 🟢
- [0124 二叉树中的最大路径和](../0124%20二叉树中的最大路径和.md) 🔴
- [0226 翻转二叉树](../0226%20翻转二叉树.md) 🟢
- [0236 二叉树的最近公共祖先](../0236%20二叉树的最近公共祖先.md) 🟡（Type B 信号 + 全局答案；附压缩版备用）
