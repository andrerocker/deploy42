package auth

import (
	"github.com/andrerocker/deploy42/config"
	"github.com/andrerocker/deploy42/http"
	"net"
)

func IpRestrictionFilter(configFile string) http.Handler {
	config := config.SimpleYAMLoad(configFile)
	ranges := resolveCIDRs(config["ip_restriction"].([]interface{}))

	return func(request http.Request) {
		userIp := net.ParseIP("127.0.0.1") //request.UserIP()

		if isInvalidIp(ranges, userIp) {
			request.Abort(403)
		}
	}
}

func isInvalidIp(ranges []*net.IPNet, userIp net.IP) bool {
	for _, cidr := range ranges {
		if cidr.Contains(userIp) {
			return false
		}
	}

	return true
}

func resolveCIDRs(ips []interface{}) []*net.IPNet {
	ranges := make([]*net.IPNet, 0)

	for _, rangeName := range ips {
		_, cidr, _ := net.ParseCIDR(rangeName.(string))
		ranges = append(ranges, cidr)
	}

	return ranges
}
