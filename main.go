package main

import (
	"github.com/chetbishop/golanganidb/apis/anidb"
	"github.com/chetbishop/golanganidb/env"
	"github.com/chetbishop/golanganidb/webserver"
	"log"
)

func main() {
	RunningConfig := env.SetupEnv()
	webserver.WebServer()
}
