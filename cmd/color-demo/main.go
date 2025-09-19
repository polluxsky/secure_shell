package main

import (
	"fmt"
	"time"

	"../pkg/term"
)

func main() {
	// 欢迎信息
	term.PrintlnColor("=== 终端颜色和样式演示 ===", term.Bold, term.FgCyan, term.BgBlack)
	fmt.Println()

	// 基本颜色演示
	demoBasicColors()
	fmt.Println()

	// 文本样式演示
	demoTextStyles()
	fmt.Println()

	// 快捷方法演示
	demoShortcuts()
	fmt.Println()

	// 复杂组合演示
	demoComplexCombinations()

	// 结束信息
	term.PrintlnColor("=== 演示结束 ===", term.Bold, term.FgCyan, term.BgBlack)
}

// 演示基本颜色
func demoBasicColors() {
	term.PrintlnColor("基本颜色演示:", term.Bold, term.FgWhite, term.BgBlack)
	
	term.PrintlnColor("黑色文本", term.Reset, term.FgBlack, term.BgBlack)
	term.PrintlnColor("红色文本", term.Reset, term.FgRed, term.BgBlack)
	term.PrintlnColor("绿色文本", term.Reset, term.FgGreen, term.BgBlack)
	term.PrintlnColor("黄色文本", term.Reset, term.FgYellow, term.BgBlack)
	term.PrintlnColor("蓝色文本", term.Reset, term.FgBlue, term.BgBlack)
	term.PrintlnColor("洋红色文本", term.Reset, term.FgMagenta, term.BgBlack)
	term.PrintlnColor("青色文本", term.Reset, term.FgCyan, term.BgBlack)
	term.PrintlnColor("白色文本", term.Reset, term.FgWhite, term.BgBlack)
	
	term.PrintlnColor("亮黑色文本", term.Reset, term.FgBrightBlack, term.BgBlack)
	term.PrintlnColor("亮红色文本", term.Reset, term.FgBrightRed, term.BgBlack)
	term.PrintlnColor("亮绿色文本", term.Reset, term.FgBrightGreen, term.BgBlack)
	term.PrintlnColor("亮黄色文本", term.Reset, term.FgBrightYellow, term.BgBlack)
	term.PrintlnColor("亮蓝色文本", term.Reset, term.FgBrightBlue, term.BgBlack)
	term.PrintlnColor("亮洋红色文本", term.Reset, term.FgBrightMagenta, term.BgBlack)
	term.PrintlnColor("亮青色文本", term.Reset, term.FgBrightCyan, term.BgBlack)
	term.PrintlnColor("亮白色文本", term.Reset, term.FgBrightWhite, term.BgBlack)
}

// 演示文本样式
func demoTextStyles() {
	term.PrintlnColor("文本样式演示:", term.Bold, term.FgWhite, term.BgBlack)
	
	term.PrintlnColor("默认样式", term.Reset, term.FgWhite, term.BgBlack)
	term.PrintlnColor("粗体样式", term.Bold, term.FgWhite, term.BgBlack)
	term.PrintlnColor("暗淡样式", term.Dim, term.FgWhite, term.BgBlack)
	term.PrintlnColor("斜体样式", term.Italic, term.FgWhite, term.BgBlack)
	term.PrintlnColor("下划线样式", term.Underline, term.FgWhite, term.BgBlack)
	term.PrintlnColor("闪烁样式", term.Blink, term.FgWhite, term.BgBlack)
	term.PrintlnColor("反转样式", term.Reverse, term.FgWhite, term.BgBlack)
	term.PrintlnColor("隐藏样式", term.Hidden, term.FgWhite, term.BgBlack)
}

// 演示快捷方法
func demoShortcuts() {
	term.PrintlnColor("快捷方法演示:", term.Bold, term.FgWhite, term.BgBlack)
	
	term.SuccessMessage("这是一条成功消息")
	term.ErrorMessage("这是一条错误消息")
	term.WarningMessage("这是一条警告消息")
	term.InfoMessage("这是一条信息消息")
	
	term.PrintlnColor("这里有一个", term.Reset, term.FgWhite, term.BgBlack)
	term.HighlightText("高亮显示的文本")
	term.PrintlnColor("示例", term.Reset, term.FgWhite, term.BgBlack)
}

