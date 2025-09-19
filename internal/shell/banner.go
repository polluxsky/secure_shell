package shell

import (
	"fmt"
	"strconv"
	"strings"
)

// 版本号变量，每次编译时自增
var Version = "1.0.0"

// DisplayBanner 显示shell的banner
func DisplayBanner() {
	// 使用字符串切片创建banner
	lines := []string{
		"  ____                                        ____                                   __ ",
		" /\\  _`\\   __ === Copyright by Pollux.Qu === /\\  _`\\ ============= __ ============= /\\ \\__====== ",
		" \\ \\ \\L\\ \\/\\_\\     __   _ __   _ __    __    \\ \\ \\L\\ \\_ __   ___  /\\_\\     __    ___\\ \\ ,_\\ ",
		"  \\ \\ ,__/\\/\\ \\  /'__`\\/\\`'__\\/\\`'__\\/'__`\\   \\ \\ ,__/\\`'__\\/ __`\\\\/\\ \\  /'__`\\ /'___\\ \\ \\/ ",
		"   \\ \\ \\/  \\ \\ \\/\\  __/\\ \\ \\/ \\ \\ \\/\\  __/    \\ \\ \\/\\ \\ \\/\\ \\L\\ \\\\ \\ \\/\\  __//\\ \\__/\\ \\ \\_ ",
		"    \\ \\_\\   \\ \\_\\ \\____\\\\ \\_\\  \\ \\_\\ \\____\\    \\ \\_\\ \\ \\_\\\\ \\____/_\\ \\ \\ \\____\\\\ \\____\\\\ \\__\\ ",
		"     \\/_/    \\/_/\\/____/ \\/_/   \\/_/ \\/____/     \\/_/  \\/_/ \\/___//\\ \\_\\ \\/____/\\/____/ \\/__/ ",
		" ================================================================= \\ \\____/===================== ",
		"       An new Version of Secure Shell                                \\/___/   Version: {$v} ",
		"",
		"",
	}
	
	// 尝试更新版本号
	tryUpdateVersion()

	// 使用strings.Join连接字符串切片
	banner := strings.Join(lines, "\n")
	
	// 替换版本占位符
	banner = strings.Replace(banner, "{$v}", Version, -1)
	
	// 输出banner
	fmt.Println(banner)
}

// tryUpdateVersion 尝试更新版本号
// 从1.0.0开始，每次编译时自增，每个小版本号自增至10时自动进位
func tryUpdateVersion() {
	// 解析版本号
	parts := strings.Split(Version, ".")
	if len(parts) != 3 {
		return
	}
	
	// 将字符串版本号转换为整数
	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])
	
	// 更新版本号（每编译一次最后一位自增1，自增至10时进位）
	patch++
	if patch >= 10 {
		patch = 0
		minor++
		if minor >= 10 {
			minor = 0
			major++
		}
	}
	
	// 更新版本号变量
	Version = fmt.Sprintf("%d.%d.%d", major, minor, patch)
}