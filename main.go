package main

import (
	//"github.com/chetbishop/golanganidb/database"
	"github.com/chetbishop/golanganidb/env"
	"github.com/chetbishop/golanganidb/webserver"
)

func main() {
	RunningConfig, DB := env.SetupEnv()
	//database.AddTable(DB, "x")
	//database.AddEpisode(DB, 239, "1", "Enter", "2014-03-15")
	webserver.WebServer(RunningConfig, DB)
}
