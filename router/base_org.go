package router

import (
	"MCS_Server/api"
	"github.com/gin-gonic/gin"
)

func InitOrgRouter(Router *gin.RouterGroup){
	orgRouter := Router.Group("organization")
	{
		orgRouter.POST("add",api.AddOrg)
		orgRouter.GET("listAll",api.ListAllOrg)
		orgRouter.DELETE("deleteOrg",api.DeleteOrg)
		orgRouter.PUT("updateOrg",api.UpdateOrg)
		orgRouter.POST("findOrg",api.FindOrg)
	}
}