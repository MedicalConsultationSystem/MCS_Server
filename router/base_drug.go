package router

import (
	"MCS_Server/api"
	"github.com/gin-gonic/gin"
)

func InitDrugRouter(Router *gin.RouterGroup){
	drugRouter := Router.Group("drug")
	{
		drugRouter.GET("listAll",api.ListAllDrug)
		drugRouter.POST("add",api.AddDrug)
	}
}
