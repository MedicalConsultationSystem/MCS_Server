package config

type Server struct {
	Mysql       Mysql       `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Information Information `mapstructure:"information" json:"information" yaml:"information"`
	Zap         Zap         `mapstructure:"zap" json:"zap" yaml:"zap"`
	JWT         JWT         `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
