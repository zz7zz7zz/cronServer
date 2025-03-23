package tasks

import (
	"cronServer/constant"
	"cronServer/database"
	"cronServer/models"
	"cronServer/utils"
	"cronServer/webhook"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Apple Stroe Review Task
type AsReviewTask struct {
	appReviewRecord *models.AppReviewRecord
}

func NewAsReviewTask(appReviewRecord *models.AppReviewRecord) *AsReviewTask {
	t := &AsReviewTask{
		appReviewRecord: appReviewRecord}
	return t
}

func (t *AsReviewTask) Run() {
	now := time.Now()
	fmt.Println("------Apple start------", now.Format("2006-01-02 15:04:05"))
	version, updateTime, err := scrapeAppStore(t.appReviewRecord.Pkg)
	version = strings.ReplaceAll(version, "Version", "")
	version = strings.TrimSpace(version)
	if err != nil {
		fmt.Println("Apple Error:", err)
	} else {

		// 转换为 UTC 时间
		tUTC := time.Unix(updateTime, 0).UTC()
		fmt.Println("UTC 时间:", tUTC.Format("2006年01月02日 15:04:05"))

		// 转换为本地时区（如 Asia/Shanghai）,会出现这个问题missing Location in call to Time.In
		// loc, err := time.LoadLocation("Asia/Shanghai")
		// if(err != nil){
		// 	fmt.Println("Failed to load location:", err)
		// }
		loc := time.FixedZone("CST", 8*60*60) // 东八区，UTC+8

		tLocal := time.Unix(updateTime, 0).In(loc)
		fmt.Println("本地时间:", tLocal.Format("2006年01月02日 15:04:05"))

		fmt.Println("------Apple version------", version, updateTime)
	}
	version = strings.ToLower(version)
	version = strings.ReplaceAll(version, "version", "")
	version = strings.TrimSpace(version)

	cmpValue := utils.VersionCompare(version, t.appReviewRecord.Ver)
	if cmpValue == 0 {
		fmt.Println("检测到版本审核-成功")
		t.appReviewRecord.ApproveTs = int(updateTime)
		hook := &webhook.ServerWebHook{}
		hook.OnWebHook(t.appReviewRecord)
		database.UpdateTaskStatus(t.appReviewRecord.Platform, t.appReviewRecord.Ver, t.appReviewRecord.Pkg, 3)
		database.UpdateStatus(t.appReviewRecord.Platform, t.appReviewRecord.Ver, t.appReviewRecord.Pkg, 1)
		StopTask(t.appReviewRecord.Ver, t.appReviewRecord.Pkg, t.appReviewRecord.Platform)
	} else {
		if cmpValue == 1 {
			fmt.Println("检测到版本审核-已有更新的版本，当前任务将忽略")
			database.UpdateTaskStatus(t.appReviewRecord.Platform, t.appReviewRecord.Ver, t.appReviewRecord.Pkg, 3)
			database.UpdateStatus(t.appReviewRecord.Platform, t.appReviewRecord.Ver, t.appReviewRecord.Pkg, 3)
			StopTask(t.appReviewRecord.Ver, t.appReviewRecord.Pkg, t.appReviewRecord.Platform)
		} else {
			fmt.Println("检测到版本审核-失败")
		}
	}
	fmt.Println("------Apple end------", t.appReviewRecord.Ver, version)
}

func scrapeAppStore(pkg string) (string, int64, error) {
	url := fmt.Sprintf(constant.AppStoreURL, pkg)
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", 0, err
	}

	// wferr := os.WriteFile("output.txt", []byte(doc.Text()), 0644)
	// if wferr != nil {
	// 	fmt.Println("文件写入失败:", err)
	// 	return "", "", wferr
	// }

	version := doc.Find(".whats-new__latest__version").Text()

	//解析出的数据是这样：Mar 19, 2025
	// time := doc.Find("time[data-test-we-datetime]").Text()
	time := extractUpdateTimeAs(doc.Text())

	return version, time, nil
}

func extractUpdateTimeAs(html string) int64 {
	// 步骤 1: 正则匹配提取时间字符串
	re := regexp.MustCompile(`"releaseTimestamp\\":\\"([^"]+)\\"`)
	matches := re.FindStringSubmatch(html)
	if len(matches) < 2 {
		fmt.Println("错误: 未找到 releaseTimestamp 时间字段")
		return 0
	}
	timeStr := matches[1]

	// 步骤 2: 解析时间字符串
	// 注意布局必须严格对应 "2006-01-02T15:04:05Z"
	t, err := time.Parse(time.RFC3339, timeStr) // RFC3339 布局即 "2006-01-02T15:04:05Z07:00"
	if err != nil {
		fmt.Println("时间解析失败:", err)
		return 0
	}

	// 步骤 3: 转换为秒级时间戳
	timestamp := t.Unix()
	return timestamp
}
