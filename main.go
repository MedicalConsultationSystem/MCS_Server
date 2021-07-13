package main

import (
	"MCS_Server/core"
	"MCS_Server/global"
	"MCS_Server/initialize"
)

func main() {
	global.MCS_Viper = core.Viper()
	global.MCS_DB = initialize.Gorm()
	if global.MCS_DB != nil{
		db,_ := global.MCS_DB.DB()
		defer db.Close()
	}
	core.RunServer()
}