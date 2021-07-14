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
)

func Login(c *gin.Context) {
	var loginMsg request.Login
	_ = c.ShouldBindJSON(&loginMsg)
	if err:=verify.Verify(loginMsg,verify.LoginVerify);err != nil{
		response.FailWithMsg(err.Error(), c)
		return
	}
	a := &model.BaseAccount{Code: loginMsg.Code,PhoneNo: loginMsg.PhoneNo}
	if err,account := service.Login(a);err !=nil{
		global.MCS_Log.Error("登陆失败!",zap.Any("err", err))
	}else {
		fmt.Println(account)
	}
}