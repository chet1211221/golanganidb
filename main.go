package main

import (
	//"github.com/chetbishop/golanganidb/database"
	"github.com/chetbishop/golanganidb/env"
	"github.com/chetbishop/golanganidb/webserver"
)

func main() {
	RunningConfig, DB := env.SetupEnv()
	webserver.WebServer(RunningConfig, DB)
}
