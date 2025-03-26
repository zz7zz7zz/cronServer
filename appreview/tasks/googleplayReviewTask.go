package tasks

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"open.com/cronServer/appreview/constant"
	"open.com/cronServer/appreview/database"
	"open.com/cronServer/appreview/models"
	"open.com/cronServer/appreview/utils"
	"open.com/cronServer/appreview/webhook"

	"github.com/PuerkitoBio/goquery"
)

// Google Stroe Review Task
type GooglePlayReviewTask struct {
	appReviewRecord *models.AppReviewRecord
}

func NewGooglePlayRewiewTask(appReviewRecord *models.AppReviewRecord) *GooglePlayReviewTask {
	t := &GooglePlayReviewTask{
		appReviewRecord: appReviewRecord}
	return t
}

func (t *GooglePlayReviewTask) Run() {
	now := time.Now()
	fmt.Println("------Google start------", now.Format("2006-01-02 15:04:05"), t)
	version, updateTime, err := scrapePlayStore(t.appReviewRecord.Pkg)
	if err != nil {
		fmt.Printf("抓取 Google Play 页面失败: %v", err)
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

		fmt.Println("------Google version------", version, updateTime)
	}

	cmpValue := utils.VersionCompare(version, t.appReviewRecord.Ver)
	if cmpValue == 0 {
		fmt.Println("检测到版本审核-成功")

		t.appReviewRecord.ApproveTs = int(updateTime)
		t.appReviewRecord.TaskStatus = int(constant.TaskStop)
		t.appReviewRecord.Status = int(constant.ReviewApproved)

		hook := &webhook.ServerWebHook{}
		hook.OnWebHook(t.appReviewRecord)

		database.UpdateTaskStatus(t.appReviewRecord)
		database.UpdateStatus(t.appReviewRecord)
		StopTask(t.appReviewRecord)
	} else {
		if cmpValue == 1 {
			fmt.Println("检测到版本审核-已有更新的版本，当前任务将忽略")

			t.appReviewRecord.TaskStatus = int(constant.TaskStop)
			t.appReviewRecord.Status = int(constant.ReviewExpired)

			database.UpdateTaskStatus(t.appReviewRecord)
			database.UpdateStatus(t.appReviewRecord)
			StopTask(t.appReviewRecord)
		} else {
			fmt.Println("检测到版本审核-失败")
		}
	}
	fmt.Println("------Google end------", t.appReviewRecord.Ver, version)
}

func scrapePlayStore(pkg string) (string, int64, error) {
	url := fmt.Sprintf(constant.PlayStoreURL, pkg)
	// 发送 HTTP 请求
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 解析 HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", 0, fmt.Errorf("解析 HTML 失败: %v", err)
	}

	html := doc.Text()
	// // 测试数据
	// html := `
	// 	<div>版本列表：</div>
	// 	[[["1.2.3"]]
	// 	[[["12.3.45"]]
	// 	[[["2.15.8"]]
	// 	无效格式：[[["123.4.5"]]  [[["a.b.c"]] [[["1.2.3.4"]]
	// `

	// 提取版本信息
	versions := extractVersions(html)

	for _, v := range versions {
		fmt.Printf("提取到版本号：%s\n", v)
	}

	if len(versions) == 0 {
		return "", 0, fmt.Errorf("未找到版本信息")
	}

	updateTime := extractUpdateTime(html)
	fmt.Println("updateTime:", updateTime)
	// 清理版本字符串
	return versions[0], updateTime, nil
}

func extractVersions(html string) []string {
	// 正则表达式说明：
	// \\[\\[\\[\"   匹配固定开头[[["
	// (\\d{1,2}\\.\\d{1,2}\\.\\d{1,2}) 捕获组匹配版本号格式
	// \"\\]\\]    匹配固定结尾"]]
	re := regexp.MustCompile(`\[\[\[\"(\d{1,2}\.\d{1,2}\.\d{1,2})\"\]\]`)

	matches := re.FindAllStringSubmatch(html, -1)
	versions := make([]string, 0, len(matches))

	for _, match := range matches {
		if len(match) >= 2 {
			versions = append(versions, match[1])
		}
	}

	return versions
}

func extractUpdateTime(html string) int64 {
	// 编译正则表达式，匹配 [A,B]]] 结构
	re := regexp.MustCompile(`\[(\d{10}),\d{9}\]{3}`)

	// 查找所有匹配项
	matches := re.FindAllStringSubmatch(html, -1)

	// 提取每个匹配中的 A 值
	for _, match := range matches {
		if len(match) >= 2 {
			// return match[1]
			num, err := strconv.ParseInt(match[1], 10, 64)
			if err != nil {
				fmt.Println("转换失败:", err)
				return 0
			}
			return num
		}
	}
	return 0
}
