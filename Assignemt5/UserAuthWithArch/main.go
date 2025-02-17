package main

import (
	"github.com/TusharKM1224/config"
	"github.com/TusharKM1224/server"
)

func main() {
	config.Loadconfig()
	router := server.InitializeGin()
	router.Run(":8080")

	//fmt.Println(os.Getenv("DB_DSN"))

}