// 演示复杂组合
func demoComplexCombinations() {
	term.PrintlnColor("复杂组合演示:", term.Bold, term.FgWhite, term.BgBlack)
	
	// 彩色标题示例
	term.PrintlnColor("=== 彩色标题示例 ===", term.Bold, term.FgYellow, term.BgBlue)
	
	// 彩色表格示例
	term.PrintColor("ID ", term.Bold, term.FgWhite, term.BgBlue)
	term.PrintColor("名称 ", term.Bold, term.FgWhite, term.BgBlue)
	term.PrintlnColor("状态", term.Bold, term.FgWhite, term.BgBlue)
	
	term.PrintColor("1  ", term.Reset, term.FgCyan, term.BgBlack)
	term.PrintColor("项目A ", term.Reset, term.FgCyan, term.BgBlack)
	term.PrintlnColor("进行中", term.Bold, term.FgGreen, term.BgBlack)
	
	term.PrintColor("2  ", term.Reset, term.FgCyan, term.BgBlack)
	term.PrintColor("项目B ", term.Reset, term.FgCyan, term.BgBlack)
	term.PrintlnColor("已完成", term.Bold, term.FgBrightGreen, term.BgBlack)
	
	term.PrintColor("3  ", term.Reset, term.FgCyan, term.BgBlack)
	term.PrintColor("项目C ", term.Reset, term.FgCyan, term.BgBlack)
	term.PrintlnColor("有问题", term.Bold, term.FgRed, term.BgBlack)
	
	// 模拟命令行提示符
	simulateCommandPrompt()
}

// 模拟命令行提示符
func simulateCommandPrompt() {
	term.PrintlnColor("\n命令行提示符模拟:", term.Bold, term.FgWhite, term.BgBlack)
	
	// 模拟bash/zsh风格的提示符
	term.PrintColor("user@machine", term.Bold, term.FgGreen, term.BgBlack)
	term.PrintColor(":", term.Reset, term.FgWhite, term.BgBlack)
	term.PrintColor("~/project", term.Bold, term.FgBlue, term.BgBlack)
	term.PrintColor("$ ", term.Bold, term.FgWhite, term.BgBlack)
	term.PrintlnColor("ls -la", term.Reset, term.FgWhite, term.BgBlack)
	
	// 模拟命令输出
	time.Sleep(500 * time.Millisecond)
	term.PrintlnColor("total 32", term.Reset, term.FgBrightBlue, term.BgBlack)
	term.PrintlnColor("drwxr-xr-x  5 user  staff  160 Sep 19 18:00 .", term.Reset, term.FgBrightBlue, term.BgBlack)
	term.PrintlnColor("drwxr-xr-x  3 user  staff   96 Sep 19 17:00 ..", term.Reset, term.FgBrightBlue, term.BgBlack)
	term.PrintColor("-rw-r--r--  1 user  staff  512 Sep 19 17:50 ", term.Reset, term.FgBrightBlue, term.BgBlack)
	term.PrintlnColor("main.go", term.Reset, term.FgGreen, term.BgBlack)
	term.PrintColor("-rwxr-xr-x  1 user  staff 2048 Sep 19 18:00 ", term.Reset, term.FgBrightBlue, term.BgBlack)
	term.PrintlnColor("color-demo", term.Reset, term.FgRed, term.BgBlack)
	
	// 模拟root提示符
	term.PrintColor("root@machine", term.Bold, term.FgRed, term.BgBlack)
	term.PrintColor(":", term.Reset, term.FgWhite, term.BgBlack)
	term.PrintColor("/root", term.Bold, term.FgBlue, term.BgBlack)
	term.PrintColor("# ", term.Bold, term.FgWhite, term.BgBlack)
	term.PrintlnColor("sudo systemctl status nginx", term.Reset, term.FgWhite, term.BgBlack)
}