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
		port = ":8080"
	}

	r.Run(port)
}
