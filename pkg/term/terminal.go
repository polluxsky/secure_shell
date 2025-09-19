package term

import (
	"fmt"
	"os"
)

// MoveCursor 相对移动光标
func MoveCursor(relative int) {
	if relative > 0 {
		fmt.Printf("\033[%dC", relative)
	} else if relative < 0 {
		fmt.Printf("\033[%dD", -relative)
	}
	os.Stdout.Sync()
}

// ClearLine 清除当前行
func ClearLine() {
	fmt.Printf("\r\033[K")
	os.Stdout.Sync()
}
