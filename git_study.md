# Git 学习文档

## 1. 基础操作

### 克隆仓库
```bash
git clone https://github.com/用户名/项目名.git
cd 项目名
git status                 # 查看修改状态
git status -s              # 简洁模式
git add .                  # 添加所有修改
git add 文件名             # 添加特定文件
git commit -m "提交说明"    # 提交到本地仓库
git commit -am "提交说明"   # 添加并提交（仅限已跟踪文件）
git push origin main       # 推送到主分支
git push origin 分支名      # 推送到指定分支