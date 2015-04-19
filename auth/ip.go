package auth

import (
	"github.com/andrerocker/deploy42/config"
	"github.com/gin-gonic/gin"
	"net"
	"strings"
)

func IpRestrictionFilter(configFile string) gin.HandlerFunc {
	config := config.SimpleYAMLoad(configFile)
	ranges := resolveCIDRs(config["ip_restriction"].([]interface{}))

	return func(context *gin.Context) {
		if isInvalidIp(ranges, resolveClientIP(context)) {
			context.String(403, "unathorized")
			context.Abort()
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

func resolveClientIP(context *gin.Context) net.IP {
	clientIp := strings.Split(context.ClientIP(), ":")[0]
	return net.ParseIP(clientIp)
}
