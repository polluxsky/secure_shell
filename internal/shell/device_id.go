package shell

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"path/filepath"
	"syscall"
)

const (
	// deviceIDDir 存储设备唯一标识的目录
	deviceIDDir = "/etc/secureshell"
	// deviceIDFile 设备唯一标识文件名
	deviceIDFile = "device_id"
)

// generateDeviceID 生成设备唯一标识文件
// 第一次启动时生成，如果文件已存在则不再生成
// 文件权限设置为只有root用户可读可写，确保应用本身无法访问
func generateDeviceID() {
	// 检查目录是否存在，如果不存在则创建
	if _, err := os.Stat(deviceIDDir); os.IsNotExist(err) {
		// 创建目录，权限设置为0700（只有root用户可读写执行）
		if err := os.MkdirAll(deviceIDDir, 0700); err != nil {
			// 如果创建目录失败，记录错误但继续执行
			// 注意：这里不输出错误信息，以避免暴露该功能
			return
		}
	}

	// 构建完整的文件路径
	deviceIDFilePath := filepath.Join(deviceIDDir, deviceIDFile)

	// 检查文件是否已存在
	if _, err := os.Stat(deviceIDFilePath); err == nil {
		// 文件已存在，不需要重新生成
		return
	}

	// 生成随机的设备唯一标识（32字节）
	id := make([]byte, 32)
	if _, err := rand.Read(id); err != nil {
		// 生成随机数失败，记录错误但继续执行
		return
	}

	// 将二进制数据转换为十六进制字符串
	idHex := hex.EncodeToString(id)

	// 创建文件并写入设备标识，权限设置为0600（只有root用户可读写）
	file, err := os.OpenFile(deviceIDFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		// 创建文件失败，记录错误但继续执行
		return
	}
	defer file.Close()

	// 写入设备标识
	if _, err := file.WriteString(idHex); err != nil {
		// 写入失败，记录错误但继续执行
		return
	}

	// 设置文件所有权为root:root
	// 注意：这只有在应用以root权限运行时才会生效
	syscall.Chown(deviceIDFilePath, 0, 0)
}