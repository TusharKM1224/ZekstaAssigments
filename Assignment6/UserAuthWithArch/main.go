package main

import (
	"github.com/TusharKM1224/config"
	"github.com/TusharKM1224/server"
)

func main() {
	config.Loadconfig()

	router := server.InitializeGin(config.Configs())
	router.Run(":8080")

}
