package tasks

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Apple Stroe Review Task
type AsReviewTask struct {
}

func NewAsReviewTask() *AsReviewTask {
	t := &AsReviewTask{}
	return t
}

func (t *AsReviewTask) Run() {
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
	appStoreURL = "https://apps.apple.com/app/id%s" // 替换为你的应用 App Store URL
)

func scrapeAppStore() (string, error) {
	url := fmt.Sprintf(appStoreURL, "1596875621")
	resp, err := http.Get(url)
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
