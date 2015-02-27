package auth

import (
	"github.com/andrerocker/deploy42/config"
	"github.com/andrerocker/deploy42/http"
	"github.com/lucasuyezu/golang-cas-client"
)

func CasFilter(configFile string) http.Filter {
	config := config.SimpleYAMLoad(configFile)
	service := cas.NewService(config["server"], config["service"])

	return func(request http.Request) {
		response, err := service.ValidateServiceTicket(request.RequestParameter("ticket"))

		if err != nil || !response.Status {
			request.Abort(401)
		}
	}
}
