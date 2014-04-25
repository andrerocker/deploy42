package main

import (
  "log"
  "net/http"
  "./util/config"
  "github.com/andrerocker/martini"
)

func main() {
  core := config.NewCoreService()
  martini := martini.Classic()

  martini.Add("GET", "/andre", func() string {
    return "MALLLCOLLLM"
  })

  log.Fatal(http.ListenAndServe(core.BindUrl(), martini))
}
