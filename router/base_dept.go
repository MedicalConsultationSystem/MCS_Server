package router

import (
	"MCS_Server/api"
	"github.com/gin-gonic/gin"
)

func InitDeptRouter(Router *gin.RouterGroup)  {
	deptRouter := Router.Group("dept")
	{
		deptRouter.POST("add",api.AddDept)
		deptRouter.GET("listAll",api.ListAllDept)
	}
}
