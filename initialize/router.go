package initialize

import (
	"MCS_Server/router"
	_ "MCS_Server/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router:= gin.Default()

	routerGroup := Router.Group("")
	{
		router.InitUserRouter(routerGroup)
		router.InitOrgRouter(routerGroup)
	}

	return Router
}
