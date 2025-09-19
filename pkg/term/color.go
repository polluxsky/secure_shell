package term

import (
	"fmt"
	"os"
	"runtime"
)

// 文本样式常量
type TextStyle string
const (
	Reset          TextStyle = "\033[0m"
	Bold           TextStyle = "\033[1m"
	Dim            TextStyle = "\033[2m"
	Italic         TextStyle = "\033[3m"
	Underline      TextStyle = "\033[4m"
	Blink          TextStyle = "\033[5m"
	Reverse        TextStyle = "\033[7m"
	Hidden         TextStyle = "\033[8m"
)

// 前景色常量
type ForegroundColor string
const (
	FgBlack        ForegroundColor = "\033[30m"
	FgRed          ForegroundColor = "\033[31m"
	FgGreen        ForegroundColor = "\033[32m"
	FgYellow       ForegroundColor = "\033[33m"
	FgBlue         ForegroundColor = "\033[34m"
	FgMagenta      ForegroundColor = "\033[35m"
	FgCyan         ForegroundColor = "\033[36m"
	FgWhite        ForegroundColor = "\033[37m"
	FgBrightBlack  ForegroundColor = "\033[90m"
	FgBrightRed    ForegroundColor = "\033[91m"
	FgBrightGreen  ForegroundColor = "\033[92m"
	FgBrightYellow ForegroundColor = "\033[93m"
	FgBrightBlue   ForegroundColor = "\033[94m"
	FgBrightMagenta ForegroundColor = "\033[95m"
	FgBrightCyan   ForegroundColor = "\033[96m"
	FgBrightWhite  ForegroundColor = "\033[97m"
)

// 背景色常量
type BackgroundColor string
const (
	BgBlack        BackgroundColor = "\033[40m"
	BgRed          BackgroundColor = "\033[41m"
	BgGreen        BackgroundColor = "\033[42m"
	BgYellow       BackgroundColor = "\033[43m"
	BgBlue         BackgroundColor = "\033[44m"
	BgMagenta      BackgroundColor = "\033[45m"
	BgCyan         BackgroundColor = "\033[46m"
	BgWhite        BackgroundColor = "\033[47m"
	BgBrightBlack  BackgroundColor = "\033[100m"
	BgBrightRed    BackgroundColor = "\033[101m"
	BgBrightGreen  BackgroundColor = "\033[102m"
	BgBrightYellow BackgroundColor = "\033[103m"
	BgBrightBlue   BackgroundColor = "\033[104m"
	BgBrightMagenta BackgroundColor = "\033[105m"
	BgBrightCyan   BackgroundColor = "\033[106m"
	BgBrightWhite  BackgroundColor = "\033[107m"
)

// ColorText 给文本添加颜色和样式
func ColorText(text string, style TextStyle, fgColor ForegroundColor, bgColor BackgroundColor) string {
	// 检查是否在Windows系统上运行，如果是，可能需要特殊处理
	if runtime.GOOS == "windows" {
		// Windows命令提示符可能不完全支持所有ANSI颜色代码
		// 这里简单返回原始文本，实际项目中可以使用更复杂的Windows终端颜色库
		return text
	}
	return string(style) + string(fgColor) + string(bgColor) + text + string(Reset)
}

// PrintColor 打印带颜色的文本到标准输出
func PrintColor(text string, style TextStyle, fgColor ForegroundColor, bgColor BackgroundColor) {
	fmt.Print(ColorText(text, style, fgColor, bgColor))
	os.Stdout.Sync()
}

// PrintlnColor 打印带颜色的文本并换行
func PrintlnColor(text string, style TextStyle, fgColor ForegroundColor, bgColor BackgroundColor) {
	fmt.Println(ColorText(text, style, fgColor, bgColor))
	os.Stdout.Sync()
}

// Example: 一些常用的快捷方法

// SuccessMessage 打印成功消息（绿色文本）
func SuccessMessage(text string) {
	PrintlnColor(text, Bold, FgGreen, BgBlack)
}

// ErrorMessage 打印错误消息（红色文本）
func ErrorMessage(text string) {
	PrintlnColor(text, Bold, FgRed, BgBlack)
}

// WarningMessage 打印警告消息（黄色文本）
func WarningMessage(text string) {
	PrintlnColor(text, Bold, FgYellow, BgBlack)
}

// InfoMessage 打印信息消息（蓝色文本）
func InfoMessage(text string) {
	PrintlnColor(text, Bold, FgBlue, BgBlack)
}

// HighlightText 高亮显示文本（黄色背景）
func HighlightText(text string) {
	PrintColor(text, Bold, FgBlack, BgYellow)
}