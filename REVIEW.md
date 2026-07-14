# 复习计划

复习 = 用户主动来问 → Claude 推荐一道 → 考 pattern 识别 → 用户去 LC 重做。
不强制频率、不维护静态清单、不写 `last_review` 字段。

---

## 触发方式

用户在会话里说"复习一道""今天复一下""我想复 BST"等任意触发语，Claude 启动复习流程。
**不主动推、不强制每天。**

---

## 选题策略（按优先级）

1. **30+ 天前刷过的 ⭐hot100 题** —— 默认池，高价值题维护
2. **历史卡过的题** —— 用户在前序会话里说过"这道卡了"的，Claude 记忆并优先推
3. **用户指定 pattern** —— 用户说"复 BST"则限定该 pattern 内

**复习池只装已毕业的 pattern**。当前正在练的 pattern 不算复习池（它还在主线训练中）。

### Claude 查询到期题（命令模板）

```bash
# 列出所有题目 .md 的首次归档日期，按日期排序
git log --diff-filter=A --name-only --format="%cs" -- "problems/" \
  | awk '/^[0-9]{4}-/ {date=$0} /\.md$/ {print date, $0}' | sort
```

结合 [INDEX.md](INDEX.md) 看哪些是 ⭐hot100，筛出 30+ 天前的候选。

---

## 流程

1. **Claude 推荐一道题**：说明上次刷的时间（如"0098，5/10 刷的，已过 25 天"）
2. **考 pattern 识别两问**：
   - Q1：这是什么 pattern？
   - Q2：**为什么用它？关键特征是什么？**（核心，识别肌肉的训练）
3. **相似题分组提醒**：按"质不同"维度列同 pattern 下其他题（不平铺列表）
4. **用户去 LC 重做**：Claude 不给思路、不给代码
5. **AC 后用户回复"过了"**：Claude 问要不要再来一道

### 用户答不上 Q2

Claude 指引去 `patterns/<系列>/<tag>.md` 看对应章节，看完再来答。**不直接给答案。**

### 用户说"卡了"

Claude 记下下次会话再推同一道（保留在跨会话记忆里）。

---

## 当前复习池（2026-07-15）

已毕业 8 个 pattern：

- **二叉树**：`tree_bfs` / `tree_dfs` / `tree_traversal` / `tree_construction` / `tree_subtree` / `tree_path` / `bst`
- **回溯**：`backtracking`（14 题通关，A/B 分类 + 8 心法）

每毕业一个新 pattern 加入复习池；DP 开刷后同理（练完再进池）。

---

## 不做的事

- ❌ 不写 `last_review` / `mastery` 等 frontmatter 字段（维护成本高）
- ❌ 不强制每天复习（破坏"用户主动"的初衷）
- ❌ 不维护静态"待复习清单"（git log 实时算更准）
- ❌ 不复习 pattern 文档本身（pattern 文档是查阅工具，不是复习对象）
- ❌ 复习时不默写 pattern 模板（背模板 ≠ 识别 pattern）
