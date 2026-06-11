# Tree Traversal（三序遍历：前/中/后序，递归 + 迭代）

## 何时使用

- 题目明确要求按某种顺序输出节点（144 / 94 / 145）
- 中序遍历 BST 得到升序序列（98 / 230 隐式用中序）
- 选择遍历顺序：前/中/后序的差别在于「访问当前节点」相对子树递归的位置——详细的"何时选哪个"见 `tree_dfs.md` 三个核心问题

## 递归实现

三种顺序仅仅是「处理当前节点」的位置不同：

```go
func preorder(node *TreeNode, result *[]int) {
    if node == nil { return }
    *result = append(*result, node.Val)  // 中
    preorder(node.Left, result)          // 左
    preorder(node.Right, result)         // 右
}

func inorder(node *TreeNode, result *[]int) {
    if node == nil { return }
    inorder(node.Left, result)
    *result = append(*result, node.Val)  // 中（左后右前）
    inorder(node.Right, result)
}

func postorder(node *TreeNode, result *[]int) {
    if node == nil { return }
    postorder(node.Left, result)
    postorder(node.Right, result)
    *result = append(*result, node.Val)  // 中（最后）
}
```

## 迭代实现（用显式栈模拟递归）

### 前序：根 → 左 → 右

栈先入右、后入左（保证左先出栈）：

```go
func preorderIterative(root *TreeNode) []int {
    result := []int{}
    if root == nil { return result }
    stack := []*TreeNode{root}
    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, node.Val)
        if node.Right != nil {
            stack = append(stack, node.Right)  // 先入右
        }
        if node.Left != nil {
            stack = append(stack, node.Left)   // 后入左 → 先出
        }
    }
    return result
}
```

### 后序：左 → 右 → 根

直接迭代后序很复杂。**取巧**：用前序变体（中→右→左）跑一遍，最后整体反转 → 得到后序（左→右→中）：

```go
func postorderIterative(root *TreeNode) []int {
    result := []int{}
    if root == nil { return result }
    stack := []*TreeNode{root}
    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, node.Val)
        if node.Left != nil {
            stack = append(stack, node.Left)   // 入左
        }
        if node.Right != nil {
            stack = append(stack, node.Right)  // 入右 → 先出
        }
    }
    // 反转
    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
    return result
}
```

### 中序：访问顺序 ≠ 处理顺序，需指针 + 栈

最复杂——前/后序都是「压谁先访问谁」，但中序必须先沉到左子树最深处再处理。
所以需要一个游标 `cur` 配合栈：

```go
func inorderIterative(root *TreeNode) []int {
    result := []int{}
    stack := []*TreeNode{}
    cur := root
    for cur != nil || len(stack) > 0 {
        for cur != nil {       // 一路向左压栈
            stack = append(stack, cur)
            cur = cur.Left
        }
        cur = stack[len(stack)-1]   // 弹出最深的左节点
        stack = stack[:len(stack)-1]
        result = append(result, cur.Val)  // 处理
        cur = cur.Right         // 转向右子树
    }
    return result
}
```

## 心法

（待积累——刷到 144 / 94 / 145 后补）

## 题目清单

- [ ] 144 二叉树的前序遍历（递归 + 迭代）
- [ ] 94 二叉树的中序遍历 ⭐ hot100（迭代最复杂）
- [ ] 145 二叉树的后序遍历（递归 + 迭代）

## 已刷题目

- [0144 二叉树的前序遍历](../../0144%20二叉树的前序遍历.md) 🟢（迭代：栈先右后左 + 递归 Type A 自递归合并）
- [0145 二叉树的后序遍历](../../0145%20二叉树的后序遍历.md) 🟢（取巧法：伪前序"先左后右"+ 整体反转）
- [0094 二叉树的中序遍历](../../0094%20二叉树的中序遍历.md) 🟢（游标 + 栈，迭代三序最复杂）

## 🎓 tree_traversal 毕业（按 CLAUDE.md「Pattern 推进策略」）

3 题覆盖三种顺序，迭代套路完整：

| 顺序 | 迭代关键技巧 | 题 |
|---|---|---|
| 前序（根→左→右） | 栈先右后左 | 0144 |
| 后序（左→右→根） | 取巧法：伪前序 + 反转 | 0145 |
| 中序（左→根→右） | 游标 + 栈（最复杂） | 0094 |

模板默写无压力 → **tree_traversal 毕业**。
