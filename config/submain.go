package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var GConfig *Config

func InitConfig() {
	// 获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("无法获取当前工作目录: %v", err)
	}

	// 构造 YAML 文件路径
	configPath := filepath.Join(dir, "config.yaml")

	// 打开 YAML 文件
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer file.Close()

	// 创建一个 Config 实例
	GConfig = &Config{}

	// 解析 YAML 文件
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&GConfig); err != nil {
		log.Fatalf("解析 YAML 文件失败: %v", err)
	}

	// 打印解析结果
	fmt.Printf("Database: %+v\n", GConfig.Database)
	fmt.Printf("Webhook: %+v\n", GConfig.Webhook)
	fmt.Printf("Server: %+v\n", GConfig.Server)
}
