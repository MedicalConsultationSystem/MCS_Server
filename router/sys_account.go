package router

import (
	"MCS_Server/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup)  {
	accountRouter := Router.Group("account")
	{
		accountRouter.POST("login",api.Login)
	}
}