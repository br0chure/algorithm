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

（待积累）

## 题目清单

- [ ] 700 BST 搜索
- [ ] 701 BST 插入
- [ ] 450 BST 删除
- [ ] 98 验证 BST ⭐ hot100
- [ ] 230 BST 第 k 小 ⭐ hot100
- [ ] 108 有序数组转 BST ⭐ hot100

## 已刷题目

（待积累）
