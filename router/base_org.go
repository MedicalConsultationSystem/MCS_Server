package router

import (
	"MCS_Server/api"
	"github.com/gin-gonic/gin"
)

func InitOrgRouter(Router *gin.RouterGroup){
	orgRouter := Router.Group("organization")
	{
		orgRouter.POST("add",api.AddOrg)
		orgRouter.GET("listall",api.ListAllOrg)
	}
}