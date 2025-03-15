package tasks

import (
	"fmt"
)

type GoogleReviewTask struct {
}

func NewGoogleRewiewTask() *GoogleReviewTask {
	t := &GoogleReviewTask{}
	return t
}

func (t *GoogleReviewTask) Run() {
	fmt.Println("------Google------")
}
