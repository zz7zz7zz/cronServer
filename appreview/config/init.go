package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var G_Config *Config

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
	G_Config = &Config{}

	// 解析 YAML 文件
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&G_Config); err != nil {
		log.Fatalf("解析 YAML 文件失败: %v", err)
	}

	// 打印解析结果
	fmt.Printf("Database: %+v\n", G_Config.Database)
	fmt.Printf("Webhook: %+v\n", G_Config.Webhook)
	fmt.Printf("Server: %+v\n", G_Config.Server)
}
