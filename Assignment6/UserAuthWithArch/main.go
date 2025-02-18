package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/TusharKM1224/config"
	"github.com/TusharKM1224/server"
)

func main() {
	config_path := flag.String("c", "", "Config file path")
	flag.Parse()
	if *config_path == "" {
		fmt.Println("Configuration File path required!!!")
		os.Exit(0)
	}

	router := server.InitializeGin(config.Loadconfig(*config_path))
	router.Run(":8080")

}
