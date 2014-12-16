package main

import (
	"./util/config"
	"github.com/gin-gonic/gin"
)

func main() {
	http := gin.Default()
	service := config.NewCoreService()
	commands := config.NewServiceCommands()

	commands.DrawRoutes(http)
	http.Run(service.BindUrl())
}
