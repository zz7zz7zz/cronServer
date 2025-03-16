package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"cronServer/models"
)

var DB *gorm.DB

func InitDb() {
	dsn := "melo:jpYU7CJ5t36e@tcp(172.18.26.11:3306)/melo_appreview?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print("Init DB Error", err)
		return
	}
	fmt.Println("Init DB Success")
}

func GetList(platform string, ver string, pkg string, status int) []models.AppReviewRecord {
	var ret []models.AppReviewRecord
	query := DB.Model(&models.AppReviewRecord{})

	// 动态拼接条件
	if platform != "" {
		query = query.Where("platform = ?", platform)
	}
	if ver != "" {
		query = query.Where("ver = ?", ver)
	}
	if pkg != "" {
		query = query.Where("pkg = ?", pkg)
	}
	if status != 0 { // 假设 status=0 表示“全部状态”
		query = query.Where("status = ?", status)
	}

	query.Find(&ret)
	return ret
}
