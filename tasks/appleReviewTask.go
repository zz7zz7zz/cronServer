package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AppleReviewTask struct {
}

func NewAppleReviewTask() *AppleReviewTask {
	t := &AppleReviewTask{}
	return t
}

func (t *AppleReviewTask) Run() {
	fmt.Println("------Apple------")
	version, err := getVersionViaAppFollow()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Version:", version)
}

func getVersionViaAppFollow() (string, error) {
	// url := "https://api.appfollow.io/version?app_id=id1596875621&country=us&device=iphone"
	// resp, err := http.Get(url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.appfollow.io/apps/app?apps_id=com.inhobichat.hobichat", nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth("AppFollow-966-2b3da05ee", "")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Version string `json:"version"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.Version, nil
}
