# Secure Shell

[English version](README.en.md)

## 安全Shell - 中文说明

Secure Shell是一个增强型的命令行界面，提供安全、高效的终端交互体验。

### 主要功能
- 智能命令自动补全
- 命令历史记录
- 路径自动补全
- 自定义颜色方案
- 安全的命令执行环境

### 安装方法

```bash
# 编译项目
go build -o secure-shell-macos cmd/secure-shell/main.go

# 运行程序
./secure-shell-macos
```

### 使用说明
- 输入命令后按Tab键获取自动补全建议
- 使用上下箭头键浏览命令历史
- 支持基本的文件和目录操作命令