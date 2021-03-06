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


// @Tags 机构
// @Summary 增加一个机构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddOrg true "机构名称"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"机构添加成功"}"
// @Router /organization/add [post]
func AddOrg(c *gin.Context) {
	var newOrgMsg request.AddOrg
	_ = c.ShouldBindJSON(&newOrgMsg)
	if err := verify.Verify(newOrgMsg,verify.AddOrgVerify);err != nil{
		response.FailWithMsg(err.Error(), c)
		return
	}
	org := model.BaseOrg{OrgName: newOrgMsg.OrgName}
	if err := service.AddOrg(org);err != nil{
		global.MCS_Log.Error("机构添加失败",zap.Any("err", err))
		response.FailWithMsg(err.Error(),c)
	}else {
		response.SuccessWithMsg("机构添加成功",c)
	}
}

// @Tags 机构
// @Summary 获取所有机构
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"机构信息获取成功"}"
// @Router /organization/listAll [get]
func ListAllOrg(c *gin.Context) {
	if err,data := service.ListAllOrg();err != nil{
		global.MCS_Log.Error("机构信息获取失败",zap.Any("err", err))
		response.FailWithMsg("机构信息获取失败",c)
	}else {
		response.SuccessWithAll(data,"机构信息获取成功",c)
	}
}


// @Tags 机构
// @Summary 根据传入的机构结构体删除数据库中对应机构信息
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body model.BaseOrg true "数据库中机构结构体"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"机构信息删除成功"}"
// @Router /organization/deleteOrg [delete]
func DeleteOrg(c *gin.Context) {
	var baseOrg model.BaseOrg
	_ = c.ShouldBindJSON(&baseOrg)
	if err := service.DeleteOrg(baseOrg);err != nil{
		global.MCS_Log.Error("机构信息删除失败",zap.Any("err",err))
		response.FailWithMsg("机构信息删除失败:"+err.Error(),c)
	}else {
		response.SuccessWithMsg("机构信息删除成功",c)
	}
}

// @Tags 机构
// @Summary 根据传入的机构结构体更新数据库中对应机构信息
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body model.BaseOrg true "数据库中机构结构体"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"机构信息更新成功"}"
// @Router /organization/updateOrg [put]
func UpdateOrg(c *gin.Context){
	var baseOrg model.BaseOrg
	_ = c.ShouldBindJSON(&baseOrg)
	if err := verify.Verify(baseOrg,verify.BaseOrgVerify);err != nil{
		response.FailWithMsg(err.Error(),c)
		return
	}
	if err := service.UpdateOrg(baseOrg);err != nil{
		global.MCS_Log.Error("机构信息更新失败",zap.Any("err",err))
		response.FailWithMsg("机构信息更新失败:"+err.Error(),c)
	}else {
		response.SuccessWithMsg("机构信息更新成功",c)
	}

}


// @Tags 机构
// @Summary 根据机构名称模糊查询
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body request.FindOrg true "机构名称"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"机构信息获取成功"}"
// @Router /organization/findOrg [post]
func FindOrg(c *gin.Context){
	var msg request.FindOrg
	_ = c.ShouldBindJSON(&msg)
	if err,data := service.FindOrg(msg.OrgName);err != nil{
		global.MCS_Log.Error("机构信息获取失败",zap.Any("err",err))
		response.FailWithMsg("机构信息获取失败",c)
	}else {
		response.SuccessWithAll(data,"机构信息获取成功",c)
	}
}