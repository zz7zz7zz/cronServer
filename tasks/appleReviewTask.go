package tasks

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type AppleReviewTask struct {
}

func NewAppleReviewTask() *AppleReviewTask {
	t := &AppleReviewTask{}
	return t
}

func (t *AppleReviewTask) Run() {
	version, err := scrapeAppStore()
	if err != nil {
		fmt.Println("Apple Error:", err)
	}
	version = strings.ToLower(version)
	version = strings.ReplaceAll(version, "version", "")
	version = strings.TrimSpace(version)
	fmt.Printf("当前版本: %s\n", version)
	if version == "2.21.0" {

	}
}

const (
	appStoreURL = "https://apps.apple.com/app/id1596875621" // 替换为你的应用 App Store URL
)

func scrapeAppStore() (string, error) {
	resp, err := http.Get(appStoreURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	version := doc.Find(".whats-new__latest__version").Text()
	return version, nil
}
