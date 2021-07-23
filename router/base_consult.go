package router

import (
	"MCS_Server/api"
	"github.com/gin-gonic/gin"
)

func InitConsultRouter(Router *gin.RouterGroup)  {
	consultRouter := Router.Group("consult")
	{
		consultRouter.POST("add",api.AddConsult)
		consultRouter.POST("findByUser",api.FindConsultByUser)
		consultRouter.POST("findByDoctor",api.FindConsultByDoctor)
		consultRouter.POST("accept",api.AcceptConsult)
		consultRouter.POST("finish",api.FinishConsult)
	}
}