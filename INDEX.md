# 题目索引

按**问题域**分组的经典题索引，与 `patterns/` 的「算法 pattern」分组互补。

> **二叉树清单来源**：`~/Documents/go求职/经典题目/二叉树.md`
> 每题标注：✅ 已刷 / ❌ 未刷；箭头指向对应 pattern 文档（去看通用心法）。

---

## 二叉树（15/30）

### 1. 三序遍历 · 递归 + 迭代（0/3）

- ❌ [144 前序遍历](https://leetcode.cn/problems/binary-tree-preorder-traversal/) → [tree_traversal](patterns/tree_traversal.md)
- ❌ [94 中序遍历](https://leetcode.cn/problems/binary-tree-inorder-traversal/) ⭐ hot100 → [tree_traversal](patterns/tree_traversal.md)
- ❌ [145 后序遍历](https://leetcode.cn/problems/binary-tree-postorder-traversal/) → [tree_traversal](patterns/tree_traversal.md)

### 2. 层序遍历 · BFS（3/3）✅

- ✅ [0102 二叉树的层序遍历](0102%20二叉树的层序遍历.md) ⭐ hot100 → [tree_bfs](patterns/tree_bfs.md)
- ✅ [0103 二叉树的锯齿形层序遍历](0103%20二叉树的锯齿形层序遍历.md) → [tree_bfs](patterns/tree_bfs.md)
- ✅ [0199 二叉树的右视图](0199%20二叉树的右视图.md) ⭐ hot100 → [tree_bfs](patterns/tree_bfs.md)

### 3. 深度与高度 · DFS 后序（5/5）✅

- ✅ [0104 二叉树的最大深度](0104%20二叉树的最大深度.md) ⭐ hot100 → [tree_dfs](patterns/tree_dfs.md) (Type A)
- ✅ [0111 二叉树的最小深度](0111%20二叉树的最小深度.md) → [tree_dfs](patterns/tree_dfs.md) (Type A 含边界)
- ✅ [0110 平衡二叉树](0110%20平衡二叉树.md) → [tree_dfs](patterns/tree_dfs.md) (Type C 哨兵)
- ✅ [0543 二叉树的直径](0543%20二叉树的直径.md) ⭐ hot100 → [tree_dfs](patterns/tree_dfs.md) (Type B)
- ✅ [0124 二叉树中的最大路径和](0124%20二叉树中的最大路径和.md) ⭐ hot100 → [tree_dfs](patterns/tree_dfs.md) (Type B 进阶)

### 4. 构造二叉树（0/2）

- ❌ [106 从中序与后序构造](https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/) → [tree_construction](patterns/tree_construction.md)
- ❌ [105 从前序与中序构造](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/) ⭐ hot100 → [tree_construction](patterns/tree_construction.md)

### 5. 子结构判断 · 双树同步递归（3/3）✅

- ✅ [0100 相同的树](0100%20相同的树.md) → [tree_subtree](patterns/tree_subtree.md)
- ✅ [0101 对称二叉树](0101%20对称二叉树.md) ⭐ hot100 → [tree_subtree](patterns/tree_subtree.md)
- ✅ [0572 另一棵树的子树](0572%20另一棵树的子树.md) → [tree_subtree](patterns/tree_subtree.md)

### 6. 路径系列 · DFS + 回溯（0/4）

- ❌ [257 二叉树的所有路径](https://leetcode.cn/problems/binary-tree-paths/) → [tree_path](patterns/tree_path.md)
- ❌ [112 路径总和](https://leetcode.cn/problems/path-sum/) → [tree_path](patterns/tree_path.md)
- ❌ [113 路径总和 II](https://leetcode.cn/problems/path-sum-ii/) → [tree_path](patterns/tree_path.md)
- ❌ [437 路径总和 III](https://leetcode.cn/problems/path-sum-iii/) ⭐ hot100 → [tree_path](patterns/tree_path.md)

### 7. 二叉搜索树（3/6）

- ✅ [0108 将有序数组转换为二叉搜索树](0108%20将有序数组转换为二叉搜索树.md) ⭐ hot100 → [bst](patterns/bst.md)
- ❌ [700 BST 搜索](https://leetcode.cn/problems/search-in-a-binary-search-tree/) → [bst](patterns/bst.md)
- ❌ [701 BST 插入](https://leetcode.cn/problems/insert-into-a-binary-search-tree/) → [bst](patterns/bst.md)
- ❌ [450 BST 删除](https://leetcode.cn/problems/delete-node-in-a-bst/) → [bst](patterns/bst.md)
- ✅ [0098 验证二叉搜索树](0098%20验证二叉搜索树.md) ⭐ hot100 → [bst](patterns/bst.md)
- ✅ [0230 二叉搜索树中第 K 小的元素](0230%20二叉搜索树中第%20K%20小的元素.md) ⭐ hot100 → [bst](patterns/bst.md)

### 8. 其他（1/4）

- ✅ [0226 翻转二叉树](0226%20翻转二叉树.md) ⭐ hot100 → [tree_dfs](patterns/tree_dfs.md)
- ❌ [617 合并二叉树](https://leetcode.cn/problems/merge-two-binary-trees/) → [tree_dfs](patterns/tree_dfs.md)
- ❌ [114 二叉树展开为链表](https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/) ⭐ hot100 → [tree_dfs](patterns/tree_dfs.md)
- ❌ [236 最近公共祖先](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/) ⭐ hot100 → [tree_dfs](patterns/tree_dfs.md) (Type B)

---

## 加菜（不在源清单内但已刷）

按 BFS pattern 推进策略额外刷的题，覆盖了源清单未列的"质"：

- ✅ [0107 二叉树的层序遍历 II](0107%20二叉树的层序遍历%20II.md) → [tree_bfs](patterns/tree_bfs.md)（顺序变换 · 全局 reverse）
- ✅ [0515 在每个树行中找最大值](0515%20在每个树行中找最大值.md) → [tree_bfs](patterns/tree_bfs.md)（聚合型）
- ✅ [0116 填充每个节点的下一个右侧节点指针](0116%20填充每个节点的下一个右侧节点指针.md) → [tree_bfs](patterns/tree_bfs.md)（横向连接）

---

## 进度统计

| 问题域 | 进度 | 备注 |
|---|---|---|
| 三序遍历 · 迭代 | 0/3 | |
| 层序遍历 · BFS | 3/3 ✅ | BFS pattern 已毕业 |
| 深度与高度 · DFS | 5/5 ✅ | 全完成 |
| 构造二叉树 | 0/2 | tree_construction（骨架已建） |
| 子结构判断 | 3/3 ✅ | tree_subtree 毕业 |
| 路径系列 | 0/4 | |
| 二叉搜索树 | 3/6 | bst 中序 + 构造维度都覆盖 |
| 其他 | 1/4 | |
| **二叉树合计** | **15/30** | |
| 加菜 | 3 | |
| **总计** | **18 题** | |
