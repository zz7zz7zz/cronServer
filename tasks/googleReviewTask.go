package tasks

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	playStoreURL = "https://play.google.com/store/apps/details?id=com.inhobichat.hobichat"                                               // 替换为你的应用 Play Store URL
	userAgents   = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36" // 随机 User-Agen
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
		fmt.Printf("当前版本ppp: %s\n", version)
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

	// 提取版本信息
	// version := doc.Find(".hAyfc .htlgb").First().Text()
	version := ""
	doc.Find(".hAyfc").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), "Current Version") {
			version = s.Find(".htlgb").First().Text()
		}
	})

	if version == "" {
		return "", fmt.Errorf("未找到版本信息")
	}

	// 清理版本字符串
	version = strings.TrimSpace(version)
	return version, nil
}
