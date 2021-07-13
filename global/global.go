package global

import (
	"MCS_Server/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	MCS_DB     *gorm.DB
	MCS_Viper  *viper.Viper
	MCS_Config config.Server
	MCS_Log    *zap.Logger
)

const (
	ConfigName = "config"
	ConfigFile = "config.yaml"
)
