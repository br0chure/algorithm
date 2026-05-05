# 题目索引

按**问题域**分组的经典题索引，与 `patterns/` 的「算法 pattern」分组互补。

> **二叉树清单来源**：`~/Documents/go求职/经典题目/二叉树.md`
> 每题标注：✅ 已刷 / ❌ 未刷；箭头指向对应 pattern 文档（去看通用心法）。

---

## 二叉树（6/30）

### 1. 三序遍历 · 递归 + 迭代（0/3）

- ❌ [144 前序遍历](https://leetcode.cn/problems/binary-tree-preorder-traversal/) → [tree_traversal](patterns/tree_traversal.md)
- ❌ [94 中序遍历](https://leetcode.cn/problems/binary-tree-inorder-traversal/) ⭐ hot100 → [tree_traversal](patterns/tree_traversal.md)
- ❌ [145 后序遍历](https://leetcode.cn/problems/binary-tree-postorder-traversal/) → [tree_traversal](patterns/tree_traversal.md)

### 2. 层序遍历 · BFS（3/3）✅

- ✅ [102 二叉树的层序遍历](p0102_binary_tree_level_order_traversal/) ⭐ hot100 → [tree_bfs](patterns/tree_bfs.md)
- ✅ [103 锯齿形层序遍历](p0103_binary_tree_zigzag_level_order_traversal/) → [tree_bfs](patterns/tree_bfs.md)
- ✅ [199 二叉树的右视图](p0199_binary_tree_right_side_view/) ⭐ hot100 → [tree_bfs](patterns/tree_bfs.md)

### 3. 深度与高度 · DFS 后序（3/5）

- ✅ [104 二叉树的最大深度](p0104_maximum_depth_of_binary_tree/) ⭐ hot100 → [tree_dfs](patterns/tree_dfs.md) (Type A)
- ✅ [111 二叉树的最小深度](p0111_minimum_depth_of_binary_tree/) → [tree_dfs](patterns/tree_dfs.md) (Type A 含边界)
- ✅ [110 平衡二叉树](p0110_balanced_binary_tree/) → [tree_dfs](patterns/tree_dfs.md) (Type C 哨兵)
- ✅ [543 二叉树的直径](p0543_diameter_of_binary_tree/) ⭐ hot100 → [tree_dfs](patterns/tree_dfs.md) (Type B)
- ✅ [124 二叉树中的最大路径和](p0124_binary_tree_maximum_path_sum/) ⭐ hot100 → [tree_dfs](patterns/tree_dfs.md) (Type B 进阶)

### 4. 构造二叉树（0/2）

- ❌ [106 从中序与后序构造](https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/) → [tree_construction](patterns/tree_construction.md)
- ❌ [105 从前序与中序构造](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/) ⭐ hot100 → [tree_construction](patterns/tree_construction.md)

### 5. 子结构判断 · 同步双递归（0/3）

- ❌ [100 相同的树](https://leetcode.cn/problems/same-tree/) → [tree_subtree](patterns/tree_subtree.md)
- ❌ [101 对称二叉树](https://leetcode.cn/problems/symmetric-tree/) ⭐ hot100 → [tree_subtree](patterns/tree_subtree.md)
- ❌ [572 另一棵树的子树](https://leetcode.cn/problems/subtree-of-another-tree/) → [tree_subtree](patterns/tree_subtree.md)

### 6. 路径系列 · DFS + 回溯（0/4）

- ❌ [257 二叉树的所有路径](https://leetcode.cn/problems/binary-tree-paths/) → [tree_path](patterns/tree_path.md)
- ❌ [112 路径总和](https://leetcode.cn/problems/path-sum/) → [tree_path](patterns/tree_path.md)
- ❌ [113 路径总和 II](https://leetcode.cn/problems/path-sum-ii/) → [tree_path](patterns/tree_path.md)
- ❌ [437 路径总和 III](https://leetcode.cn/problems/path-sum-iii/) ⭐ hot100 → [tree_path](patterns/tree_path.md)

### 7. 二叉搜索树（0/6）

- ❌ [108 有序数组转 BST](https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/) ⭐ hot100 → [bst](patterns/bst.md)
- ❌ [700 BST 搜索](https://leetcode.cn/problems/search-in-a-binary-search-tree/) → [bst](patterns/bst.md)
- ❌ [701 BST 插入](https://leetcode.cn/problems/insert-into-a-binary-search-tree/) → [bst](patterns/bst.md)
- ❌ [450 BST 删除](https://leetcode.cn/problems/delete-node-in-a-bst/) → [bst](patterns/bst.md)
- ❌ [98 验证 BST](https://leetcode.cn/problems/validate-binary-search-tree/) ⭐ hot100 → [bst](patterns/bst.md)
- ❌ [230 BST 第 k 小](https://leetcode.cn/problems/kth-smallest-element-in-a-bst/) ⭐ hot100 → [bst](patterns/bst.md)

### 8. 其他（0/4）

- ✅ [226 翻转二叉树](p0226_invert_binary_tree/) ⭐ hot100 → [tree_dfs](patterns/tree_dfs.md)
- ❌ [617 合并二叉树](https://leetcode.cn/problems/merge-two-binary-trees/) → [tree_dfs](patterns/tree_dfs.md)
- ❌ [114 二叉树展开为链表](https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/) ⭐ hot100 → [tree_dfs](patterns/tree_dfs.md)
- ❌ [236 最近公共祖先](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/) ⭐ hot100 → [tree_dfs](patterns/tree_dfs.md) (Type B)

---

## 加菜（不在源清单内但已刷）

按 BFS pattern 推进策略额外刷的题，覆盖了源清单未列的"质"：

- ✅ [107 层序遍历 II](p0107_binary_tree_level_order_traversal_ii/) → [tree_bfs](patterns/tree_bfs.md)（顺序变换 · 全局 reverse）
- ✅ [515 在每个树行中找最大值](p0515_find_largest_value_in_each_tree_row/) → [tree_bfs](patterns/tree_bfs.md)（聚合型）
- ✅ [116 填充每个节点的下一个右侧节点指针](p0116_populating_next_right_pointers_in_each_node/) → [tree_bfs](patterns/tree_bfs.md)（横向连接）

---

## 进度统计

| 问题域 | 进度 | 备注 |
|---|---|---|
| 三序遍历 · 迭代 | 0/3 | |
| 层序遍历 · BFS | 3/3 ✅ | BFS pattern 已毕业 |
| 深度与高度 · DFS | 5/5 ✅ | 全完成 |
| 构造二叉树 | 0/2 | tree_construction（骨架已建） |
| 子结构判断 | 0/3 | |
| 路径系列 | 0/4 | |
| 二叉搜索树 | 0/6 | bst（骨架已建） |
| 其他 | 0/4 | |
| **二叉树合计** | **9/30** | |
| 加菜 | 3 | |
| **总计** | **12 题** | |
