---
id: 0226
title: 翻转二叉树
difficulty: 简单
url: https://leetcode.cn/problems/invert-binary-tree/
tags: [tree_dfs]
---

# 翻转二叉树

## 题面

给你一棵二叉树的根节点 `root`，**翻转**这棵二叉树（每个节点的左右子树整体互换），返回根节点。

样例：`[4,2,7,1,3,6,9]` → `[4,7,2,9,6,3,1]`。

约束：节点数 `0 ≤ n ≤ 100`，`-100 ≤ Node.val ≤ 100`。

## 思路

**Pattern：tree_dfs · Type A · in-place 修改**。

每个节点要做的事：把它的左右孩子指针互换。然后递归处理子树。

### 完整代码

```go
func invertTree(root *TreeNode) *TreeNode {
    invert(root)
    return root
}

func invert(root *TreeNode) {
    if root == nil { return }
    invert(root.Left)
    invert(root.Right)
    root.Left, root.Right = root.Right, root.Left
}
```

> 用 void helper 是为了**显式表达「这是 in-place 修改」**——所有翻转动作都改在原树上，外层 `invertTree` 返回 root 只是为了对齐 LC 签名要求，并不携带计算结果。等价的写法是直接让 `invertTree` 自己递归（省掉 helper）；两种都对，看个人偏好。

## 前 / 中 / 后序对照

| 顺序 | 是否 AC | 说明 |
|---|---|---|
| **前序**（先 swap，再递归） | ✅ | 每次先交换当前 root 的左右指针，再去翻转新位置上的两棵子树 |
| **后序**（先递归，再 swap） | ✅ | 主解写法。先把两棵子树各自翻完，最后交换 |
| **中序**（递归左 + swap + 递归右） | ❌ | **会写错** |

### 为什么中序会错

```go
invert(root.Left)              // 翻转原左子树 ✓
root.Left, root.Right = ...    // swap → 此时 root.Left 是原右子树，root.Right 是已翻转的原左子树
invert(root.Right)             // ❌ 这里翻的是「已翻转的原左子树」，反翻了
```

关键陷阱：**swap 改变了 `root.Left` / `root.Right` 的指向**。中序在 swap 之后再递归 `root.Right`，访问到的是 swap 后的位置——不再是原右子树。

> 这是「翻转操作不可交换」的直观体现：swap 与递归的相对顺序敏感。前序在 swap 前没递归，后序在 swap 后没递归，都不踩这个坑。

## 复杂度

- **时间：O(n)**——每个节点恰好被访问 1 次，节点级动作（指针互换）O(1)。
- **空间：O(h)**——递归栈，h 为树高（最坏链状 O(n)、平衡 O(log n)）。

  栈帧只装 `root` 指针 + 元信息，O(1)；函数体没有任何动态分配。
