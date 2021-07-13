package initialize

import (
	"MCS_Server/global"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	msg := global.MCS_Config.Mysql
	if msg.Dbname == ""{
		return nil
	}
	dburl := msg.Username + ":" + msg.Password + "@tcp(" + msg.Path + ")/" + msg.Dbname + "?" + msg.Config

}
