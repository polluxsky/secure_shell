package shell

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"

	gterm "golang.org/x/term"
)

// 允许的命令列表
var allowedCommands = map[string]bool{
	"ls":      true,
	"ll":      true,
	"cd":      true,
	"cat":     true,
	"pwd":     true,
	"vim":     true,
	"mkdir":   true,
	"history": true,
	"exit":    true,
	"help":    true,
}

// 存储命令历史
var commandHistory []string

// 正常退出状态码
const normalExitCode = 100

// 执行命令并输出结果
func executeCommand(cmd string) bool {
	// 添加到命令历史
	if cmd != "" {
		commandHistory = append(commandHistory, cmd)
	}

	// 分割命令和参数
	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return true
	}
	cmdName := parts[0]
	
	// 检查命令是否被允许
	if !allowedCommands[cmdName] {
		fmt.Printf("%s: command not found\n", cmdName)
		return true
	}

	// 处理特殊命令
	switch cmdName {
	case "cd":
		return handleCdCommand(parts)
	case "pwd":
		return handlePwdCommand()
	case "history":
		return handleHistoryCommand()
	case "exit":
		return handleExitCommand()
	case "ls", "ll":
		return handleLsCommand(cmd, parts)
	case "help":
		return handleHelpCommand()
	default:
		return handleOtherCommand(cmdName, cmd, parts)
	}
}

// 处理cd命令
func handleCdCommand(parts []string) bool {
	var dir string
	if len(parts) < 2 {
		// 如果没有参数，切换到用户主目录
		u, err := user.Current()
		if err != nil {
			fmt.Printf("cd: failed to get home directory: %v\n", err)
			return true
		}
		dir = u.HomeDir
	} else {
		dir = parts[1]
		// 处理波浪号
		if strings.HasPrefix(dir, "~") {
			u, err := user.Current()
			if err == nil {
				dir = u.HomeDir + dir[1:]
			}
		}
	}

	// 执行目录切换
	err := os.Chdir(dir)
	if err != nil {
		fmt.Printf("cd: %v\n", err)
	}
	return true
}

// 处理pwd命令
func handlePwdCommand() bool {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("pwd: %v\n", err)
	} else {
		fmt.Println(dir)
	}
	return true
}

// 处理history命令
func handleHistoryCommand() bool {
	for i, cmd := range commandHistory {
		fmt.Printf("%d  %s\n", i+1, cmd)
	}
	return true
}

// 处理help命令
func handleHelpCommand() bool {
	// 命令说明映射
	commandDescriptions := map[string]string{
		"ls":      "列出目录内容",
		"ll":      "详细列出目录内容",
		"cd":      "切换目录",
		"cat":     "查看文件内容",
		"pwd":     "显示当前目录",
		"vim":     "文本编辑器",
		"mkdir":   "创建目录",
		"history": "查看命令历史",
		"help":    "显示帮助信息",
	}

	fmt.Println("支持的命令列表：")
	fmt.Println("================")
	
	// 遍历允许的命令并显示说明（exit除外）
	for cmd := range allowedCommands {
		if cmd != "exit" {
			// 使用制表符对齐命令名和说明
			fmt.Printf("%s\t\t-- %s\n", cmd, commandDescriptions[cmd])
		}
	}
	
	fmt.Println("================")
	return true
}

// 处理exit命令
func handleExitCommand() bool {

	fmt.Print("Enter password to exit: ")
	// 切换终端到原始模式以隐藏输入
	fd := int(os.Stdin.Fd())
	oldState, err := gterm.MakeRaw(fd)
	if err != nil {
		fmt.Println("Failed to set terminal to raw mode")
		return true
	}
	defer gterm.Restore(fd, oldState)
	defer fmt.Println() // 在函数结束时添加换行
	
	// 读取密码（隐藏输入）
	password, err := gterm.ReadPassword(fd)
	if err != nil {
		fmt.Println("Failed to read password")
		return true
	}
	
	// 验证密码
	if string(password) == "nkwya" {
		fmt.Println()
		fmt.Println("Exiting secure shell...\n") // 添加换行符确保移至新行首列
		return false // 返回false表示要退出
	} else {
		fmt.Println()
		fmt.Println("Incorrect password. Exit aborted.")
		return true
	}
}

// 处理ls和ll命令
func handleLsCommand(cmd string, parts []string) bool {
	// 确保颜色显示
	if !strings.Contains(cmd, "--color") {
		cmd += " --color=always"
	}

	// 处理ll命令
	if parts[0] == "ll" {
		cmd = "ls -l --color=always " + strings.Join(parts[1:], " ")
	}

	// 执行命令
	c := exec.Command("sh", "-c", cmd)
	c.Env = append(os.Environ(), "TERM=xterm-256color") // 确保终端支持颜色
	output, err := c.CombinedOutput()
	
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	if len(output) > 0 {
		fmt.Print(string(output))
	}
	
	return true
}

// 处理其他命令
func handleOtherCommand(cmdName, cmd string, parts []string) bool {
	var output []byte
	var err error
	
	if cmdName == "vim" {
		// vim需要交互式运行
		c := exec.Command("vim", parts[1:]...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Env = os.Environ()
		err = c.Run()
	} else {
		// 其他命令
		c := exec.Command("sh", "-c", cmd)
		c.Env = os.Environ()
		output, err = c.CombinedOutput()
	}

	if err != nil {
		fmt.Printf("%v\n", err)
	}
	if len(output) > 0 {
		fmt.Print(string(output))
	}

	return true
}
