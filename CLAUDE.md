# Algorithm 刷题仓库

本仓库用于学习算法，使用 **Go** 完成 LeetCode 等题目。

## 协作规则

- 全程使用中文交流。
- 修改或新增任何项目文件前，必须先向用户说明计划并获得确认。
- 每次写入文件后，立即同步 Git：`git add` + `git commit` + `git push`（远程 `origin → github.com/br0chure/algorithm` 已配置）。

## Skill 使用

- 调用 `leetcode-teacher` skill 时，将其输出翻译成中文后再交付给用户。
- skill 默认输出 Python / TypeScript / Kotlin / Swift。本仓库统一用 Go 归档，必要时把 skill 给出的代码翻译成 Go。

## 刷题流程

- 在 LeetCode 网页写题、运行、提交（网页本身就是测试场）。
- AC 后把代码贴回本地仓库归档。
- **不写本地单元测试**：所有验证以 LeetCode 提交为准；本地仓库只做归档，不放 `*_test.go` 文件。
- **复杂度分析要详细**：不能只甩一行 `O(n)`。要写清楚——
  - **时间**：每个节点 / 元素被访问几次？有没有嵌套？最坏 / 平均 / 摊销分别是什么？
  - **空间**：哪些数据结构占空间（队列、栈、递归栈、辅助哈希）？峰值占用是多少？为什么是这个量级（给出推导，比如「满二叉树最后一层 ≈ n/2 个节点」）？输出本身的 `result` 是否计入？
  - 一眼能看出来的（如双指针 O(n)/O(1)）可以简写；**只要不是一眼能看出来，就要给推导过程**，便于将来翻看时复盘。

## 目录结构（方案 C：题号平铺 + 标签索引）

```
algorithm/
├── CLAUDE.md
├── README.md
├── INDEX.md                   ← 按 pattern 分组的题号索引
├── go.mod
├── patterns/                  ← 各 pattern 的概念 + 模板 + 心法 + 已刷题
│   ├── tree_bfs.md
│   └── ...
├── p0001_two_sum/
│   ├── solution.go            // package p0001
│   └── README.md              // frontmatter + 题面 + 思路 + 复杂度
├── p0015_three_sum/
└── p0206_reverse_linked_list/
```

约定：

- **目录名**：`pXXXX_snake_case_name`（4 位题号 + 蛇形英文短名）。
- **包名**：取目录前缀 `pXXXX`（如 `package p0001`），避免跨题符号冲突。
- **`solution.go`**：解法实现，签名照搬 LeetCode 题目（便于直接粘贴提交）。
- **`README.md`**：按下方模板填写（全文中文）。
- **一题多解**：同目录新增 `solution_v2.go`，或在同 package 内追加函数。

### 每题 README.md 模板

```markdown
---
id: 0001
title: 两数之和
difficulty: 简单
url: https://leetcode.cn/problems/two-sum/
tags: [hash_map, array]
---

# 两数之和

## 题面
（题目描述或链接）

## 思路
...

## 复杂度
- 时间：O(n)，<推导：每个节点访问几次、有无嵌套>
- 空间：O(n)，<推导：哪个数据结构吃空间、峰值多少、为什么>
```

## Pattern 文档

`patterns/` 目录下，每个 pattern 一份 markdown 文档（文件名 = tag，如 `tree_bfs.md`），记录：

- **Pattern 概念**：什么时候用、典型场景
- **Go 通用模板**：可背诵的标准代码结构
- **心法 / 易错点**：通用陷阱与记忆要点
- **常见变体**：同 pattern 下的题目分类
- **已刷题目列表**：链接到对应 `pXXXX_xxx/` 目录

通用模板与心法集中在 pattern 文档里；每题 `README.md` 只写本题特有的思路、易错点和复杂度。每归档一道题，把链接追加进对应 pattern 文档的「已刷题目」列表。

## Pattern 标签字典（与 leetcode-teacher skill 对齐）

每题 README frontmatter 的 `tags` 用以下 snake_case 名（保持术语统一）：

### Array & String Patterns
- `two_pointers` — Two Pointers
- `sliding_window` — Sliding Window
- `fast_slow_pointers` — Fast & Slow Pointers

### Tree & Graph Patterns
- `tree_bfs` — Tree BFS
- `tree_dfs` — Tree DFS
- `graph_bfs` — Graph BFS
- `graph_dfs` — Graph DFS
- `topological_sort` — Topological Sort

### Dynamic Programming Patterns
- `knapsack_01` — 0/1 Knapsack
- `knapsack_unbounded` — Unbounded Knapsack
- `fibonacci` — Fibonacci Numbers
- `lcs` — Longest Common Subsequence

### Other Essential Patterns
- `binary_search` — Modified Binary Search
- `top_k_elements` — Top K Elements
- `k_way_merge` — K-Way Merge
- `backtracking` — Backtracking
- `union_find` — Union Find
- `intervals` — Intervals
- `monotonic_stack` — Monotonic Stack
- `trie` — Trie

辅助标签（不在 skill 20 种内、但实用）：`hash_map`、`array`、`string`、`linked_list`、`stack`、`queue`、`heap`、`greedy`、`bit_manipulation`、`math`、`recursion`、`simulation`。

## INDEX.md

`INDEX.md` 用于按 pattern 列出题号索引。当前手动维护；题量多时可写脚本扫描各题 README frontmatter 的 `tags` 字段自动生成。
