package main

import (
	"cronServer/config"

	"github.com/gin-gonic/gin"
)

//gin
//https://www.bilibili.com/video/BV1Rd4y1C7A1/?spm_id_from=333.337.search-card.all.click&vd_source=630aca8d31fad0f6a159cf69cf0dca35

func main() {

	config.InitConfig()

	// hook := webhook.ServerWebHook{}
	// hook.OnWebHook()

	// database.InitDb()

	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// routers.InitRouters(r)

	r.Run()
}
