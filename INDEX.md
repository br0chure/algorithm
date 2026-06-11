# 题目索引

按**问题域**分组的经典题索引，与 `patterns/` 的「算法 pattern」分组互补。

> **清单来源**：各系列对应 `~/Documents/go求职/经典题目/<类型>.md`（理论基础 + 精选题）。
> 每题标注：✅ 已刷 / 🔄 进行中 / ❌ 未刷；箭头指向对应 pattern 文档（去看通用心法）。

---

## 二叉树（30/30）🎉

### 1. 三序遍历 · 递归 + 迭代（3/3）✅

- ✅ [0144 二叉树的前序遍历](0144%20二叉树的前序遍历.md) → [tree_traversal](patterns/二叉树/tree_traversal.md)
- ✅ [0094 二叉树的中序遍历](0094%20二叉树的中序遍历.md) ⭐ hot100 → [tree_traversal](patterns/二叉树/tree_traversal.md)
- ✅ [0145 二叉树的后序遍历](0145%20二叉树的后序遍历.md) → [tree_traversal](patterns/二叉树/tree_traversal.md)

### 2. 层序遍历 · BFS（3/3）✅

- ✅ [0102 二叉树的层序遍历](0102%20二叉树的层序遍历.md) ⭐ hot100 → [tree_bfs](patterns/二叉树/tree_bfs.md)
- ✅ [0103 二叉树的锯齿形层序遍历](0103%20二叉树的锯齿形层序遍历.md) → [tree_bfs](patterns/二叉树/tree_bfs.md)
- ✅ [0199 二叉树的右视图](0199%20二叉树的右视图.md) ⭐ hot100 → [tree_bfs](patterns/二叉树/tree_bfs.md)

### 3. 深度与高度 · DFS 后序（5/5）✅

- ✅ [0104 二叉树的最大深度](0104%20二叉树的最大深度.md) ⭐ hot100 → [tree_dfs](patterns/二叉树/tree_dfs.md) (Type A)
- ✅ [0111 二叉树的最小深度](0111%20二叉树的最小深度.md) → [tree_dfs](patterns/二叉树/tree_dfs.md) (Type A 含边界)
- ✅ [0110 平衡二叉树](0110%20平衡二叉树.md) → [tree_dfs](patterns/二叉树/tree_dfs.md) (Type C 哨兵)
- ✅ [0543 二叉树的直径](0543%20二叉树的直径.md) ⭐ hot100 → [tree_dfs](patterns/二叉树/tree_dfs.md) (Type B)
- ✅ [0124 二叉树中的最大路径和](0124%20二叉树中的最大路径和.md) ⭐ hot100 → [tree_dfs](patterns/二叉树/tree_dfs.md) (Type B 进阶)

### 4. 构造二叉树（2/2）✅

- ✅ [0106 从中序与后序遍历序列构造二叉树](0106%20从中序与后序遍历序列构造二叉树.md) → [tree_construction](patterns/二叉树/tree_construction.md)
- ✅ [0105 从前序与中序遍历序列构造二叉树](0105%20从前序与中序遍历序列构造二叉树.md) ⭐ hot100 → [tree_construction](patterns/二叉树/tree_construction.md)

### 5. 子结构判断 · 双树同步递归（3/3）✅

- ✅ [0100 相同的树](0100%20相同的树.md) → [tree_subtree](patterns/二叉树/tree_subtree.md)
- ✅ [0101 对称二叉树](0101%20对称二叉树.md) ⭐ hot100 → [tree_subtree](patterns/二叉树/tree_subtree.md)
- ✅ [0572 另一棵树的子树](0572%20另一棵树的子树.md) → [tree_subtree](patterns/二叉树/tree_subtree.md)

### 6. 路径系列 · DFS + 回溯（4/4）✅

- ✅ [0257 二叉树的所有路径](0257%20二叉树的所有路径.md) → [tree_path](patterns/二叉树/tree_path.md)
- ✅ [0112 路径总和](0112%20路径总和.md) → [tree_path](patterns/二叉树/tree_path.md)
- ✅ [0113 路径总和 II](0113%20路径总和%20II.md) → [tree_path](patterns/二叉树/tree_path.md)
- ✅ [0437 路径总和 III](0437%20路径总和%20III.md) ⭐ hot100 → [tree_path](patterns/二叉树/tree_path.md)

### 7. 二叉搜索树（6/6）✅

