
## Background

Martini is a lightweight web framework for Go.

Martini provides a HTTP Server along with a light weight routing DSL and parameter parsing.

Martini is a great fit for services and APIs.


## Installing Go

This guide will focus on OSX.

First ensure homebrew is installed and up to date.

[http://brew.sh/](http://brew.sh/)

Next install Go

`brew install go`

Next ensure environment variables are set appropriately

```bash
export GOROOT=/usr/local/opt/go/libexec
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/opt/go/libexec/bin
export PATH=$PATH:$GOPATH/bin
```

In order for things like godeps to work correctly you must put your projects in `$GOPATH/src`


## Setup the project

Create a namespace folder for your projects in __$GOPATH/src__

`mkdir $GOPATH/src/getting-started.md`

Create a folder for your new martini project in __$GOPATH/src/getting-started.md__

`mkdir $GOPATH/src/getting-started.md/hello`


## Installing Martini

First install Martini in the project folder.

In the __hello__ folder, 

`$ go get github.com/go-martini/martini`


## A simple server


Create your Martini appilication by making a new file named `server.go` in the 
__hello__ folder.



**server.go**

```go
package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Run()
}
```

First we declare our package "main" since this file will be the main entry point into our application.

Next we import dependencies, in this case martini. This works in conjunction with the `go get` command to import the 3rd party package.

Next we declare the main function, which is the code execution entry point.

m is declared as a martini server instance.
then we use the router DSL `.Get` to delcare a new route "/"

When a get request is sent for /, the code declared in the function will be executed.

Basic routes can just return a string, in this case "Hello World"

Lastly we tell our server to run.


## Templates and Views

Martini ships with no view / templating, but setting it up is made simple thanks to a great community of gophers solving all these problems and sharing their effort.

Amber is a popular Go templating language, and has a Jade / HAML feel. I am a big fan of Jade / HAML so i've chosen Amber as the templating language for this demo.

First install Martini-Amber with `go get github.com/llun/martini-amber`

Then add it to the imports section of your __server.go__

```go
import (
  "github.com/go-martini/martini"
  "github.com/llun/martini-amber"
)
```

Add the renderer to martini.

`m.Use(martini_amber.Renderer(map[string]string{}))`

Now in a new route simply render a template.

```go
m.Get("/hello", func(r martini_amber.Render) {
  r.Amber(200, "hello", map[string]interface{}{
      "title": "Hello World",
  })
})
```

Now create a hello template in __templates/hello.amber__

```
h2 #{title}
```


## Run it!

The server can be started with

`go run server.go`

Now just open your web browser to

`http://127.0.0.1:3000`

If you visit "/hello" you should be greeted with "Hello World"

## Package it!

The last step is to package your server with godeps.

First install Godeps, `go get github.com/tools/godep`

Then simply `godep save` to create a godeps package. This will help ensure reliable builds across all platforms. This step will be vital in getting your app to run on heroku.
