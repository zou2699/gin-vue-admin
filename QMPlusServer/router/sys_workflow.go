package router

import (
	"github.com/gin-gonic/gin"

	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"
)

func InitWorkflowRouter(Router *gin.RouterGroup) {
	WorkflowRouter := Router.Group("workflow").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		WorkflowRouter.POST("createWorkFlow", api.CreateWorkFlow) // 创建工作流
	}
}
