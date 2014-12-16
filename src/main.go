package main

import (
	"./util/config"
	"flag"
	"github.com/gin-gonic/gin"
)

func main() {
	configFile := flag.String("config", "deploy.yml", "configuration file")
	flag.Parse()

	http := gin.Default()
	service := config.NewCoreService(*configFile)
	commands := config.NewServiceCommands()

	commands.DrawRoutes(http)
	http.Run(service.BindUrl())
}
