package shell

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"unicode"

	gterm "golang.org/x/term"
	pkgterm "secure-shell/pkg/term"
)

// 读取带Tab补全的输入
func readInputWithCompletion() (string, error) {
	fd := int(os.Stdin.Fd())
	oldState, err := gterm.MakeRaw(fd)
	if err != nil {
		return "", err
	}
	defer gterm.Restore(fd, oldState)

	input := make([]rune, 0)
	cursorPos := 0

	for {
		// 读取单个字符
		buf := make([]byte, 3)
		n, err := os.Stdin.Read(buf)
		if err != nil {
			return string(input), err
		}

		// 处理输入
		if n == 1 {
			switch buf[0] {
			case '\n', '\r': // 回车
				fmt.Println()
				return string(input), nil
			case '\b', 127: // 退格
				if len(input) > 0 && cursorPos > 0 {
					input = append(input[:cursorPos-1], input[cursorPos:]...)
					cursorPos--
					redrawInput(input, cursorPos)
				}
			case '\t': // Tab补全
				currentInput := string(input)
				completions := getCompletions(currentInput)
				
				if len(completions) == 1 {
					// 唯一补全
					parts := strings.Fields(currentInput)
					var newInput string
					
					if len(parts) == 0 {
						newInput = completions[0] + " "
					} else if len(parts) == 1 && currentInput == parts[0] {
						newInput = completions[0] + " "
					} else {
						newInput = strings.Join(parts[:len(parts)-1], " ") + " " + completions[0]
					}
					
					input = []rune(newInput)
					cursorPos = len(input)
					redrawInput(input, cursorPos)
				} else if len(completions) > 1 {
					// 多个补全选项
					fmt.Println()
					for i, comp := range completions {
						if i > 0 && i%4 == 0 { // 每行显示4个选项
							fmt.Println()
						}
						fmt.Printf("%-20s", comp)
					}
					fmt.Println()
					showPrompt()
					fmt.Print(string(input))
					// 移动光标到正确位置
					moveLeft := len(input) - cursorPos
					if moveLeft > 0 {
						pkgterm.MoveCursor(-moveLeft)
					}
				}
			default:
				// 普通字符
				if r := rune(buf[0]); unicode.IsPrint(r) {
					// 在光标位置插入字符
					input = append(input[:cursorPos], append([]rune{r}, input[cursorPos:]...)...)
					cursorPos++
					redrawInput(input, cursorPos)
				}
			}
		} else if n == 3 && buf[0] == 27 && buf[1] == 91 {
			// 处理方向键
			switch buf[2] {
			case 'C': // 右箭头
				if cursorPos < len(input) {
					cursorPos++
					pkgterm.MoveCursor(1)
				}
			case 'D': // 左箭头
				if cursorPos > 0 {
					cursorPos--
					pkgterm.MoveCursor(-1)
				}
			}
		}
	}
}

// 重绘输入内容
func redrawInput(input []rune, cursorPos int) {
	// 改进的光标控制逻辑
	// 1. 移动到行首
	fmt.Print("\r")
	// 2. 清除整行（从行首到行尾）
	fmt.Print("\033[2K")
	// 3. 重新显示输入内容
	fmt.Print(string(input))
	// 4. 移动光标到正确位置
	if cursorPos >= 0 && cursorPos <= len(input) {
		// 计算需要向左移动的距离
		moveLeft := len(input) - cursorPos
		if moveLeft > 0 {
			fmt.Printf("\033[%dD", moveLeft)
		}
	}
	// 确保所有输出都被刷新到终端
	os.Stdout.Sync()
}

// 获取补全选项
func getCompletions(input string) []string {
	// 分割输入为命令和参数
	parts := strings.Fields(input)
	var prefix string
	var completions []string
	
	if len(parts) == 0 {
		// 没有输入，返回所有命令
		for cmd := range allowedCommands {
			completions = append(completions, cmd)
		}
		return completions
	} else if len(parts) == 1 && strings.TrimSpace(input) == parts[0] {
		// 只有命令部分，可能需要补全命令
		prefix = parts[0]
		completions = getCommandCompletions(prefix)
		
		// 如果没有命令补全，尝试路径补全
		if len(completions) == 0 {
			completions = getPathCompletions(prefix)
		}
	} else {
		// 有参数，补全最后一个参数（路径）
		prefix = parts[len(parts)-1]
		completions = getPathCompletions(prefix)
	}
	
	return completions
}

// 获取可能的命令补全
func getCommandCompletions(prefix string) []string {
	var completions []string
	for cmd := range allowedCommands {
		if strings.HasPrefix(cmd, prefix) {
			completions = append(completions, cmd)
		}
	}
	return completions
}

// 获取可能的路径补全
func getPathCompletions(prefix string) []string {
	var completions []string
	
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
		return completions
	}
	
	// 查找匹配的文件/目录
	for _, f := range files {
		name := f.Name()
		if strings.HasPrefix(name, file) {
			fullPath := filepath.Join(dir, name)
			
			// 如果是目录，添加斜杠
			if f.IsDir() {
				name += "/"
			}
			
			// 如果前缀包含目录部分，保留目录结构
			if dir != "." {
				relPath, err := filepath.Rel(".", fullPath)
				if err == nil {
					name = relPath
					if f.IsDir() && !strings.HasSuffix(name, "/") {
						name += "/"
					}
				}
			}
			
			// 替换回波浪号表示法（如果适用）
			if strings.HasPrefix(prefix, "~") {
				u, err := user.Current()
				if err == nil && strings.HasPrefix(name, u.HomeDir) {
					name = "~" + name[len(u.HomeDir):]
				}
			}
			
			completions = append(completions, name)
		}
	}
	
	return completions
}
