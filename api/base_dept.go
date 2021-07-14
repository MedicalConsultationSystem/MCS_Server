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

func AddDept(c *gin.Context){
	var newDeptMsg request.AddDept
	_ = c.ShouldBindJSON(&newDeptMsg)
	if err := verify.Verify(newDeptMsg,verify.AddDeptVerify);err!=nil{
		response.FailWithMsg(err.Error(),c)
		return
	}
	dept := model.BaseDept{DeptName: newDeptMsg.DeptName}
	if err:= service.AddDept(dept);err != nil{
		global.MCS_Log.Error("科室添加失败",zap.Any("err",err))
		response.FailWithMsg(err.Error(),c)
	}else {
		response.SuccessWithMsg("科室添加成功",c)
	}
}

func ListAllDept(c *gin.Context) {
	if err,data := service.ListAllDept();err!=nil{
		global.MCS_Log.Error("科室信息获取失败",zap.Any("err",err))
		response.FailWithMsg("科室信息获取失败",c)
	}else {
		response.SuccessWithAll(data,"科室信息获取成功",c)
	}
}