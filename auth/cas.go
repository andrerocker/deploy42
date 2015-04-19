package auth

import (
	"github.com/andrerocker/deploy42/config"
	"github.com/andrerocker/golang-cas-client"
	"github.com/gin-gonic/gin"
)

func CasFilter(configFile string) gin.HandlerFunc {
	config := config.SimpleYAMLoad(configFile)
	service := cas.NewService(config["server"].(string), config["service"].(string))

	return func(context *gin.Context) {
		request := context.Request
		request.ParseForm()
		response, err := service.ValidateServiceTicket(request.Form.Get("ticket"))

		if err != nil || !response.Status {
			context.String(401, "unauthorized")
			context.Abort()
		}
	}
}
