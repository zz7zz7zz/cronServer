package tasks

import (
	"cronServer/constant"
	"cronServer/database"
	"cronServer/models"
	"fmt"

	"github.com/robfig/cron/v3"
)

var GCron *cron.Cron

var GPendingTasks = make(map[string]cron.EntryID)

func StartTasks(appReviewRecord *models.AppReviewRecord, key string) {
	GCron = cron.New(cron.WithSeconds())
	if appReviewRecord.Platform == constant.Android {
		task := NewGpRewiewTask(appReviewRecord)
		id := innerStartTask("10 * * * * * ", task)
		GPendingTasks[key] = id
	} else if appReviewRecord.Platform == constant.Ios {
		task := NewAsReviewTask(appReviewRecord)
		id2 := innerStartTask("10 * * * * * ", task)
		GPendingTasks[key] = id2
	}
}

func innerStartTask(spec string, cmd cron.Job) cron.EntryID {
	id2, err := GCron.AddJob(spec, cmd)
	if err != nil {
		fmt.Println("Error is ", err.Error())
		return -1
	}
	fmt.Println("innerStartTask ID: ", id2)
	GCron.Start()
	return id2
}

func StartTask(ver string, pkg string, platform string) bool {
	key := fmt.Sprintf("%s_%s_%s", platform, ver, pkg)
	_, flag := GPendingTasks[key]
	if !flag {
		StartTasks(&models.AppReviewRecord{Pkg: pkg, Ver: ver, Platform: platform}, key)
		database.Insert(platform, ver, pkg, 0, 1)
	}
	return flag
}

func StopTask(ver string, pkg string, platform string) bool {
	key := fmt.Sprintf("%s_%s_%s", platform, ver, pkg)
	value, flag := GPendingTasks[key]
	if flag {
		delete(GPendingTasks, key)
		GCron.Remove(value)
		database.UpdateTaskStatus(platform, ver, pkg, 3)
	}
	return flag
}

func Ternary(b bool, valueIfTrue, valueIfFalse string) string {
	if b {
		return valueIfTrue
	}
	return valueIfFalse
}
