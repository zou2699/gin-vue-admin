package router

import (
	"github.com/gin-gonic/gin"

	"gin-vue-admin/controller/api"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("regist", api.Regist)
		BaseRouter.POST("login", api.Login)
	}
	return BaseRouter
}
