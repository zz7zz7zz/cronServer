package main

import (
	"cronServer/config"
	"cronServer/models"
	"cronServer/tasks"
	"strconv"

	"github.com/gin-gonic/gin"
)

//gin
//https://www.bilibili.com/video/BV1Rd4y1C7A1/?spm_id_from=333.337.search-card.all.click&vd_source=630aca8d31fad0f6a159cf69cf0dca35

func main() {

	config.InitConfig()
	// database.InitDb()

	// go startTasks()
	// select {}

	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// routers.InitRouters(r)

	task := tasks.NewAsReviewTask(&models.AppReviewRecord{Ver: "2.21.6", Pkg: "1596875621", Platform: "ios"})
	task.Run()

	// task := tasks.NewGpRewiewTask(&models.AppReviewRecord{Ver: "2.21.2", Pkg: "com.inhobichat.hobichat", Platform: "android"})
	// task.Run()

	r.Run(":" + strconv.Itoa(config.GConfig.Server.Port))
}
