package main

import (
	"MCS_Server/core"
	"MCS_Server/global"
	"MCS_Server/initialize"
)

// @title Swagger MCS_Server API
// @version 0.0.1
// @description MCS_Server API 接口文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.MCS_Viper = core.Viper()
	global.MCS_Log = core.Zap()
	global.MCS_DB = initialize.Gorm()
	if global.MCS_DB != nil {
		db, _ := global.MCS_DB.DB()
		defer db.Close()
	}
	core.RunServer()
}
