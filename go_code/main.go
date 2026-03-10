package main

import (
	"fmt"
	"go_code/config"
)

func main() {
	config.InitConfig()
	fmt.Println(config.AppConfig.App.Port)
}
