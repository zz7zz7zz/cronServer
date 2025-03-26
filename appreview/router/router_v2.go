package router

import (
	"fmt"
	"net/http"
	"strconv"

	"open.com/cronServer/appreview/constant"
	"open.com/cronServer/appreview/database"
	"open.com/cronServer/appreview/models"
	"open.com/cronServer/appreview/tasks"
	"open.com/cronServer/appreview/utils"

	"github.com/gin-gonic/gin"
)

func InitAppreviewV2(group *gin.RouterGroup) {

	appreview := group.Group("/appreview/v2")

	//查询审核状态
	appreview.POST("/list", func(c *gin.Context) {
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
		appleReviewRecords := database.GetList(platform, ver, pkg, constant.ReviewStatus(status), constant.TaskNotStart)
		c.JSON(http.StatusOK, appleReviewRecords)
	})

	//
	appreview.POST("/start", func(c *gin.Context) {
		pkg := c.Query("pkg")
		ver := c.Query("ver")
		platform := c.Query("platform")

		appReviewRecord, err := database.GetMaxVersionRecord(pkg, platform)
		maxVer := ""
		if err == nil {
			maxVer = appReviewRecord.Ver
		}

		in := &models.AppReviewRecord{
			Ver:        ver,
			Pkg:        pkg,
			Platform:   platform,
			Status:     int(constant.ReviewPending),
			TaskStatus: int(constant.TaskRunning),
		}
		status, taskStatus, err := tasks.StartTask(in)
		fmt.Print("status taskStatus err ", status, taskStatus, err, maxVer)

		message := tasks.Ternary(taskStatus == constant.TaskRunning, "已存在相同任务 ", "启动-定时任务成功")
		if maxVer != "" {
			cmpValue := utils.VersionCompare(appReviewRecord.Ver, ver)
			if cmpValue == 1 {
				message = message + fmt.Sprintf("（已存在审核通过的更高版本%s，但是仍然为你执行相应的任务）", maxVer)
			} else if cmpValue == 0 {
				message = message + "（记录显示该版本已审核通过，但是仍然为你执行相应的任务）"
			}

		}

		c.JSON(http.StatusOK, gin.H{
			"ver":      ver,
			"pkg":      pkg,
			"platform": platform,
			"key":      fmt.Sprintf("%s_%s_%s", platform, ver, pkg),
			"status":   "start",
			"message":  message,
		})
	})

	//
	appreview.POST("/stop", func(c *gin.Context) {
		ver := c.Query("ver")
		pkg := c.Query("pkg")
		platform := c.Query("platform")

		appReviewRecord := &models.AppReviewRecord{
			Ver:        ver,
			Pkg:        pkg,
			Platform:   platform,
			TaskStatus: int(constant.TaskStop),
		}

		flag := tasks.StopTask(appReviewRecord)
		c.JSON(http.StatusOK, gin.H{
			"ver":      ver,
			"pkg":      pkg,
			"platform": platform,
			"status":   "stop",
			"message":  tasks.Ternary(flag, "停止-定时任务成功 ", "没有对应任务"),
		})
	})
}
