package router

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")

	apiRouter.GET("/feed")
	apiRouter.GET("/user/")
	apiRouter.POST("/user/register/")
	apiRouter.POST("/user/login/")
	apiRouter.POST("/publish/action/")
	apiRouter.GET("/publish/list/")
}
