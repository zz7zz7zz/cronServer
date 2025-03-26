package models

type RequestStart struct {
	Ver      string `json:"ver" form:"ver" binding:"required"`
	Pkg      string `json:"pkg" form:"pkg" binding:"required"`
	Platform string `json:"platform" form:"platform" binding:"required,oneof=web ios android"`
	Status   string `json:"status" form:"status" binding:"omitempty"`
}

type Platform string

const (
	PlatformWeb     Platform = "web"
	PlatformIOS     Platform = "ios"
	PlatformAndroid Platform = "android"
)
