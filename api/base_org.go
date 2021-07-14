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
		response.FailWithMsg("机构添加失败",c)
	}else {
		response.SuccessWithMsg("机构添加成功",c)
	}
}

func ListAllOrg(c *gin.Context) {
	if err,data := service.ListAllOrg();err != nil{
		global.MCS_Log.Error("机构信息获取失败",zap.Any("err", err))
		response.FailWithMsg("机构信息获取失败",c)
	}else {
		response.SuccessWithAll(data,"机构信息获取成功",c)
	}
}
