package core

import (
	"MCS_Server/global"
	"MCS_Server/initialize"
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"time"
)

func RunServer(){
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.MCS_Config.Information.Port)
	server := endless.NewServer(address,Router)
	server.ReadHeaderTimeout = 10 * time.Millisecond
	server.WriteTimeout = 10 * time.Second
	server.MaxHeaderBytes = 1 << 20

	fmt.Printf("server run on localhost:%d\n",global.MCS_Config.Information.Port)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("server err: %v", err)
	}
}
