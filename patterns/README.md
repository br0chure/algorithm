# Patterns 工作流

本目录每个文件 = 一种算法 pattern，集中沉淀通用模板和心法。
搭配根目录 [WORKFLOW.md](../WORKFLOW.md)（题目工作流）和 [REVIEW.md](../REVIEW.md)（复习节奏）一起使用。

---

## Pattern 推进策略

掌握一个 pattern 的标准 **不是刷完所有变体**，而是：

- **能默写模板**（不回看 pattern 文档）
- **做过 3 种「质」不同的变体**（不是 3 道同类题）

满足这两条就毕业，切换下一个 pattern。

### 「质不同」的判断

层级动作类型变了、数据结构变了、思路核心变了。

例：BFS 里"收集全部 / 选择性收集 / 聚合 / 顺序变换 / 横向连接"是 5 种"质"。

### 不能因惰性持续推同质题

每完成一题后主动检查是否达到退出标准，达到就建议切换 pattern。
例：515 已练过聚合，637 就不必再刷。

---

## Pattern 文档结构

每个 pattern 一份 markdown（文件名 = tag 名，如 `tree_bfs.md`），含：

| 段落 | 内容 |
|---|---|
| **Pattern 概念** | 什么时候用、典型场景 |
| **Go 通用模板** | 可背诵的标准代码结构 |
| **心法 / 易错点** | 通用陷阱、记忆要点 |
| **常见变体** | 同 pattern 下的题目分类 |
| **已刷题目列表** | 链接到对应 `XXXX 中文标题.md` |

**分工**：通用模板与心法集中在 pattern 文档；每题 .md 只写本题特有的思路、易错点和复杂度。每归档一题，把链接追加进对应 pattern 文档的"已刷题目"列表。

### 心法写入门槛（宁少勿多）

加新心法前过三问自检：
1. 这个心法能从已有概念推出来吗？能推出就不加。
2. 它在未来题里只会出现一两次吗？只用一两次就不加，写题目 README 即可。
3. 它和已有心法重叠吗？重叠就合并不新增。

---

## Pattern 标签字典

每题 frontmatter 的 `tags` 用以下 snake_case 名（与 leetcode-teacher skill 对齐，术语统一）：

### Array & String

| 标签 | 含义 |
|---|---|
| `two_pointers` | Two Pointers |
| `sliding_window` | Sliding Window |
| `fast_slow_pointers` | Fast & Slow Pointers |

### Tree & Graph

| 标签 | 含义 |
|---|---|
| `tree_bfs` | 层序遍历 |
| `tree_dfs` | 深度优先 / 后序聚合通用框架 |
| `tree_traversal` | 三序遍历（含迭代版本，144/94/145） |
| `tree_construction` | 二叉树构造（前/后序 + 中序，105/106） |
| `tree_subtree` | 子结构判断 · 双树同步递归（100/101/572） |
| `tree_path` | 二叉树路径 · DFS + 回溯/前缀和（257/112/113/437） |
| `bst` | 二叉搜索树（98/230/108/700/701/450） |
| `graph_bfs` | Graph BFS |
| `graph_dfs` | Graph DFS |
| `topological_sort` | Topological Sort |

### Dynamic Programming

| 标签 | 含义 |
|---|---|
| `knapsack_01` | 0/1 Knapsack |
| `knapsack_unbounded` | Unbounded Knapsack |
| `fibonacci` | Fibonacci Numbers |
| `house_robber` | 打家劫舍型（选或不选 + 相邻互斥）|
| `lcs` | Longest Common Subsequence |

### Other Essential

| 标签 | 含义 |
|---|---|
| `binary_search` | Modified Binary Search |
| `top_k_elements` | Top K Elements |
| `k_way_merge` | K-Way Merge |
| `backtracking` | Backtracking |
| `union_find` | Union Find |
| `intervals` | Intervals |
| `monotonic_stack` | Monotonic Stack |
| `trie` | Trie |

### 辅助标签

不在 skill 20 种内、但实用：

`hash_map`、`array`、`string`、`linked_list`、`stack`、`queue`、`heap`、`greedy`、`bit_manipulation`、`math`、`recursion`、`simulation`
