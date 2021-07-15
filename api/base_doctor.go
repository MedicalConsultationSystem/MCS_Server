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
)

// @Tags 医生
// @Summary 获取所有医生信息
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"医生信息获取成功"}"
// @Router /doctor/listAll [get]
func ListAllDoctor(c *gin.Context) {
	if err, data := service.ListAllDoctor(); err != nil {
		global.MCS_Log.Error("医生信息获取失败", zap.Any("err", err))
		response.FailWithMsg("医生信息获取失败:"+err.Error(), c)
	} else {
		response.SuccessWithAll(data, "药物信息获取成功", c)
	}
}


// @Tags 医生
// @Summary 注册医生信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddDoctor true "医生id，医生姓名，机构id，机构名称，科室id，科室名称，医生头像链接，医生职称代码，医生职称"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"医生注册成功"}"
// @Router /doctor/add [post]
func AddDoctor(c *gin.Context) {
	var newDoctorMsg request.AddDoctor
	_ = c.ShouldBindJSON(&newDoctorMsg)
	if err:=verify.Verify(newDoctorMsg,verify.AddDoctorVerify);err!=nil{
		response.FailWithMsg(err.Error(),c)
		return
	}
	doctor := model.BaseDoctor{DoctorId: newDoctorMsg.DoctorId,DoctorName: newDoctorMsg.DoctorName,OrgId: newDoctorMsg.OrgId,
		OrgName: newDoctorMsg.OrgName,DeptId: newDoctorMsg.DeptId,DeptName: newDoctorMsg.DeptName,AvatarUrl: newDoctorMsg.AvatarUrl,
		LevelCode: newDoctorMsg.LevelCode,LevelName: newDoctorMsg.LevelName}
	if err:= service.AddDoctor(doctor);err!=nil{
		global.MCS_Log.Error("医生注册失败",zap.Any("err",err))
		response.FailWithMsg(err.Error(),c)
	}else {
		response.SuccessWithMsg("医生注册成功",c)
	}
}
