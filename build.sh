#!/bin/bash

# 版本配置文件路径
VERSION_FILE="version.txt"

# 检查版本文件是否存在，如果不存在则创建并设置默认版本
if [ ! -f "$VERSION_FILE" ]; then
  echo "1.0.0" > "$VERSION_FILE"
  echo "创建默认版本配置文件: $VERSION_FILE，版本号: 1.0.0"
fi

# 读取当前版本号
CURRENT_VERSION=$(cat "$VERSION_FILE")

# 解析版本号为主要、次要和补丁版本
IFS='.' read -r MAJOR MINOR PATCH <<< "$CURRENT_VERSION"

# 递增补丁版本号
((PATCH++))

# 实现逢10进位逻辑
if [ $PATCH -ge 10 ]; then
  PATCH=0
  ((MINOR++))
  
  if [ $MINOR -ge 10 ]; then
    MINOR=0
    ((MAJOR++))
    
    # 可选：如果主版本号也想限制在10以内，可以取消下面两行的注释
    # if [ $MAJOR -ge 10 ]; then
    #   MAJOR=0
    # fi
  fi
fi

# 构建新版本号
NEW_VERSION="$MAJOR.$MINOR.$PATCH"

# 编译命令
GOROOT=/usr/local/Cellar/go/1.25.0/libexec GOOS=darwin GOARCH=amd64 go build -ldflags "-X 'secure-shell/internal/shell.Version=$NEW_VERSION'" -o secure-shell-macos cmd/secure-shell/main.go

# 检查编译是否成功
if [ $? -eq 0 ]; then
  # 更新版本文件
  echo "$NEW_VERSION" > "$VERSION_FILE"
  echo "编译成功! 版本号已从 $CURRENT_VERSION 更新为: $NEW_VERSION"
  echo "生成的可执行文件: secure-shell-macos"
else
  echo "编译失败，版本号未更新 (当前版本: $CURRENT_VERSION)"
  exit 1
fi