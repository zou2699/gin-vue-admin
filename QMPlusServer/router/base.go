package router

import (
	"github.com/gin-gonic/gin"

	"qmserver/controller/api"
)

func InitBaseRouter(Router *gin.Engine) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("regist", api.Regist)
		BaseRouter.POST("login", api.Login)
	}
}
