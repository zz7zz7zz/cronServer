package tasks

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Apple Stroe Review Task
type AsReviewTask struct {
	ReviewTask
}

func NewAsReviewTask(ver string, pkg string, platform string) *AsReviewTask {
	t := &AsReviewTask{
		ReviewTask: ReviewTask{
			Ver:      ver,
			Pkg:      pkg,
			Platform: platform,
		}}
	return t
}

func (t *AsReviewTask) Run() {
	fmt.Println("------Apple------")
	version, err := scrapeAppStore(t.Pkg)
	if err != nil {
		fmt.Println("Apple Error:", err)
	}
	version = strings.ToLower(version)
	version = strings.ReplaceAll(version, "version", "")
	version = strings.TrimSpace(version)
	if version == t.Pkg {
		fmt.Println("版本一致，无需更新")
	} else {
		fmt.Println("版本不一致，需要更新")
	}
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
