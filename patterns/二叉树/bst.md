# BST（二叉搜索树）

## BST 的两条性质

- **左小右大**：节点的左子树所有值 < 节点值；右子树所有值 > 节点值（递归适用于每个子树）
- **中序有序**：中序遍历 BST 得到**严格升序**序列

> 大多数 BST 题的钥匙都是这两条之一。**搜索/插入用左小右大；验证/排名用中序有序**。

## 模式 1：搜索 / 插入（用左小右大递归）

```go
// 搜索（700）
func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil { return nil }
    if val < root.Val { return searchBST(root.Left, val) }
    if val > root.Val { return searchBST(root.Right, val) }
    return root
}

// 插入（701）：递归到 nil 处创建新节点
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil { return &TreeNode{Val: val} }
    if val < root.Val {
        root.Left = insertIntoBST(root.Left, val)
    } else {
        root.Right = insertIntoBST(root.Right, val)
    }
    return root
}
```

时间复杂度：平衡 BST 是 O(log n)，最坏链状是 O(n)。

## 模式 2：验证 / 排名（用中序遍历）

### 验证 BST（98）：中序应严格递增

```go
func isValidBST(root *TreeNode) bool {
    var prev *TreeNode    // 中序上一个节点
    var dfs func(*TreeNode) bool
    dfs = func(node *TreeNode) bool {
        if node == nil { return true }
        if !dfs(node.Left) { return false }       // 左
        if prev != nil && node.Val <= prev.Val {  // 中：与前一个节点比
            return false
        }
        prev = node
        return dfs(node.Right)                    // 右
    }
    return dfs(root)
}
```

### 第 k 小（230）：中序到第 k 个停止

```go
func kthSmallest(root *TreeNode, k int) int {
    count, result := 0, 0
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil || count >= k { return }
        dfs(node.Left)
        count++
        if count == k {
            result = node.Val
            return
        }
        dfs(node.Right)
    }
    dfs(root)
    return result
}
```

## 模式 3：删除节点（450）

最复杂——找到目标节点后五种情况：

1. **没找到**（root == nil）→ 返回 nil
2. **目标是叶子**（左右都空）→ 返回 nil
3. **左空右非空** → 用右子树继位
4. **右空左非空** → 用左子树继位
5. **左右都非空** → 把左子树挂到右子树最左下角，用右子树继位

```go
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil { return nil }
    if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else {
        // 找到了
        if root.Left == nil { return root.Right }
        if root.Right == nil { return root.Left }
        // 左右都非空：把左挂到右子树最左下角
        cur := root.Right
        for cur.Left != nil { cur = cur.Left }
        cur.Left = root.Left
        return root.Right
    }
    return root
}
```

## 模式 4：构造平衡 BST（108）

有序数组 + 二分递归 = 平衡 BST：取中点为根，左右半数组递归构造左右子树。

```go
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 { return nil }
    mid := len(nums) / 2
    return &TreeNode{
        Val:   nums[mid],
        Left:  sortedArrayToBST(nums[:mid]),
        Right: sortedArrayToBST(nums[mid+1:]),
    }
}
```

## 心法

### 1. 修改类题用「父节点 = 子调用返回值」修剪树指针

凡是要改变树结构的题（插入、删除、构造），递归调用必须用 `=` **把返回值接回到父节点的 `.Left` / `.Right` 字段**：

```go
root.Left = insertIntoBST(root.Left, val)    // ✅ 把可能新建的节点挂到父节点
// 而不是
insertIntoBST(root.Left, val)                // ❌ 新节点返回了但没人接，飘走
```

**两种情况自动统一**：

- 中间节点：返回原 root.Left，赋值无副作用
- base case（在 nil 处建新节点）：返回新节点，赋值让它真正挂上去

**只搜不改的题（700）不需要**——因为不动结构，子调用结果用不上。

**0701 / 0450 / 0108** 都用了这个模式，新刷的修改类树题（删除链表节点等）也通用。

### 2. 中序题的通用骨架：闭包 + 共享状态 + 入口短路

BST 中序遍历得到升序——多数 BST 题都能用这个性质，骨架完全一致：

```go
func solve(root *TreeNode) ResultT {
    // 共享状态（取决于题目要什么）
    var prev *TreeNode      // 验证 BST 用
    // 或 count, result := 0, 0    // 第 k 小用

    var dfs func(*TreeNode) bool
    dfs = func(node *TreeNode) bool {
        if node == nil || /* 终止条件 */ {
            return true
        }
        if !dfs(node.Left) { return false }   // 入口短路：左子树失败立即返回
        // —— 在中序位置处理 node ——
        // 用 prev 或 count 等共享状态做事
        return dfs(node.Right)
    }
    return /* 用 result 或其他 */
}
```

**共享状态变量取什么取决于题目**：

| 题 | 共享状态 | 终止条件 |
|---|---|---|
| 0098 验证 BST | `prev *TreeNode` | `prev.Val >= node.Val` 时 false |
| 0230 第 k 小 | `count + result` | `count >= k` 短路 |
| 未来：最小绝对差、众数 | `prev + best` | 类似 |

**记忆抓手**：题目说"BST" + "中序" / "排序" / "k 小 / 大" / "差" → 用这个骨架。

## 题目清单

- [ ] 700 BST 搜索
- [ ] 701 BST 插入
- [ ] 450 BST 删除
- [ ] 98 验证 BST ⭐ hot100
- [ ] 230 BST 第 k 小 ⭐ hot100
- [ ] 108 有序数组转 BST ⭐ hot100

## 已刷题目

- [0098 验证二叉搜索树](../../problems/二叉树/0098%20验证二叉搜索树.md) 🟡（上下界 + 中序两版）
- [0230 二叉搜索树中第 K 小的元素](../../problems/二叉树/0230%20二叉搜索树中第%20K%20小的元素.md) 🟡（中序闭包 + count 短路）
- [0108 将有序数组转换为二叉搜索树](../../problems/二叉树/0108%20将有序数组转换为二叉搜索树.md) 🟢（数组上二分递归构造）
- [0700 二叉搜索树中的搜索](../../problems/二叉树/0700%20二叉搜索树中的搜索.md) 🟢（左小右大单边递归）
- [0701 二叉搜索树中的插入操作](../../problems/二叉树/0701%20二叉搜索树中的插入操作.md) 🟡（叶子插入 + 返回值修剪树指针）
- [0450 删除二叉搜索树中的节点](../../problems/二叉树/0450%20删除二叉搜索树中的节点.md) 🟡（4 case + 子树重新拼接）

## 🎓 BST 毕业（按 CLAUDE.md「Pattern 推进策略」）

BST 系列 6 题全刷完，覆盖以下"质"：

| 质 | 题 |
|---|---|
| 中序闭包（用中序有序性质） | 0098, 0230 |
| 二分递归构造 | 0108 |
| 左小右大单边递归（搜索） | 0700 |
| 左小右大 + 返回值修剪树指针（插入） | 0701 |
| 子树拼接（删除） | 0450 |

模板默写无压力 → **BST 毕业**。
