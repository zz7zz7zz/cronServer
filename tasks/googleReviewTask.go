package tasks

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

const (
	playStoreURL = "https://play.google.com/store/apps/details?id=com.inhobichat.hobichat" // 替换为你的应用 Play Store URL
)

type GoogleReviewTask struct {
}

func NewGoogleRewiewTask() *GoogleReviewTask {
	t := &GoogleReviewTask{}
	return t
}

func (t *GoogleReviewTask) Run() {
	fmt.Println("------Google------")
	version, err := scrapePlayStore()
	if err != nil {
		fmt.Printf("抓取 Google Play 页面失败: %v", err)
	} else {
		fmt.Printf("当前版本: %s\n", version)
	}
}

func scrapePlayStore() (string, error) {
	// 发送 HTTP 请求
	resp, err := http.Get(playStoreURL)
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
	fmt.Println("提取到的版本号：")
	for _, v := range versions {
		fmt.Println(v)
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
