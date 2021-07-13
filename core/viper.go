package core

import (
	"MCS_Server/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	viper.AddConfigPath(".")
	v := viper.New()
	v.SetConfigName(global.ConfigName)
	v.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.MCS_Config); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.MCS_Config); err != nil {
		fmt.Println(err)
	}

	return v
}