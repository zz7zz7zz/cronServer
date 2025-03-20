package database

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"cronServer/config"
	"cronServer/models"
)

var DB *gorm.DB

func InitDb() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.GConfig.Database.User, config.GConfig.Database.Password, config.GConfig.Database.Host, config.GConfig.Database.Port, config.GConfig.Database.Name)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print("Init DB Error", err)
		return
	}
	fmt.Println("Init DB Success")
}

func GetList(platform string, ver string, pkg string, status int, taskstatus int) []models.AppReviewRecord {
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
	if taskstatus != 0 { // 假设 taskstatus=0 表示“全部状态”
		query = query.Where("task_status = ?", taskstatus)
	}
	// 添加倒序排序（按时间戳或ID倒序）
	// query = query.Order("time_stamp DESC") // 或 "id DESC"
	// 修改排序逻辑：先按 ver 倒序，再按 time_stamp 倒序
	query = query.Order("ver DESC, time_stamp DESC")

	query.Find(&ret)
	return ret
}

func Insert(platform string, ver string, pkg string, status int, taskstatus int) error {
	// 1. 检查记录是否已存在
	var existingRecord models.AppReviewRecord
	result := DB.Where("platform = ? AND ver = ? AND pkg = ?", platform, ver, pkg).First(&existingRecord)

	// 2. 如果已存在，更新 taskstatus 字段
	if result.Error == nil {
		existingRecord.TaskStatus = taskstatus
		if err := DB.Save(&existingRecord).Error; err != nil {
			return fmt.Errorf("更新失败: %v", err)
		}
		return nil
	}

	// 3. 如果是"未找到记录"错误，继续插入新数据
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 创建新记录
		newRecord := models.AppReviewRecord{
			Platform:   platform,
			Ver:        ver,
			Pkg:        pkg,
			Status:     status,
			TimeStamp:  int(time.Now().Unix()), // 添加时间戳（根据字段类型调整）
			TaskStatus: taskstatus,
		}

		// 插入数据库
		if err := DB.Create(&newRecord).Error; err != nil {
			return fmt.Errorf("插入失败: %v", err)
		}
		return nil
	}

	// 4. 其他数据库错误（如连接问题）
	return fmt.Errorf("查询失败: %v", result.Error)
}

//若未创建联合唯一索引，OnConflict 将无法触发，导致重复插入
// func Insert(platform string, ver string, pkg string, status int, taskStatus int) error {
// 	// 创建记录对象（包含时间戳）
// 	newRecord := models.AppReviewRecord{
// 		Platform:   platform,
// 		Ver:        ver,
// 		Pkg:        pkg,
// 		Status:     status,
// 		TimeStamp:  int(time.Now().Unix()),
// 		TaskStatus: taskStatus,
// 		Channel:    "", // 根据需求初始化或从参数传入
// 	}

// 	// 执行 Upsert 操作
// 	result := DB.Clauses(clause.OnConflict{
// 		Columns: []clause.Column{
// 			{Name: "platform"},
// 			{Name: "ver"},
// 			{Name: "pkg"}, // 依赖这三个字段的唯一索引
// 		},
// 		DoUpdates: clause.AssignmentColumns([]string{"task_status"}), // 冲突时更新 task_status
// 	}).Create(&newRecord)

// 	if result.Error != nil {
// 		return fmt.Errorf("操作失败: %v", result.Error)
// 	}

// 	// 可选：返回插入/更新的 ID
// 	fmt.Printf("记录 ID: %d\n", newRecord.ID)
// 	return nil
// }

func UpdateTaskStatus(platform string, ver string, pkg string, taskstatus int) error {
	if platform == "" || ver == "" || pkg == "" {
		return fmt.Errorf("platform/ver/pkg 参数不可为空")
	}
	// 查找符合条件的记录
	result := DB.Model(&models.AppReviewRecord{}).
		Where("platform = ? AND ver = ? AND pkg = ?", platform, ver, pkg).
		Update("task_status", taskstatus)

	// 检查是否有错误
	if result.Error != nil {
		return fmt.Errorf("更新失败: %v", result.Error)
	}

	// 检查是否有记录被更新
	if result.RowsAffected == 0 {
		return errors.New("没有符合条件的记录被更新")
	}

	return nil
}

func UpdateStatus(platform string, ver string, pkg string, status int) error {
	if platform == "" || ver == "" || pkg == "" {
		return fmt.Errorf("platform/ver/pkg 参数不可为空")
	}
	// 查找符合条件的记录
	result := DB.Model(&models.AppReviewRecord{}).
		Where("platform = ? AND ver = ? AND pkg = ?", platform, ver, pkg).
		Update("status", status)

	// 检查是否有错误
	if result.Error != nil {
		return fmt.Errorf("更新失败: %v", result.Error)
	}

	// 检查是否有记录被更新
	if result.RowsAffected == 0 {
		return errors.New("没有符合条件的记录被更新")
	}

	return nil
}
