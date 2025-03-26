package models

type RequestStart struct {
	Ver      string `json:"ver"`
	Pkg      string `json:"pkg"`
	Platform string `json:"platform"`
	Status   string `json:"status"`
}
