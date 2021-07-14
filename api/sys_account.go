package api

import (
	"MCS_Server/model/request"
	"MCS_Server/model/response"
	"MCS_Server/utils/verify"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var r request.Login
	_ = c.ShouldBindJSON(&r)
	if err:=verify.Verify(r,verify.LoginVerify);err != nil{
		response.FailWithMsg(err.Error(), c)
	}

}