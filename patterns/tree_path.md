# Tree Path（二叉树路径 · DFS + 回溯 / 前缀和）

## 何时使用

- 收集从根到叶的所有路径（257）
- 判断/收集和为目标的根到叶路径（112 / 113）
- 计数任意起止的下行路径数（437，路径不必从根开始也不必到叶子结束）

## 模式 1：前序 DFS + 回溯（适合 257 / 112 / 113）

**核心套路**：进入节点时把当前值加入路径，递归处理左右子树，**离开节点时弹出**（回溯）。

```go
func paths(root *TreeNode) [][]int {
    result := [][]int{}
    path := []int{}
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil { return }
        path = append(path, node.Val)             // 进入：加入

        if node.Left == nil && node.Right == nil {
            // 叶子：收集（注意拷贝 path！）
            result = append(result, append([]int{}, path...))
        } else {
            dfs(node.Left)
            dfs(node.Right)
        }

        path = path[:len(path)-1]                 // 离开：弹出（回溯）
    }
    dfs(root)
    return result
}
```

**两个常踩坑**：

1. **必须拷贝 `path` 再 append 到 result**——`append([]int{}, path...)` 创建新切片。否则后续的回溯 `path = path[:len(path)-1]` 会修改已收集的结果（slice 共享底层数组）。
2. **回溯一定要写**——不写就变成"路径只增不减"，下一条路径会拼接上一条的尾巴。

## 模式 2：前缀和 + 哈希表（437 路径总和 III）

437 路径**不必从根开始也不必到叶子结束**——朴素枚举每个节点作为起点是 O(n²)。

**优化**：在 DFS 过程中维护「从根到当前节点的累加和」`pref`。问题转化为：是否存在一个祖先节点 a，使 `pref(curr) - pref(a) == target`？等价于检查哈希表里是否有 `pref(curr) - target` 出现过。

```go
func pathSum(root *TreeNode, target int) int {
    count := 0
    prefSum := map[int]int{0: 1}              // 「虚拟前缀和」：处理"从根开始"的路径

    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, pref int) {
        if node == nil { return }
        pref += node.Val
        count += prefSum[pref - target]       // 命中数 = 之前出现 (pref - target) 的次数
        prefSum[pref]++

        dfs(node.Left, pref)
        dfs(node.Right, pref)

        prefSum[pref]--                       // 回溯：出当前节点的递归子树时减回
    }
    dfs(root, 0)
    return count
}
```

**核心要点**：
- `prefSum[0] = 1`（虚拟前缀和）让从根开始的合法路径也能被统计
- 离开节点时 `prefSum[pref]--`——确保哈希表只记录「**当前路径上**」的祖先前缀和，不会污染兄弟分支
- 注意 `int` 类型溢出（深路径累加）

## 心法

（待积累）

## 题目清单

- [ ] 257 二叉树的所有路径
- [ ] 112 路径总和
- [ ] 113 路径总和 II
- [ ] 437 路径总和 III ⭐ hot100

## 已刷题目

- [0257 二叉树的所有路径](../0257%20二叉树的所有路径.md) 🟢（DFS + 回溯入门）
- [0112 路径总和](../0112%20路径总和.md) 🟢（策略 A：纯 DFS + 减法 + 返回值）
