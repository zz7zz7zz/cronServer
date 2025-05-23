package main

import (
	"strconv"

	"open.com/cronServer/appreview/config"
	"open.com/cronServer/appreview/database"
	"open.com/cronServer/appreview/tasks"
	"open.com/cronServer/cmd"
	"open.com/cronServer/routers"

	"github.com/gin-gonic/gin"
)

//gin
//https://www.bilibili.com/video/BV1Rd4y1C7A1/?spm_id_from=333.337.search-card.all.click&vd_source=630aca8d31fad0f6a159cf69cf0dca35

func main() {

	main_cmd()
	// main_appreview()
}

func main_cmd() {
	cmd.Execute()
}

func main_appreview() {

	config.InitConfig()
	database.InitDb()

	tasks.RecoverAppReviewTasks()

	//-------------------- Test start --------------------
	// task := tasks.NewAsReviewTask(&models.AppReviewRecord{Ver: "2.21.6", Pkg: "1596875621", Platform: constant.Ios})
	// task.Run()
	// task := tasks.NewGpRewiewTask(&models.AppReviewRecord{Ver: "2.21.2", Pkg: "com.inhobichat.hobichat", Platform: constant.Android})
	// task.Run()
	//-------------------- Test end --------------------

	r := gin.Default()
	// gin.SetMode(gin.ReleaseMode)
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	routers.InitRouters(r)
	r.Run(":" + strconv.Itoa(config.G_Config.Server.Port))

	// 不退出
	// go startTasks()
	// select {}
}
