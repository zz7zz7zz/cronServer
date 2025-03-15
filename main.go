package main

import (
	"fmt"
	"net/http"

	"cronServer/tasks"

	"github.com/robfig/cron/v3"

	"github.com/gin-gonic/gin"
)

//gin
//https://www.bilibili.com/video/BV1Rd4y1C7A1/?spm_id_from=333.337.search-card.all.click&vd_source=630aca8d31fad0f6a159cf69cf0dca35

var cr *cron.Cron

var taskMap = make(map[string]cron.EntryID)

func main() {
	// go startTasks()
	// select {}

	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	//查询审核状态
	r.GET("/appreview/list", func(c *gin.Context) {
		ver := c.Query("ver")
		pkg := c.Query("pkg")
		platform := c.Query("platform")
		c.JSON(http.StatusOK, gin.H{
			"ver":      ver,
			"pkg":      pkg,
			"platform": platform,
			"status":   "审核通过",
		})
	})

	//
	r.GET("/appreview/start", func(c *gin.Context) {
		pkg := c.Query("pkg")
		ver := c.Query("ver")
		platform := c.Query("platform")

		key := fmt.Sprintf("%s_%s_%s", platform, ver, pkg)
		_, flag := taskMap[key]
		if !flag {
			startTasks()
		}
		c.JSON(http.StatusOK, gin.H{
			"ver":      ver,
			"pkg":      pkg,
			"platform": platform,
			"key":      key,
			"status":   "审核通过",
			"cron":     ternary(flag, "已经在任务中 ", "首次添加"),
		})
	})

	//
	r.GET("/appreview/stop", func(c *gin.Context) {
		ver := c.Query("ver")
		pkg := c.Query("pkg")
		platform := c.Query("platform")
		key := fmt.Sprintf("%s_%s_%s", platform, ver, pkg)
		value, flag := taskMap[key]
		if flag {
			delete(taskMap, key)
			cr.Remove(value)
		}
		c.JSON(http.StatusOK, gin.H{
			"ver":      ver,
			"pkg":      pkg,
			"platform": platform,
			"status":   "审核通过",
			"cron":     ternary(flag, "移除成功 ", "未添加任务"),
		})
	})

	r.Run()
}

func startTasks() {

	cr = cron.New(cron.WithSeconds())

	task := tasks.NewGoogleRewiewTask()
	id := startTaskItem("* * * * * * ", task)
	taskMap["android_2.20.1_com.inhobichat.hobichat"] = id

	task2 := tasks.NewAppleReviewTask()
	id2 := startTaskItem("* * * * * * ", task2)
	taskMap["ios_2.20.1_com.inhobichat.hobichat"] = id2
}

func startTaskItem(spec string, cmd cron.Job) cron.EntryID {
	id2, err := cr.AddJob(spec, cmd)
	if err != nil {
		fmt.Println("Error is ", err.Error())
		return -1
	}
	fmt.Println("ID is ", id2)
	cr.Start()
	return id2
}

func ternary(b bool, valueIfTrue, valueIfFalse string) string {
	if b {
		return valueIfTrue
	}
	return valueIfFalse
}
