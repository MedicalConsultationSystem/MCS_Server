package api

import (
	"MCS_Server/model/request"
	"MCS_Server/utils/verify"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var r request.Login
	_ = c.ShouldBindJSON(&r)
	if err:=verify.Verify(r,verify.LoginVerify);err != nil{
		fmt.Println(err)
	}
	fmt.Println("get this")
}