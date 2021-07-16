package router

import (
	"MCS_Server/api"
	"github.com/gin-gonic/gin"
)

func InitPrescriptionRouter(Router *gin.RouterGroup) {
	prescription := Router.Group("prescription")
	{
		prescription.POST("add",api.AddPrescription)
	}
}