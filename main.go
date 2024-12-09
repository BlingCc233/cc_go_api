package main

import (
	"studyGo/config"
	"studyGo/router"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()

	port := config.AppConfig.App.Port
	if port == "" {
		port = "3051"
	}

	r.Run(port)

}
