package main

import (
	"go_code/config"
	"go_code/router"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()

	port := config.AppConfig.App.Port

	if port == "" {
		port = ":9090"
	}

	r.Run(port)
}
