package initialize

import (
	"MCS_Server/router"
	_ "MCS_Server/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router:= gin.Default()
	//Router.Use(middleware.Cors())
	routerGroup := Router.Group("")
	{
		router.InitUserRouter(routerGroup)
		router.InitOrgRouter(routerGroup)
		router.InitDeptRouter(routerGroup)
	}

	return Router
}
