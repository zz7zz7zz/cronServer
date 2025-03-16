package models

type AppReviewRecord struct {
	ID        int    `json:"id"`
	Ver       string `json:"ver"`
	Pkg       string `json:"pkg"`
	Platform  string `json:"platform"`
	Status    int    `json:"status"`
	TimeStamp int    `json:"time_stamp"`
}
