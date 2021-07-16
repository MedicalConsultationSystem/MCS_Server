package initialize

import (
	_ "MCS_Server/docs"
	"MCS_Server/global"
	"MCS_Server/middleware"
	"MCS_Server/router"
	_ "MCS_Server/router"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {
	Router:= gin.Default()
	Router.Use(middleware.Cors())
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.MCS_Log.Info("register swagger handler")
	routerGroup := Router.Group("")
	{
		router.InitUserRouter(routerGroup)
		router.InitOrgRouter(routerGroup)
		router.InitDeptRouter(routerGroup)
		router.InitDrugRouter(routerGroup)
		router.InitDoctorRouter(routerGroup)
		router.InitConsultRouter(routerGroup)
		router.InitPrescriptionRouter(routerGroup)
	}

	return Router
}
