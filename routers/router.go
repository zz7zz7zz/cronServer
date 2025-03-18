package routers

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.Engine) {
	api := r.Group("/api")
	InitAppreview(api)
}
