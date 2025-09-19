package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 定义ANSI颜色代码
	const (
		Reset          = "\033[0m"
		Bold           = "\033[1m"
		Dim            = "\033[2m"
		Italic         = "\033[3m"
		Underline      = "\033[4m"
		
		// 前景色
		FgBlack        = "\033[30m"
		FgRed          = "\033[31m"
		FgGreen        = "\033[32m"
		FgYellow       = "\033[33m"
		FgBlue         = "\033[34m"
		FgMagenta      = "\033[35m"
		FgCyan         = "\033[36m"
		FgWhite        = "\033[37m"
		FgBrightBlack  = "\033[90m"
		FgBrightRed    = "\033[91m"
		FgBrightGreen  = "\033[92m"
		FgBrightYellow = "\033[93m"
		FgBrightBlue   = "\033[94m"
		FgBrightMagenta = "\033[95m"
		FgBrightCyan   = "\033[96m"
		FgBrightWhite  = "\033[97m"
		
		// 背景色
		BgBlack        = "\033[40m"
		BgRed          = "\033[41m"
		BgGreen        = "\033[42m"
		BgYellow       = "\033[43m"
		BgBlue         = "\033[44m"
		BgMagenta      = "\033[45m"
		BgCyan         = "\033[46m"
		BgWhite        = "\033[47m"
		BgBrightBlack  = "\033[100m"
		BgBrightRed    = "\033[101m"
		BgBrightGreen  = "\033[102m"
		BgBrightYellow = "\033[103m"
		BgBrightBlue   = "\033[104m"
		BgBrightMagenta = "\033[105m"
		BgBrightCyan   = "\033[106m"
		BgBrightWhite  = "\033[107m"
	)

	// 打印带颜色的文本的辅助函数
	printColor := func(text, style, fg, bg string) {
		fmt.Print(style + fg + bg + text + Reset)
		os.Stdout.Sync()
	}

	printlnColor := func(text, style, fg, bg string) {
		fmt.Println(style + fg + bg + text + Reset)
		os.Stdout.Sync()
	}

	// 欢迎信息
	printlnColor("=== 终端颜色和样式演示 ===", Bold, FgCyan, BgBlack)
	fmt.Println()

	// 模拟命令行提示符
	printlnColor("命令行提示符模拟:", Bold, FgWhite, BgBlack)
	
	// 模拟bash/zsh风格的提示符
	printColor("user@machine", Bold, FgGreen, BgBlack)
	printColor(":", Reset, FgWhite, BgBlack)
	printColor("~/project", Bold, FgBlue, BgBlack)
	printColor("$ ", Bold, FgWhite, BgBlack)
	printlnColor("ls -la", Reset, FgWhite, BgBlack)
	
	// 模拟命令输出
	time.Sleep(500 * time.Millisecond)
	printlnColor("total 32", Reset, FgBrightBlue, BgBlack)
	printlnColor("drwxr-xr-x  5 user  staff  160 Sep 19 18:00 .", Reset, FgBrightBlue, BgBlack)
	printlnColor("drwxr-xr-x  3 user  staff   96 Sep 19 17:00 ..", Reset, FgBrightBlue, BgBlack)
	printColor("-rw-r--r--  1 user  staff  512 Sep 19 17:50 ", Reset, FgBrightBlue, BgBlack)
	printlnColor("main.go", Reset, FgGreen, BgBlack)
	printColor("-rwxr-xr-x  1 user  staff 2048 Sep 19 18:00 ", Reset, FgBrightBlue, BgBlack)
	printlnColor("color-demo", Reset, FgRed, BgBlack)
	
	// 模拟root提示符
	printColor("root@machine", Bold, FgRed, BgBlack)
	printColor(":", Reset, FgWhite, BgBlack)
	printColor("/root", Bold, FgBlue, BgBlack)
	printColor("# ", Bold, FgWhite, BgBlack)
	printlnColor("sudo systemctl status nginx", Reset, FgWhite, BgBlack)
	
	fmt.Println()
	printlnColor("=== 演示结束 ===", Bold, FgCyan, BgBlack)
}