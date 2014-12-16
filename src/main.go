package main

import (
	"./engine"
)

func main() {
	engine := engine.New()
	engine.Draw()
	engine.Start()
}
