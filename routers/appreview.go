package routers

import (
	"cronServer/database"
	"cronServer/tasks"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitAppreview(group *gin.RouterGroup) {

	appleReviewRecords := database.GetList("", "", "", 0, 1)
	//自动开启以下任务
	for _, record := range appleReviewRecords {
		if record.TaskStatus == 1 {
			key := fmt.Sprintf("%s_%s_%s", record.Platform, record.Ver, record.Pkg)
			_, flag := tasks.GPendingTasks[key]
			if !flag {
				fmt.Println("自动开启任务 ", key)
				tasks.StartTasks(&record, key)
			}
		}
	}

	appreview := group.Group("/appreview")

	//查询审核状态
	appreview.GET("/list", func(c *gin.Context) {
		ver := c.Query("ver")
		pkg := c.Query("pkg")
		platform := c.Query("platform")
		statusStr := c.DefaultQuery("status", "0")
		status, err := strconv.Atoi(statusStr)
		if err != nil {
			// 参数无效时的处理逻辑（如返回错误响应）
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status parameter"})
			return
		}
		appleReviewRecords := database.GetList(platform, ver, pkg, status, 0)
		c.JSON(http.StatusOK, appleReviewRecords)
	})

	//
	appreview.GET("/start", func(c *gin.Context) {
		pkg := c.Query("pkg")
		ver := c.Query("ver")
		platform := c.Query("platform")

		flag := tasks.StartTask(ver, pkg, platform)
		c.JSON(http.StatusOK, gin.H{
			"ver":      ver,
			"pkg":      pkg,
			"platform": platform,
			"key":      fmt.Sprintf("%s_%s_%s", platform, ver, pkg),
			"status":   "start",
			"cron":     tasks.Ternary(flag, "已存在相同任务 ", "启动-定时任务成功"),
		})
	})

	//
	appreview.GET("/stop", func(c *gin.Context) {
		ver := c.Query("ver")
		pkg := c.Query("pkg")
		platform := c.Query("platform")

		flag := tasks.StopTask(ver, pkg, platform)
		c.JSON(http.StatusOK, gin.H{
			"ver":      ver,
			"pkg":      pkg,
			"platform": platform,
			"status":   "stop",
			"cron":     tasks.Ternary(flag, "停止-定时任务成功 ", "没有对应任务"),
		})
	})
}
