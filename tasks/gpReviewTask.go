package tasks

import (
	"cronServer/database"
	"cronServer/models"
	"cronServer/webhook"
	"fmt"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

const (
	playStoreURL = "https://play.google.com/store/apps/details?id=%s" // 替换为你的应用 Play Store URL
)

// Google Stroe Review Task
type GpReviewTask struct {
	appReviewRecord *models.AppReviewRecord
}

func NewGpRewiewTask(appReviewRecord *models.AppReviewRecord) *GpReviewTask {
	t := &GpReviewTask{
		appReviewRecord: appReviewRecord}
	return t
}

func (t *GpReviewTask) Run() {
	fmt.Println("------Google start------", t)
	version, err := scrapePlayStore(t.appReviewRecord.Pkg)
	if err != nil {
		fmt.Printf("抓取 Google Play 页面失败: %v", err)
	} else {
		if version == t.appReviewRecord.Ver {
			fmt.Println("版本一致，无需更新")
			hook := &webhook.ServerWebHook{}
			hook.OnWebHook(t.appReviewRecord)
			database.UpdateTaskStatus(t.appReviewRecord.Platform, t.appReviewRecord.Ver, t.appReviewRecord.Pkg, 3)
			database.UpdateStatus(t.appReviewRecord.Platform, t.appReviewRecord.Ver, t.appReviewRecord.Pkg, 1)
		} else {
			fmt.Println("版本不一致，需要更新")
		}
	}
	fmt.Println("------Google end------")
}

func scrapePlayStore(pkg string) (string, error) {
	url := fmt.Sprintf(playStoreURL, pkg)
	// 发送 HTTP 请求
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 解析 HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("解析 HTML 失败: %v", err)
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
		return "", fmt.Errorf("未找到版本信息")
	}

	// 清理版本字符串
	return versions[0], nil
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
