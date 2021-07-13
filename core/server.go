package core

import (
	"MCS_Server/global"
	"MCS_Server/initialize"
	"fmt"
	"github.com/fvbock/endless"
	"go.uber.org/zap"
	"log"
	"time"
)

func RunServer() {
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.MCS_Config.Information.Port)
	server := endless.NewServer(address, Router)
	server.ReadHeaderTimeout = 10 * time.Millisecond
	server.WriteTimeout = 10 * time.Second
	server.MaxHeaderBytes = 1 << 20

	global.MCS_Log.Info("server run success on ", zap.String("address", address))
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("server err: %v", err)
	}
}
