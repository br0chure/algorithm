# Tree Subtree（子结构判断 · 同步双递归）

## 何时使用

- 比较两棵树是否完全相同（100）
- 判断一棵树是否轴对称（101，本质是 root 的左右子树互为镜像）
- 判断一棵树是否包含另一棵作为子树（572）

## 核心思路

**两棵树同时递归**——基础情况同时处理两个 nil/非 nil 组合，递归时同步前进。

## Go 通用模板：判断两棵树相同（100）

```go
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil { return true }   // 两个都空 → 相同
    if p == nil || q == nil { return false }  // 一个空一个不空 → 不同
    if p.Val != q.Val { return false }        // 值不同 → 不同
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
```

`base case` 三种组合：

- 两个都是 nil → 相同
- 一个 nil 一个非 nil → 不同
- 两个都非 nil → 比较 Val，再递归子树

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

（待积累）
