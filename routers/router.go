package routers

import (
	"open.com/cronServer/appreview/router"

	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	api := r.Group("/api")
	router.InitAppreview(api)
	router.InitAppreviewV1(api)
	router.InitAppreviewV2(api)
}
