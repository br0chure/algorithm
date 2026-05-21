# Tree Construction（二叉树构造）

## 何时使用

- 给定两种遍历序列，重建原二叉树
- 经典组合：**前序 + 中序** 或 **中序 + 后序**

> ⚠️ 仅给前序 + 后序 **无法**唯一重建——左右子树边界识别不出来。

## 核心思路

1. **用前序/后序确定当前根节点**
   - 前序：第一个元素是根
   - 后序：最后一个元素是根
2. **在中序里找到根的位置**——以这个位置为切分点：
   - 左侧 = 左子树中序
   - 右侧 = 右子树中序
3. **根据左子树节点数**反推前序/后序里的左右子树范围
4. **递归构造左右子树**

## Go 通用模板（前序 + 中序 → 树）

用左闭右闭的索引区间表示子序列范围（避免重复切片）：

```go
func buildTree(preorder []int, inorder []int) *TreeNode {
    return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func build(pre []int, preL, preR int, in []int, inL, inR int) *TreeNode {
    if preL > preR { return nil }

    rootVal := pre[preL]
    root := &TreeNode{Val: rootVal}

    // 在中序里找根的位置
    idx := inL
    for in[idx] != rootVal {
        idx++
    }
    leftSize := idx - inL

    // 递归构造左右子树
    root.Left = build(pre, preL+1, preL+leftSize, in, inL, idx-1)
    root.Right = build(pre, preL+leftSize+1, preR, in, idx+1, inR)
    return root
}
```

**关键索引关系**：

- 左子树前序范围：`[preL+1, preL+leftSize]`（紧跟在根后面，长度 = leftSize）
- 右子树前序范围：`[preL+leftSize+1, preR]`（左子树之后到末尾）
- 左子树中序范围：`[inL, idx-1]`
- 右子树中序范围：`[idx+1, inR]`

中序 + 后序变体只需把"前序的根 = 第一个"改成"后序的根 = 最后一个"，其余索引推导逻辑一致。

### 优化：哈希表加速找根

线性查找 `idx` 让总复杂度变成 O(n²)。预建 `inorder` 的 `值 → 索引` 哈希表，每次 O(1) 查找，总复杂度降到 O(n)。

## 心法

（待积累）

## 题目清单

- [ ] 105 从前序与中序构造 ⭐ hot100
- [ ] 106 从中序与后序构造

## 已刷题目

- [0105 从前序与中序遍历序列构造二叉树](../0105%20从前序与中序遍历序列构造二叉树.md) 🟡（前序定根 + 中序切分）
