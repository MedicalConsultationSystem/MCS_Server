package core

import "MCS_Server/initialize"

func RunServer(){
	Router := initialize.Routers()
	_ = Router.Run(":23333")
}
