package api

import (
	"MCS_Server/global"
	"MCS_Server/model"
	"MCS_Server/model/request"
	"MCS_Server/model/response"
	"MCS_Server/service"
	"MCS_Server/utils/verify"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)


// @Tags 处方
// @Summary 增加一个处方
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddPrescription true "机构id，用户id，问诊信息id，处方类型，医生id，医生名称"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"处方新增成功"}"
// @Router /prescription/add [post]
func AddPrescription(c *gin.Context){
	var newPrescriptionMsg request.AddPrescription
	_ = c.ShouldBindJSON(&newPrescriptionMsg)
	if err := verify.Verify(newPrescriptionMsg,verify.AddPrescriptionVerify);err!=nil{
		response.FailWithMsg(err.Error(),c)
		return
	}
	prescription := model.BasePrescription{OrgId: newPrescriptionMsg.OrgId,UserId: newPrescriptionMsg.UserId,
		ConsultId: newPrescriptionMsg.ConsultId,PrescriptionType: newPrescriptionMsg.PrescriptionType,DoctorId: newPrescriptionMsg.DoctorId,
		DoctorName: newPrescriptionMsg.DoctorName,CreateTime: time.Now(),PrescriptionStatus: "0"}
	if err := service.AddPrescription(prescription);err!=nil{
		global.MCS_Log.Error("处方新增失败",zap.Any("err",err))
		response.FailWithMsg(err.Error(),c)
	}else {
		response.SuccessWithMsg("处方新增成功",c)
	}
}

