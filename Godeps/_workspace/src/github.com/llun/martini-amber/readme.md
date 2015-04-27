Martini-Amber
=============

[Amber](https://github.com/eknkc/amber) template for Martini.

## Usage

This project uses [martini-contrib/render](https://github.com/martini-contrib/render) as template. So the usage is almost the same.

```go
// main.go
package main

import (
  "github.com/go-martini/martini"
  "github.com/llun/martini-amber"
)

func main() {
  m := martini.Classic()
  // render html templates from templates directory
  m.Use(martini_amber.Renderer())

  m.Get("/", func(r martini_amber.Render) {
    r.Amber(200, "template_name", map[string]interface{}{
		"title": "Hello, World"
	})
  })

  m.Run()
}
```

```amber
<!-- templates/template_name.amber -->
h2 $title

