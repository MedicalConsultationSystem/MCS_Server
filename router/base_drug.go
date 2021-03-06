package router

import (
	"MCS_Server/api"
	"github.com/gin-gonic/gin"
)

func InitDrugRouter(Router *gin.RouterGroup) {
	drugRouter := Router.Group("drug")
	{
		drugRouter.GET("listAll", api.ListAllDrug)
		drugRouter.POST("add", api.AddDrug)
		drugRouter.POST("findByPinyin", api.FindDrugByPinyin)
		drugRouter.POST("findByName", api.FindDrugByName)
		drugRouter.DELETE("deleteDrug", api.DeleteDrug)
		drugRouter.PUT("updateDrug",api.UpdateDrug)
	}
	drugUsageRouter := drugRouter.Group("usage")
	{
		drugUsageRouter.GET("listAll", api.ListAllUsage)
	}
	drugFrequencyRouter := drugRouter.Group("frequency")
	{
		drugFrequencyRouter.GET("listAll", api.ListAllFrequency)
	}
}
