package tasks

import (
	"fmt"
)

type AppleReviewTask struct {
}

func NewAppleReviewTask() *AppleReviewTask {
	t := &AppleReviewTask{}
	return t
}

func (t *AppleReviewTask) Run() {
	fmt.Println("------Apple------")
}
