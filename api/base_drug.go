package api

import (
	"MCS_Server/global"
	"MCS_Server/model/response"
	"MCS_Server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)



// @Tags 药物
//@Summary 获取所有药物信息
//@Security ApiKeyAuth
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"药物信息获取成功"}"
// @Router /drug/listAll [get]
func ListAllDrug(c *gin.Context) {
	if err,data := service.ListAllDrug();err!=nil{
		global.MCS_Log.Error("药物信息获取失败",zap.Any("err",err))
		response.FailWithMsg("药物信息获取失败:"+err.Error(),c)
	}else {
		response.SuccessWithAll(data,"药物信息获取成功",c)
	}
}