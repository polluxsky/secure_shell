package shell

import (
	"fmt"
	"strings"
)

// 版本号变量，默认为1.1.0
// 实际版本号通过自动版本管理系统更新：
// 运行 ./build.sh 自动递增补丁版本号，采用逢10进位逻辑
// - 补丁版本号达到10时，重置为0，次要版本号+1 (如：1.1.9 -> 1.2.0)
// - 次要版本号达到10时，重置为0，主要版本号+1 (如：1.9.9 -> 2.0.0)
// 版本号存储在version.txt文件中
var Version = "1.1.0"

// DisplayBanner 显示shell的banner
func DisplayBanner() {
	// 使用字符串切片创建简化版banner，避免复杂转义
	lines := []string{
		"=========================================",
		" Secure Shell (Version: {$v})",
		" Copyright by Pollux.Qu",
		"=========================================",
		"",
		"",
	}

	// 使用strings.Join连接字符串切片
	banner := strings.Join(lines, "\n")
	
	// 替换版本占位符
	banner = strings.Replace(banner, "{$v}", Version, -1)
	
	// 输出banner
	fmt.Println(banner)
}