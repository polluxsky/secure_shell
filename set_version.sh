#!/bin/bash

# 设置版本号为传入的参数
if [ -z "$1" ]; then
  echo "用法: $0 <版本号>"
  echo "示例: $0 1.1.9"
  exit 1
fi

VERSION_FILE="version.txt"
echo "$1" > "$VERSION_FILE"
echo "版本号已设置为: $1"