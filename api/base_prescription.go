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
func AddPrescription(c *gin.Context) {
	var newPrescriptionMsg request.AddPrescription
	_ = c.ShouldBindJSON(&newPrescriptionMsg)
	if err := verify.Verify(newPrescriptionMsg, verify.AddPrescriptionVerify); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	prescription := model.BasePrescription{OrgId: newPrescriptionMsg.OrgId, UserId: newPrescriptionMsg.UserId,
		ConsultId: newPrescriptionMsg.ConsultId, PrescriptionType: newPrescriptionMsg.PrescriptionType, DoctorId: newPrescriptionMsg.DoctorId,
		DoctorName: newPrescriptionMsg.DoctorName, CreateTime: time.Now(), PrescriptionStatus: "0"}
	if err := service.AddPrescription(prescription); err != nil {
		global.MCS_Log.Error("处方新增失败", zap.Any("err", err))
		response.FailWithMsg(err.Error(), c)
	} else {
		response.SuccessWithMsg("处方新增成功", c)
	}
}

// @Tags 处方
// @Summary 提交所有处方
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SubmitPrescription true "机构id，用户id，问诊信息id，处方类型，医生id，医生名称"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"处方提交成功"}"
// @Router /prescription/submit [post]
func SubmitPrescription(c *gin.Context) {
	response.SuccessWithMsg("接口还没写完，让你失望了", c)
}

// @Tags 处方
// @Summary 查询处方信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ListPrescription true "问诊信息 id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"处方查询成功"}"
// @Router /prescription/list [post]
func ListPrescription(c *gin.Context) {
	var msg request.ListPrescription
	_ = c.ShouldBindJSON(&msg)
	if err := verify.Verify(msg, verify.ListPrescriptionVerify); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	if err,data := service.ListPrescription(msg.ConsultId);err != nil{
		global.MCS_Log.Error("处方信息获取失败",zap.Any("err",err))
		response.FailWithMsg("处方信息获取失败:"+err.Error(),c)
	}else {
		response.SuccessWithAll(data,"处方信息获取成功",c)
	}
	//response.SuccessWithMsg("接口还没写完，让你失望了", c)
}
