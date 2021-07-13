package MCS_Server

import (
	"MCS_Server/core"
	"MCS_Server/global"
)

func main() {
	global.MCS_Viper = core.Viper()
}