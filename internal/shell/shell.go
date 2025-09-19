package shell

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// SetupSignalHandler 设置信号处理器
func SetupSignalHandler() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	
	go func() {
		for sig := range sigChan {
			fmt.Printf("\nReceived signal %v, ignoring...\n", sig)
		}
	}()
}

// Start 启动shell
func Start() {
	// 显示banner
	DisplayBanner()
	
	// 设置信号处理器
	SetupSignalHandler()
	
	// 运行基于go-prompt的命令行界面
	RunPrompt()
}
