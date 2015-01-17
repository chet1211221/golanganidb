package background

import (
	"github.com/chetbishop/golanganidb/apis/anidb"
	"github.com/chetbishop/golanganidb/database"
	"time"
)

func UpdateShowsBackground() {
	for _ = range time.Tick(1 * time.Hour) {
		anidbapi.AnimeTitlesCheck(runningConfig)
		shows := database.ListShows(DB)
		for _, show := range shows {
			anidbapi.AnimeDetailsCheck(show.Name, runningConfig)
			result := anidbapi.AnimeDetailsParse(runningConfig.ProgramConfigPath + "/cache/" + show.Aid + ".xml")
			database.PopulateShowWithEpisode(DB, result, show.Lang)
			time.Sleep(1 * time.Second)
		}
	}
}
