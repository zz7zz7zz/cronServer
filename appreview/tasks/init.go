package tasks

import (
	"fmt"
	"time"

	"open.com/cronServer/appreview/constant"
	"open.com/cronServer/appreview/database"
	"open.com/cronServer/appreview/models"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

var G_Cron *cron.Cron

var GPendingTasks = make(map[string]cron.EntryID)

func RecoverAppReviewTasks() {
	appleReviewRecords := database.GetList("", "", "", constant.ReviewPending, constant.TaskRunning)
	//自动开启以下任务
	for _, record := range appleReviewRecords {
		if record.TaskStatus == 1 {
			key := fmt.Sprintf("%s_%s_%s", record.Platform, record.Ver, record.Pkg)
			_, flag := GPendingTasks[key]
			if !flag {
				fmt.Println("自动开启任务 ", key)
				StartTasks(&record, key)
			}
		}
	}
}

func StartTasks(appReviewRecord *models.AppReviewRecord, key string) {
	G_Cron = cron.New(cron.WithSeconds())
	if appReviewRecord.Platform == constant.Android {
		task := NewGooglePlayRewiewTask(appReviewRecord)
		id := innerStartTask("10 * * * * * ", task)
		if gin.Mode() == gin.ReleaseMode {
			id = innerStartTask("0 */10 * * * *", task)
		}
		GPendingTasks[key] = id
	} else if appReviewRecord.Platform == constant.Ios {
		task := NewAppstoreReviewTask(appReviewRecord)
		id2 := innerStartTask("10 * * * * * ", task)
		if gin.Mode() == gin.ReleaseMode {
			id2 = innerStartTask("0 */10 * * * *", task)
		}
		GPendingTasks[key] = id2
	}
}

func innerStartTask(spec string, cmd cron.Job) cron.EntryID {
	id2, err := G_Cron.AddJob(spec, cmd)
	if err != nil {
		fmt.Println("Error is ", err.Error())
		return -1
	}
	fmt.Println("innerStartTask ID: ", id2)
	G_Cron.Start()
	return id2
}

func StartTask(appReviewRecord *models.AppReviewRecord) (constant.ReviewStatus, constant.TaskStatus, error) {
	platform := appReviewRecord.Platform
	ver := appReviewRecord.Ver
	pkg := appReviewRecord.Pkg
	key := fmt.Sprintf("%s_%s_%s", platform, ver, pkg)
	_, flag := GPendingTasks[key]
	if !flag {
		StartTasks(&models.AppReviewRecord{Pkg: pkg, Ver: ver, Platform: platform, TaskCreateTs: int(time.Now().Unix())}, key)
		return database.Insert(appReviewRecord)
	}
	return constant.ReviewPending, constant.TaskNotStart, nil
}

func StopTask(appReviewRecord *models.AppReviewRecord) bool {
	platform := appReviewRecord.Platform
	ver := appReviewRecord.Ver
	pkg := appReviewRecord.Pkg
	key := fmt.Sprintf("%s_%s_%s", platform, ver, pkg)
	value, flag := GPendingTasks[key]
	if flag {
		delete(GPendingTasks, key)
		G_Cron.Remove(value)
		database.UpdateTaskStatus(appReviewRecord)
	}
	return flag
}

func Ternary(b bool, valueIfTrue, valueIfFalse string) string {
	if b {
		return valueIfTrue
	}
	return valueIfFalse
}
