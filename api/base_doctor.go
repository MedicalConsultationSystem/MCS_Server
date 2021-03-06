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
		response.SuccessWithAll(data, "医生信息获取成功", c)
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
	if err := verify.Verify(newDoctorMsg, verify.AddDoctorVerify); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	doctor := model.BaseDoctor{DoctorId: newDoctorMsg.DoctorId, DoctorName: newDoctorMsg.DoctorName, OrgId: newDoctorMsg.OrgId,
		OrgName: newDoctorMsg.OrgName, DeptId: newDoctorMsg.DeptId, DeptName: newDoctorMsg.DeptName, AvatarUrl: newDoctorMsg.AvatarUrl,
		LevelCode: newDoctorMsg.LevelCode, LevelName: newDoctorMsg.LevelName}
	if err := service.AddDoctor(doctor); err != nil {
		global.MCS_Log.Error("医生注册失败", zap.Any("err", err))
		response.FailWithMsg(err.Error(), c)
	} else {
		response.SuccessWithMsg("医生注册成功", c)
	}
}

// @Tags 医生
// @Summary 根据姓名模糊查询医生
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body request.FindDoctorByName true "医生姓名"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"医生信息获取成功"}"
// @Router /doctor/findByName [post]
func FindDoctorByName(c *gin.Context) {
	var msg request.FindDoctorByName
	_ = c.ShouldBindJSON(&msg)
	if err, data := service.FindDoctorByName(msg.DoctorName); err != nil {
		global.MCS_Log.Error("医生信息获取失败", zap.Any("err", err))
		response.FailWithMsg("医生信息获取失败:"+err.Error(), c)
	} else {
		response.SuccessWithAll(data, "医生信息获取成功", c)
	}
}


// @Tags 医生
// @Summary 更新医生信息
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body model.BaseDoctor true "数据库中医生结构体"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"医生信息更新成功"}"
// @Router /doctor/updateDoctor [put]
func UpdateDoctor(c *gin.Context) {
	var baseDoctor model.BaseDoctor
	_ = c.ShouldBindJSON(&baseDoctor)
	if err := verify.Verify(baseDoctor, verify.BaseDoctorVerify); err != nil {
		response.FailWithMsg(err.Error(),c)
		return
	}
	if err := service.UpdateDoctor(baseDoctor);err != nil{
		global.MCS_Log.Error("医生信息更新失败",zap.Any("err",err))
		response.FailWithMsg("医生信息更新失败:"+err.Error(),c)
	}else {
		response.SuccessWithMsg("医生信息更新成功",c)
	}
}

// @Tags 医生
// @Summary 删除医生信息
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body request.DeleteDoctor true "医生id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"医生信息删除成功"}"
// @Router /doctor/deleteDoctor [delete]
func DeleteDoctor(c *gin.Context){
	var msg request.DeleteDoctor
	_ = c.ShouldBindJSON(&msg)
	if err := service.DeleteDoctor(msg.DoctorId);err !=nil{
		global.MCS_Log.Error("医生信息删除失败",zap.Any("err",err))
		response.FailWithMsg("医生信息删除失败:"+err.Error(),c)
	}else {
		response.SuccessWithMsg("医生信息删除成功",c)
	}
}