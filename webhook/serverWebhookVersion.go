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

// 定义嵌套结构体
type VersionDetail struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		ID        int                       `json:"id"`
		Verno     string                    `json:"verno"`
		VerVal    int                       `json:"ver_val"`
		Creator   string                    `json:"creator"`
		CreatedAt string                    `json:"created_at"`
		UpdatedAt string                    `json:"updated_at"`
		Control   map[string]map[string]int `json:"control"` // 核心定义
	} `json:"data"`
}
