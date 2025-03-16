package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
