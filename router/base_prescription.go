package router

import (
	"MCS_Server/api"
	"github.com/gin-gonic/gin"
)

func InitPrescriptionRouter(Router *gin.RouterGroup) {
	prescription := Router.Group("prescription")
	{
		prescription.POST("add",api.AddPrescription)
		prescription.POST("submit",api.SubmitPrescription)
		prescription.POST("list",api.ListPrescription)
		prescription.POST("addDrug",api.AddPrescriptionDrug)
		prescription.DELETE("delDrug",api.DeletePrescriptionDrug)
		prescription.DELETE("delPre",api.DeletePrescription)

	}
}