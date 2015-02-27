package main

import (
	"github.com/alecthomas/kingpin"
	"github.com/andrerocker/deploy42"
)

var (
	baseConfig = kingpin.Flag("config", "deploy42 configuration file").Short('c').Default("/etc/deploy42/config.yml").String()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	deploy42 := deploy42.New(*baseConfig)
	deploy42.Draw()
	deploy42.Start()
}
