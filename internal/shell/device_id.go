package shell

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

const (
	// deviceIDFile 设备唯一标识文件名
	deviceIDFile = "device_id"
)

// generateDeviceID 生成设备唯一标识文件
// 第一次启动时生成，如果文件已存在则不再生成
func generateDeviceID() {
	// 获取用户主目录作为存储位置，避免需要root权限
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("警告: 无法获取用户主目录，无法生成设备唯一标识: %v\n", err)
		return
	}
	
	// 使用用户可访问的目录存储设备ID
	deviceIDDir := filepath.Join(homeDir, ".secureshell")

	// 检查目录是否存在，如果不存在则创建
	if _, err := os.Stat(deviceIDDir); os.IsNotExist(err) {
		// 创建目录，权限设置为0700（只有当前用户可读写执行）
		if err := os.MkdirAll(deviceIDDir, 0700); err != nil {
			fmt.Printf("警告: 无法创建设备ID目录 %s: %v\n", deviceIDDir, err)
			return
		}
	}

	// 构建完整的文件路径
	deviceIDFilePath := filepath.Join(deviceIDDir, deviceIDFile)

	// 检查文件是否已存在
	if _, err := os.Stat(deviceIDFilePath); err == nil {
		// 文件已存在，检查内容是否为空
		content, err := os.ReadFile(deviceIDFilePath)
		if err != nil {
			fmt.Printf("警告: 无法读取设备ID文件 %s: %v\n", deviceIDFilePath, err)
		} else if len(content) == 0 {
			fmt.Printf("注意: 设备ID文件存在但内容为空，将重新生成\n")
		} else {
			fmt.Printf("设备ID文件已存在，跳过生成\n")
			return
		}
	}

	// 生成随机的设备唯一标识（32字节）
	id := make([]byte, 32)
	if _, err := rand.Read(id); err != nil {
		fmt.Printf("警告: 无法生成随机设备ID: %v\n", err)
		return
	}

	// 将二进制数据转换为十六进制字符串
	idHex := hex.EncodeToString(id)

	// 创建文件并写入设备标识，权限设置为0600（只有当前用户可读写）
	file, err := os.OpenFile(deviceIDFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Printf("警告: 无法创建设备ID文件 %s: %v\n", deviceIDFilePath, err)
		return
	}
	defer file.Close()

	// 写入设备标识（添加换行符确保文件格式正确）
	if _, err := file.WriteString(idHex + "\n"); err != nil {
		fmt.Printf("警告: 无法写入设备ID到文件 %s: %v\n", deviceIDFilePath, err)
		return
	}

	// 显示设备ID文件创建成功的提示
	fmt.Printf("设备唯一标识文件已成功创建！\n")
}