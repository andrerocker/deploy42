package main

import (
	"github.com/andrerocker/deploy42"
	"github.com/andrerocker/deploy42/auth"
	"gopkg.in/alecthomas/kingpin.v1"
)

var (
	baseConfig = kingpin.Flag("base-config", "base configuration file").Short('c').Default("/etc/deploy42/base.yml").String()
	authConfig = kingpin.Flag("auth-config", "auth configuration file").Short('a').Default("/etc/deploy42/auth.yml").String()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	deploy42 := deploy42.New(*baseConfig)
	deploy42.RegisterFilter("cas_tickets", auth.CasFilter(*authConfig))
	deploy42.Draw()
	deploy42.Start()
}
