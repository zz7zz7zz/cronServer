package webhook

type VersionList struct {
	BeanBase
	Version `json:"data"`
}

type Version struct {
	Total int           `json:"total"`
	Data  []VersionItem `json:"data"`
}

type VersionItem struct {
	Id         int    `json:"id"`
	Verno      string `json:"verno"`
	VerVal     int    `json:"ver_val"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
