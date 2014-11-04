package background

import (
	"github.com/chetbishop/golanganidb/apis/anidb"
	//"github.com/chetbishop/golanganidb/apis/newznab"
	"github.com/chetbishop/golanganidb/database"
	//"github.com/chetbishop/golanganidb/env"
	//"html/template"
	"log"
	//"net/http"
	//"strconv"
	//"strings"
	//"database/sql"
	"time"
)

func UpdateShowsBackground() {
	for _ = range time.Tick(1 * time.Minute) {
		anidbapi.AnimeTitlesCheck(runningConfig)
		shows := database.ListShows(DB)
		for x, show := range shows {
			anidbapi.AnimeDetailsCheck(show.Name, runningConfig)
			result := anidbapi.AnimeDetailsParse(runningConfig.ProgramConfigPath + "/cache/" + show.Aid + ".xml")
			database.PopulateShowWithEpisode(DB, result, show.Lang)
			log.Println(x)
			log.Println(show.Lang)
			time.Sleep(1 * time.Second)
		}
	}
}
