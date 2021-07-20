package api

import (
	"MCS_Server/global"
	"MCS_Server/model"
	"MCS_Server/model/request"
	"MCS_Server/model/response"
	"MCS_Server/service"
	"MCS_Server/utils/verify"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// @Tags 问诊
// @Summary 创建一个问诊信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddConsult true "问诊信息的各个字段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"问诊信息添加成功"}"
// @Router /consult/add [post]
func AddConsult(c *gin.Context)  {
	var newConsultMsg request.AddConsult
	_ = c.ShouldBindJSON(&newConsultMsg)
	fmt.Println(newConsultMsg)
	if err := verify.Verify(newConsultMsg,verify.AddConsultVerify);err != nil{
		response.FailWithMsg(err.Error(),c)
		return
	}
	consult := model.BaseConsult{OrgId: newConsultMsg.OrgId,OrgName: newConsultMsg.OrgName,DeptId: newConsultMsg.DeptId,
		DeptName: newConsultMsg.DeptName,DoctorId: newConsultMsg.DoctorId,DoctorName: newConsultMsg.DoctorName,
		DoctorLevelCode: newConsultMsg.DoctorLevelCode,DoctorLevelName: newConsultMsg.DoctorLevelName,CreateUserId: newConsultMsg.CreateUserId,
		PersonName: newConsultMsg.PersonName,PersonCardType: newConsultMsg.PersonCardType,PersonCardId: newConsultMsg.PersonCardId,
		PersonGenderCode: newConsultMsg.PersonGenderCode,PersonGenderName: newConsultMsg.PersonGenderName,PersonBirthDate: newConsultMsg.PersonBirthDate,
		PersonAge: newConsultMsg.PersonAge,PersonPhoneNo: newConsultMsg.PersonPhoneNo,Question: newConsultMsg.Question,
		Diagnosis: newConsultMsg.Diagnosis,DrugIds: newConsultMsg.DrugIds,DrugNames: newConsultMsg.DrugNames,PhotoIds: newConsultMsg.PhotoIds,
		ConsultStatus: newConsultMsg.ConsultStatus,CreateTime: time.Now()}
	if err := service.AddConsult(consult);err!=nil{
		global.MCS_Log.Error("问诊信息添加失败",zap.Any("err",err))
		response.FailWithMsg(err.Error(),c)
	}else {
		response.SuccessWithMsg("问诊信息添加成功",c)
	}
}


// @Tags 问诊
// @Summary 根据名称用户id查询该用户的问诊信息
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body request.FindConsultByUser true "用户id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"问诊信息获取成功"}"
// @Router /consult/findByUser [post]
func FindConsultByUser(c *gin.Context) {
	var msg request.FindConsultByUser
	_ = c.ShouldBindJSON(&msg)
	if err,data := service.FindConsultByUser(msg.CreateUserId);err !=nil{
		global.MCS_Log.Error("问诊信息获取失败",zap.Any("err",err))
		response.FailWithMsg("问诊信息获取失败:"+err.Error(),c)
	}else {
		response.SuccessWithAll(data,"问诊信息获取成功",c)
	}
}

// @Tags 问诊
// @Summary 根据医生id查询问诊信息
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body request.FindConsultByDoctor true "医生id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"问诊信息获取成功"}"
// @Router /consult/findByDoctor [post]
func FindConsultByDoctor(c *gin.Context){
	var msg request.FindConsultByDoctor
	_ = c.ShouldBindJSON(&msg)
	if err,data := service.FindConsultByDoctor(msg.DoctorId);err != nil{
		global.MCS_Log.Error("问诊信息获取失败",zap.Any("err",err))
		response.FailWithMsg("问诊信息获取失败:"+err.Error(),c)
	}else {
		response.SuccessWithAll(data,"问诊信息获取成功",c)
	}
}