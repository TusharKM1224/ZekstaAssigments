package main

import (
	"github.com/TusharKM1224/config"
	"github.com/TusharKM1224/internals/handler"
	"github.com/TusharKM1224/server"
)

func main() {
	config.Loadconfig()

	router,serviceInstacec:=server.InitializeGin()
	handler.GetserviceInstace(serviceInstacec)

	router.Run(":8080")

	
}
