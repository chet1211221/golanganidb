package main

import (
	"github.com/chetbishop/golanganidb/apis/anidb"
	"github.com/chetbishop/golanganidb/env"
	"github.com/chetbishop/golanganidb/webserver"
)

func main() {
	RunningConfig := env.SetupEnv()
	go webserver.WebServer()
	animexml := anidbapi.AnimeParse(RunningConfig.ProgramConfigPath + "/cache/anime-titles.xml")
	anidbapi.AnimeSearch(animexml, "Naruto", "en")

}
