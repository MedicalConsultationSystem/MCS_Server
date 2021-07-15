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

// @Tags 药物
// @Summary 增加一种药物
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddDrug true "药物通用名称,商品名,拼音码,药品规格,包装单位,药品价格,剂量,剂量单位,产地,批准文号"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"药物添加成功"}"
// @Router /drug/add [post]
func AddDrug(c *gin.Context) {
	var newDrugMsg request.AddDrug
	_ = c.ShouldBindJSON(&newDrugMsg)
	if err := verify.Verify(newDrugMsg,verify.AddDrugVerify);err!=nil{
		response.FailWithMsg(err.Error(),c)
		return
	}
	drug := model.BaseDrug{DrugName: newDrugMsg.DrugName,TradeName: newDrugMsg.TradeName,PinyinCode: newDrugMsg.PinyinCode,
		Specification: newDrugMsg.Specification,PackUnit: newDrugMsg.PackUnit,Price: newDrugMsg.Price,Dose: newDrugMsg.Dose,
		DoseUnit: newDrugMsg.DoseUnit,FactoryName: newDrugMsg.FactoryName,ApprovalNumber: newDrugMsg.ApprovalNumber}
	if err := service.AddDrug(drug);err !=nil{
		global.MCS_Log.Error("药物添加失败",zap.Any("err",err))
		response.FailWithMsg(err.Error(),c)
	}else {
		response.SuccessWithMsg("药物添加成功",c)
	}

}