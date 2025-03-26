package utils

import (
	"strconv"
	"strings"
)

// VersionCompare 比较两个版本号字符串
// 返回 -1 表示 v1 < v2，0 表示相等，1 表示 v1 > v2
func VersionCompare(v1, v2 string) int {
	// 分割版本号为分段
	parts1 := strings.Split(v1, ".")
	parts2 := strings.Split(v2, ".")

	// 获取最大分段数
	maxLength := max(len(parts1), len(parts2))

	for i := 0; i < maxLength; i++ {
		// 获取当前分段数值（缺省为0）
		num1 := getPartValue(parts1, i)
		num2 := getPartValue(parts2, i)

		// 数值比较
		if num1 < num2 {
			return -1
		} else if num1 > num2 {
			return 1
		}
	}
	return 0
}

// 获取分段数值（带越界保护）
func getPartValue(parts []string, index int) int {
	if index >= len(parts) {
		return 0
	}

	// 字符串转数值（忽略错误，非法字符按0处理）
	num, _ := strconv.Atoi(parts[index])
	return num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
