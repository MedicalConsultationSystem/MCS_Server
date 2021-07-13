package global

import (
	"MCS_Server/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	MCS_DB *gorm.DB
	MCS_Viper *viper.Viper
	MCS_Config config.Server
)

const (
	ConfigName = "config"
	ConfigFile = "config.yaml"
)
