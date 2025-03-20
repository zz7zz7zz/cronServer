package tasks

import (
	"cronServer/models"
	"cronServer/webhook"
	"fmt"
	"net/http"
	"strings"

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
	fmt.Println("------Apple start------", t)
	version, err := scrapeAppStore(t.appReviewRecord.Pkg)
	version = strings.ReplaceAll(version, "Version", "")
	version = strings.TrimSpace(version)
	if err != nil {
		fmt.Println("Apple Error:", err)
	} else {
		fmt.Println("------Apple version------", version)
	}
	version = strings.ToLower(version)
	version = strings.ReplaceAll(version, "version", "")
	version = strings.TrimSpace(version)
	if version == t.appReviewRecord.Ver {
		fmt.Println("版本一致，无需更新")
		hook := &webhook.ServerWebHook{}
		hook.OnWebHook(t.appReviewRecord)
	} else {
		fmt.Print("版本不一致，需要更新\n", t.appReviewRecord.Ver, version)
	}
	fmt.Println("------Apple end------")
}

const (
	appStoreURL = "https://apps.apple.com/app/id%s" // 替换为你的应用 App Store URL
)

func scrapeAppStore(pkg string) (string, error) {
	url := fmt.Sprintf(appStoreURL, pkg)
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
