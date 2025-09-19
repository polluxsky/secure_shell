package monitor

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// Start 启动监控进程
func Start() {
	for {
		// 创建子进程
		cmd := exec.Command(os.Args[0], "shell")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Env = append(os.Environ(), "TERM=xterm-256color") // 确保终端支持颜色
		
		err := cmd.Start()
		if err != nil {
			fmt.Printf("Failed to start shell: %v\n", err)
			return
		}
		
		// 等待子进程退出
		err = cmd.Wait()
		
		// 检查退出状态码
		if exitError, ok := err.(*exec.ExitError); ok {
			// 获取退出状态码
			status := exitError.Sys().(syscall.WaitStatus).ExitStatus()
			
			// 如果是正常退出状态码，终止监控循环
			if status == 100 {
				fmt.Println("Shell exited normally.")
				return
			}
			
			fmt.Printf("Shell exited with code %d, restarting...\n", status)
		} else if err != nil {
			fmt.Printf("Shell exited with error: %v, restarting...\n", err)
		} else {
			fmt.Println("Shell exited normally, restarting...")
		}
	}
}
