package shell

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"os/user"

	"github.com/c-bata/go-prompt"
)

// 显示命令提示符（保留旧函数以保持兼容性）
func showPrompt() {
	username := getUsername()
	hostname := getHostname()
	currDir := getCurrentDir()
	fmt.Printf("[%s@%s %s]$ ", username, hostname, currDir)
	os.Stdout.Sync()
}

// 获取当前用户名
func getUsername() string {
	u, err := user.Current()
	if err != nil {
		return "unknown"
	}
	return u.Username
}

// 获取主机名
func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hostname
}

// 获取当前工作目录，格式化为相对路径（如果在用户主目录下）
func getCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return "unknown"
	}

	// 尝试将路径转换为相对于home目录的形式
	u, err := user.Current()
	if err == nil && strings.HasPrefix(dir, u.HomeDir) {
		return "~" + dir[len(u.HomeDir):]
	}
	return dir
}

// 创建自定义的命令行提示器
func createPrompt() *prompt.Prompt {
	// 动态前缀函数（用于OptionLivePrefix）
	livePrefix := func() (string, bool) {
		// 获取当前目录
		currentDir, err := os.Getwd()
		if err != nil {
			currentDir = "unknown"
		}
		
		// 简化显示，将用户主目录替换为~符号
		if homeDir, err := os.UserHomeDir(); err == nil {
			if strings.HasPrefix(currentDir, homeDir) {
				currentDir = "~" + currentDir[len(homeDir):]
			}
		}
		
		// 格式化提示前缀（不使用直接的终端转义序列）
		return fmt.Sprintf("%s $ ", currentDir), true
	}

	// 自动补全函数
	autoComplete := func(d prompt.Document) []prompt.Suggest {
		// 获取输入内容
		input := d.Text
		parts := strings.Fields(input)
		var prefix string
		var suggestions []prompt.Suggest

		if len(parts) == 0 {
			// 没有输入，返回所有命令
			for cmd := range allowedCommands {
				suggestions = append(suggestions, prompt.Suggest{
					Text: cmd,
					Description: getCommandDescription(cmd),
				})
			}
			return suggestions
		} else if len(parts) == 1 && strings.TrimSpace(input) == parts[0] {
			// 只有命令部分，可能需要补全命令
			prefix = parts[0]
			// 尝试命令补全
			for cmd := range allowedCommands {
				if strings.HasPrefix(cmd, prefix) {
					suggestions = append(suggestions, prompt.Suggest{
						Text: cmd,
						Description: getCommandDescription(cmd),
					})
				}
			}

			// 如果没有命令补全，尝试路径补全
			if len(suggestions) == 0 {
				pathSuggestions := getPathSuggestions(prefix)
				suggestions = append(suggestions, pathSuggestions...)
			}
		} else {
			// 有参数，补全最后一个参数（路径）
			prefix = parts[len(parts)-1]
			suggestions = getPathSuggestions(prefix)
		}

		return suggestions
	}
	
	// 历史记录处理函数
	history := func(s string) {
		if s != "" {
			commandHistory = append(commandHistory, s)
		}
	}
	
	// 创建并返回prompt实例
	return prompt.New(
		// 执行命令的函数
		func(in string) {
			in = strings.TrimSpace(in)
			if in == "" {
				return
			}
			
			// 添加到历史记录
			history(in)
			
			// 执行命令
			continueShell := executeCommand(in)
			if !continueShell {
				// 正常退出
				os.Exit(normalExitCode)
			}
		},
		// 自动补全函数
		autoComplete,
		// 配置选项 - 更美观的配色方案
		prompt.OptionLivePrefix(livePrefix),
		prompt.OptionPrefixBackgroundColor(prompt.DarkBlue), // 深蓝色前缀背景
		prompt.OptionPrefixTextColor(prompt.LightGray),      // 浅灰色前缀文本
		prompt.OptionInputTextColor(prompt.Cyan),            // 青色输入文本
		prompt.OptionSuggestionTextColor(prompt.LightGray),  // 浅灰色建议文本
		prompt.OptionSuggestionBGColor(prompt.DarkGray),     // 深灰色建议背景
		prompt.OptionSelectedSuggestionTextColor(prompt.Black), // 选中建议的黑色文本
		prompt.OptionSelectedSuggestionBGColor(prompt.Cyan),    // 选中建议的青色背景
		prompt.OptionDescriptionTextColor(prompt.LightGray), // 浅灰色描述文本
		prompt.OptionDescriptionBGColor(prompt.DarkGray),    // 深灰色描述背景
		prompt.OptionSelectedDescriptionTextColor(prompt.Black), // 选中描述的黑色文本
		prompt.OptionSelectedDescriptionBGColor(prompt.Cyan),    // 选中描述的青色背景
		prompt.OptionHistory([]string{}), // 空的历史记录初始化
		prompt.OptionMaxSuggestion(0), // 设置最大显示建议数为0，禁用下拉提示但保留自动补全功能
		prompt.OptionAddKeyBind(prompt.KeyBind{
			Key: prompt.ControlC,
			Fn: func(buffer *prompt.Buffer) {
				// 忽略Ctrl+C，继续运行
				buffer.InsertText("", true, false)
			},
		}),
	)
}

