package router

import (
	"github.com/ConnorPoor/MyTikTok/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")

	apiRouter.GET("/feed")
	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/")
	apiRouter.GET("/publish/list/")
}
