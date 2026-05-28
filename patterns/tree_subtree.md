# Tree Subtree（子结构判断 · 双树同步递归）

## 何时使用

- 比较两棵树是否完全相同（100）
- 判断一棵树是否轴对称（101，本质是 root 的左右子树互为镜像）
- 判断一棵树是否包含另一棵作为子树（572）

> **关键标志：函数签名带两个 node 参数**（`f(p, q *TreeNode)`），递归时**两棵树同时推进**。
> 区分：单树 DFS（如 104、三序遍历）虽然也写 `dfs(left); dfs(right)`，但参数只有一个 node——那是「同一棵树的两个子树」，不属于本 pattern。

## Go 通用模板：判断两棵树相同（100）

```go
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil { return true }   // 两个都空 → 相同
    if p == nil || q == nil { return false }  // 一个空一个不空 → 不同
    if p.Val != q.Val { return false }        // 值不同 → 不同
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
```

## 三类变体

### 1. 完全相同（100）

左对左、右对右递归（如上模板）。

### 2. 镜像对称（101）

把根的左右子树看成两棵树，**左对右、右对左**递归：

```go
func isSymmetric(root *TreeNode) bool {
    return isMirror(root.Left, root.Right)
}

func isMirror(p, q *TreeNode) bool {
    if p == nil && q == nil { return true }
    if p == nil || q == nil { return false }
    if p.Val != q.Val { return false }
    return isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)  // ← 交叉
}
```

### 3. 一棵包含另一棵子树（572）

**双重递归**——外层遍历 root 的所有节点，内层判断当前节点是否和 subRoot 相同：

```go
func isSubtree(root, subRoot *TreeNode) bool {
    if root == nil { return false }
    return isSameTree(root, subRoot) ||
           isSubtree(root.Left, subRoot) ||
           isSubtree(root.Right, subRoot)
}
```

## 心法

（待积累）

## 题目清单

- [ ] 100 相同的树
- [ ] 101 对称二叉树 ⭐ hot100
- [ ] 572 另一棵树的子树

## 已刷题目

- [0101 对称二叉树](../0101%20对称二叉树.md) 🟢
- [0100 相同的树](../0100%20相同的树.md) 🟢
- [0572 另一棵树的子树](../0572%20另一棵树的子树.md) 🟢
- [0617 合并二叉树](../0617%20合并二叉树.md) 🟢（双树同步递归 · 构造型 + 原地修改）

## 🎓 tree_subtree 毕业（按 CLAUDE.md「Pattern 推进策略」）

3 种「质」全覆盖：

| 质 | 题 |
|---|---|
| 同向配对 | 0100 |
| 交叉配对（镜像） | 0101 |
| 嵌套递归（双重递归） | 0572 |

模板默写无压力 → **tree_subtree 毕业**。
