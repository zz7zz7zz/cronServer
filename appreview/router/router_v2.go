package router

import (
	"fmt"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/proto"
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
		// 1. 读取请求体
		body, err := c.GetRawData()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		}

		// 2. 解析 Protobuf 数据
		var req models.AppReviewRequest
		if err := proto.Unmarshal(body, &req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid protobuf data"})
			return
		}

		pkg := req.Pkg
		ver := req.Ver
		platform := req.Platform
		status := req.Status

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
		
		// 3. 处理业务逻辑
		response := &models.AppReviewResponse{
			Message:  message,
			Ver:      ver,
			Pkg:      pkg,
			Platform: platform,
			status:"start",
			Key:fmt.Sprintf("%s_%s_%s", platform, ver, pkg)
		}

		// 4. 序列化响应
		protoData, err := proto.Marshal(response)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to marshal response"})
			return
		}

		// 5. 设置响应头并返回
		c.Data(http.StatusOK, "application/x-protobuf", protoData)
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
