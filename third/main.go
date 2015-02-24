package main

import (
	"github.com/andrerocker/deploy42"
)

func main() {
	deploy42 := deploy42.New()
	deploy42.Draw()
	deploy42.Start()
}