// 获取命令描述
func getCommandDescription(cmd string) string {
	descriptions := map[string]string{
		"ls":      "List directory contents",
		"ll":      "List directory contents in long format",
		"cd":      "Change the working directory",
		"cat":     "Concatenate and display files",
		"pwd":     "Print the current working directory",
		"vim":     "Open file with Vim editor",
		"mkdir":   "Create new directory",
		"history": "Show command history",
		"exit":    "Exit the shell",
	}
	return descriptions[cmd]
}

// 获取路径补全建议
func getPathSuggestions(prefix string) []prompt.Suggest {
	var suggestions []prompt.Suggest
	
	// 处理波浪号
	expandedPrefix := prefix
	if strings.HasPrefix(prefix, "~") {
		u, err := user.Current()
		if err == nil {
			expandedPrefix = u.HomeDir + prefix[1:]
		}
	}
	
	// 获取目录和文件名
	dir := filepath.Dir(expandedPrefix)
	file := filepath.Base(expandedPrefix)
	
	// 读取目录内容
	files, err := os.ReadDir(dir)
	if err != nil {
		return suggestions
	}
	
	// 查找匹配的文件/目录
	for _, f := range files {
		name := f.Name()
		if strings.HasPrefix(name, file) {
			fullPath := filepath.Join(dir, name)
			var description string
			var displayName string
			
			// 如果是目录，添加斜杠和描述
			if f.IsDir() {
				name += "/"
				description = "directory"
			} else {
				description = "file"
			}
			
			// 构建显示路径
			displayName = name
			if dir != "." {
				relPath, err := filepath.Rel(".", fullPath)
				if err == nil {
					displayName = relPath
					if f.IsDir() && !strings.HasSuffix(displayName, "/") {
						displayName += "/"
					}
				}
			}
			
			// 替换回波浪号表示法
			if strings.HasPrefix(prefix, "~") {
				u, err := user.Current()
				if err == nil && strings.HasPrefix(displayName, u.HomeDir) {
					displayName = "~" + displayName[len(u.HomeDir):]
				}
			}
			
			suggestions = append(suggestions, prompt.Suggest{
				Text:        displayName,
				Description: description,
			})
		}
	}
	
	return suggestions
}

// 搜索历史记录
func searchHistory(text string, index int) (string, int) {
	if len(commandHistory) == 0 {
		return "", 0
	}
	
	// 简单的历史搜索
	start := index
	if index >= len(commandHistory) {
		start = 0
	}
	
	for i := start; i < len(commandHistory); i++ {
		if strings.Contains(commandHistory[i], text) {
			return commandHistory[i], i + 1
		}
	}
	
	for i := 0; i < start; i++ {
		if strings.Contains(commandHistory[i], text) {
			return commandHistory[i], i + 1
		}
	}
	
	return "", 0
}

// RunPrompt 运行基于go-prompt的命令行界面
func RunPrompt() {
	// 创建并运行提示器
	p := createPrompt()
	p.Run()
}
