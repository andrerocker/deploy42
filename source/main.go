package main

import (
  "net/http"
  "./util/config"
  "github.com/andrerocker/martini"
)

func main() {
  martini := martini.Classic()
  service := config.NewCoreService()
  commands := config.NewServiceCommands()

  commands.DrawRoutes(martini)
  http.ListenAndServe(service.BindUrl(), martini)
}
