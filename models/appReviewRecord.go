package models

type AppReviewRecord struct {
	ID           int    `json:"id"`
	Ver          string `json:"ver"`
	Pkg          string `json:"pkg"`
	Platform     string `json:"platform"`
	Status       int    `json:"status"`
	ApproveTs    int    `json:"approve_ts"`
	TaskStatus   int    `json:"task_status"`
	TaskCreateTs int    `json:"task_create_ts"`
	Channel      string `json:"channel"`
}
