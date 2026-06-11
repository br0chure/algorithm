# Algorithm 刷题仓库

本仓库用于学习算法，使用 **Go** 完成 LeetCode 等题目。

## 协作规则

- 全程使用中文交流。
- 修改或新增任何项目文件前，必须先向用户说明计划并获得确认。
- 每次写入文件后，立即同步 Git：`git add` + `git commit` + `git push`（远程 `origin → github.com/br0chure/algorithm` 已配置）。

## Skill 使用

- 调用 `leetcode-teacher` skill 时，将其输出翻译成中文后再交付给用户。
- skill 默认输出 Python / TypeScript / Kotlin / Swift。本仓库统一用 Go 归档，必要时把 skill 给出的代码翻译成 Go。

## 目录结构（一题一文件，平铺）

```
algorithm/
├── CLAUDE.md           ← 本文件，协作规则 + 文档导航
├── WORKFLOW.md         ← 刷题流程 + 复杂度分析要求 + 每题模板
├── REVIEW.md           ← 复习计划
├── INDEX.md            ← 按问题域分组的题号索引
├── README.md
├── patterns/           ← 各 pattern 的概念 + 模板 + 心法 + 已刷题
│   ├── README.md       ← Pattern 推进策略 + 标签字典
│   ├── 二叉树/          ← 按系列分子目录
│   │   ├── tree_bfs.md  tree_dfs.md  bst.md  ……
│   │   └── （三序 / 构造 / 子结构 / 路径）
│   └── 回溯/            ← 新系列各自独立，互不混杂
├── 0001 两数之和.md      ← 一题一个 markdown，文件名 = 题号 + 中文标题
├── 0102 二叉树的层序遍历.md
└── ...
```

**命名约定**：
- 题目文件：`XXXX 中文标题.md`（4 位题号 + 一个空格 + 中文标题）
- 运维文件：全大写英文（`CLAUDE.md` / `WORKFLOW.md` / `REVIEW.md` / `INDEX.md`），视觉上与题目区分
- 代码：直接嵌入题目 markdown 的 ` ```go ... ``` ` 代码块，**不创建独立 `.go` 文件**

## 文档导航

| 文档 | 何时查阅 |
|---|---|
| [WORKFLOW.md](WORKFLOW.md) | 写新题归档时 —— 流程、复杂度分析要求、题目模板 |
| [patterns/README.md](patterns/README.md) | 切 pattern / 写 pattern 文档时 —— 推进策略、毕业标准、标签字典 |
| [REVIEW.md](REVIEW.md) | 用户说"复习一道"时 —— 选题策略、考 pattern 识别、流程 |
| [INDEX.md](INDEX.md) | 想找某题 / 看进度时 —— 按问题域分组的题号索引 |