- ✅ [0108 将有序数组转换为二叉搜索树](0108%20将有序数组转换为二叉搜索树.md) ⭐ hot100 → [bst](patterns/二叉树/bst.md)
- ✅ [0700 二叉搜索树中的搜索](0700%20二叉搜索树中的搜索.md) → [bst](patterns/二叉树/bst.md)
- ✅ [0701 二叉搜索树中的插入操作](0701%20二叉搜索树中的插入操作.md) → [bst](patterns/二叉树/bst.md)
- ✅ [0450 删除二叉搜索树中的节点](0450%20删除二叉搜索树中的节点.md) → [bst](patterns/二叉树/bst.md)
- ✅ [0098 验证二叉搜索树](0098%20验证二叉搜索树.md) ⭐ hot100 → [bst](patterns/二叉树/bst.md)
- ✅ [0230 二叉搜索树中第 K 小的元素](0230%20二叉搜索树中第%20K%20小的元素.md) ⭐ hot100 → [bst](patterns/二叉树/bst.md)

### 8. 其他（4/4）✅

- ✅ [0226 翻转二叉树](0226%20翻转二叉树.md) ⭐ hot100 → [tree_dfs](patterns/二叉树/tree_dfs.md)
- ✅ [0617 合并二叉树](0617%20合并二叉树.md) → [tree_subtree](patterns/二叉树/tree_subtree.md)
- ✅ [0114 二叉树展开为链表](0114%20二叉树展开为链表.md) ⭐ hot100 → [tree_dfs](patterns/二叉树/tree_dfs.md) + [tree_traversal](patterns/二叉树/tree_traversal.md)
- ✅ [0236 二叉树的最近公共祖先](0236%20二叉树的最近公共祖先.md) ⭐ hot100 → [tree_dfs](patterns/二叉树/tree_dfs.md) (Type B)

---

## 加菜（不在源清单内但已刷）

按 BFS pattern 推进策略额外刷的题，覆盖了源清单未列的"质"：

- ✅ [0107 二叉树的层序遍历 II](0107%20二叉树的层序遍历%20II.md) → [tree_bfs](patterns/二叉树/tree_bfs.md)（顺序变换 · 全局 reverse）
- ✅ [0515 在每个树行中找最大值](0515%20在每个树行中找最大值.md) → [tree_bfs](patterns/二叉树/tree_bfs.md)（聚合型）
- ✅ [0116 填充每个节点的下一个右侧节点指针](0116%20填充每个节点的下一个右侧节点指针.md) → [tree_bfs](patterns/二叉树/tree_bfs.md)（横向连接）

---

## 回溯（0/14）🔜

> 清单来源：`~/Documents/go求职/经典题目/回溯.md`
> 回溯骨架 = **选择 → 递归 → 撤销**，与二叉树 [tree_path](patterns/二叉树/tree_path.md) 一脉相承。pattern 文档放 `patterns/回溯/`，随做题沉淀。题目归档后再补 `.md` 链接。

### 1. 组合 · startIdx 控制不重复

- ❌ 0077 组合
- ❌ 0216 组合总和 III
- ❌ 0039 组合总和 ⭐ hot100
- ❌ 0040 组合总和 II（同层去重）
- ❌ 0017 电话号码的字母组合 ⭐ hot100
- ❌ 0131 分割回文串 ⭐ hot100
- ❌ 0093 复原 IP 地址

### 2. 子集

- ❌ 0078 子集 ⭐ hot100
- ❌ 0090 子集 II（同层去重）

### 3. 排列

- 🔄 0046 全排列 ⭐ hot100 ← **破冰进行中**
- ❌ 0047 全排列 II（同层去重）

### 4. 其他 · 棋盘 / 网格 / 字符串

- ❌ 0022 括号生成 ⭐ hot100
- ❌ 0079 单词搜索 ⭐ hot100
- ❌ 0051 N 皇后 ⭐ hot100

---

## 进度统计

| 问题域 | 进度 | 备注 |
|---|---|---|
| 三序遍历 · 迭代 | 3/3 ✅ | tree_traversal 毕业 |
| 层序遍历 · BFS | 3/3 ✅ | BFS pattern 已毕业 |
| 深度与高度 · DFS | 5/5 ✅ | 全完成 |
| 构造二叉树 | 2/2 ✅ | tree_construction 毕业 |
| 子结构判断 | 3/3 ✅ | tree_subtree 毕业 |
| 路径系列 | 4/4 ✅ | tree_path 毕业（437 前缀和优化待补）|
| 二叉搜索树 | 6/6 ✅ | bst 毕业 |
| 其他 | 4/4 ✅ | 全完成 |
| **二叉树合计** | **30/30** 🎉 | 全部完成 |
| 加菜 | 3 | |
| **回溯系列** | **0/14** 🔜 | 0046 破冰中（组合 / 子集 / 排列 / 其他）|
| **总计** | **已归档 33 题** | 回溯 14 道待刷 |
