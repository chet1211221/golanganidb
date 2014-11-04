package main

import (
	"github.com/chetbishop/golanganidb/background"
	"github.com/chetbishop/golanganidb/env"
	"github.com/chetbishop/golanganidb/webserver"
)

func main() {
	RunningConfig, DB := env.SetupEnv()
	go background.Background(RunningConfig, DB)
	webserver.WebServer(RunningConfig, DB)

}
