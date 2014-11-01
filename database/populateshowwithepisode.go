package database

import (
	"database/sql"
	"github.com/chetbishop/golanganidb/apis/anidb"
	//"log"
	"strconv"
)

//database.AddEpisode(DB, 239, "1", "Enter", "2014-03-15")
func PopulateShowWithEpisode(db *sql.DB, animedetails anidbapi.AnimeDetails, animelang string) {
	var aid int
	var epno int
	var name string
	var airdate string
	var status string
	for _, episode := range animedetails.Episodes.Episode {
		aid = animedetails.Aid
		epno, _ = strconv.Atoi(episode.Epno)
		airdate = episode.Airdate
		status = "skipped"
		for _, title := range episode.Title {
			if title.Lang == animelang {
				name = title.Name
			}
		}
		if epno != 0 {
			//log.Println(aid, epno, name, airdate)
			AddEpisode(db, aid, strconv.Itoa(epno), name, airdate, status)
		}
	}
}

//func PopulateShowExtraInfo(db *sql.DB, animedetails anidbapi.AnimeDetails)
