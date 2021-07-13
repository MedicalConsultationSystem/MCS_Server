package initialize

import (
	"MCS_Server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	msg := global.MCS_Config.Mysql
	if msg.Dbname == ""{
		return nil
	}
	dbUrl := msg.Username + ":" + msg.Password + "@tcp(" + msg.Path + ")/" + msg.Dbname + "?" + msg.Config
	if db,err := gorm.Open(mysql.Open(dbUrl),&gorm.Config{});err !=nil{
		return nil
	}else {
		sqlDB,_:=db.DB()
		sqlDB.SetMaxIdleConns(msg.MaxIdleConns)
		sqlDB.SetMaxOpenConns(msg.MaxOpenConns)
		return db
	}
}
