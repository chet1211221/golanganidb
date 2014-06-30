package main

import (
	//"fmt"
	"github.com/chetbishop/golanganidb/env"
	"github.com/chetbishop/golanganidb/webserver"
)

func main() {
	env.SetupEnv()
	webserver.WebServer()
}
