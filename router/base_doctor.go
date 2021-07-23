package router

import (
	"MCS_Server/api"
	"github.com/gin-gonic/gin"
)

func InitDoctorRouter(Router *gin.RouterGroup){
	doctorRouter := Router.Group("doctor")
	{
		doctorRouter.GET("listAll",api.ListAllDoctor)
		doctorRouter.POST("add",api.AddDoctor)
		doctorRouter.POST("findByName",api.FindDoctorByName)
		doctorRouter.PUT("updateDoctor",api.UpdateDoctor)
	}
}
