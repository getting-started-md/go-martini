package main

import (
  "github.com/go-martini/martini"
  "github.com/llun/martini-amber"
)

func main() {
  m := martini.Classic()

  m.Use(martini_amber.Renderer(map[string]string{}))

  m.Get("/", func() string {
    return "Hello world!"
  })

  m.Get("/hello", func(r martini_amber.Render) {
    r.Amber(200, "hello", map[string]interface{}{
        "title": "Hello, World",
    })
  })

  m.Run()
}