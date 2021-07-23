package api

import (
	"MCS_Server/global"
	"MCS_Server/middleware"
	"MCS_Server/model/request"
	"MCS_Server/model/response"
	"MCS_Server/service"
	"MCS_Server/utils/verify"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// @Tags 账户
// @Summary 用户登录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Login true "微信小程序 Appid AppSecret code"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登录成功"}"
// @Router /account/login [post]
func Login(c *gin.Context) {
	var loginMsg request.Login
	_ = c.ShouldBindJSON(&loginMsg)
	if err := verify.Verify(loginMsg, verify.LoginVerify); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	if err, account := service.Login(loginMsg); err != nil {
		global.MCS_Log.Error("登陆失败!"+err.Error(), zap.Any("err", err))
		response.FailWithMsg(err.Error(), c)
	} else {
		j := &middleware.JWT{SigningKey: []byte(global.MCS_Config.JWT.SigningKey)}
		claims := request.Claims{
			UserId:     loginMsg.OpenId,
			BufferTime: global.MCS_Config.JWT.BufferTime,
			StandardClaims: jwt.StandardClaims{
				NotBefore: time.Now().Unix() - 1000,
				ExpiresAt: time.Now().Unix() + global.MCS_Config.JWT.ExpiresTime,
				Issuer:    "MCS_admin",
			},
		}
		token, err := j.CreateToken(claims)
		if err != nil {
			global.MCS_Log.Error("获取Token失败！", zap.Any("err", err))
			response.FailWithMsg(err.Error(), c)
		}
		response.SuccessWithAll(response.AccountWithToken{
			BaseAccount: account, Token: token,
		}, "登录成功", c)
	}
}
