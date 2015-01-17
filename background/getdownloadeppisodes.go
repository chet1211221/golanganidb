package background

import (
	//"github.com/chetbishop/golanganidb/apis/anidb"
	"github.com/chetbishop/golanganidb/apis/newznab"
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

func GetDowloadEppisodes() {
	shows := database.ListShows(DB)
	for _, show := range shows {
		//log.Println(show.Aid)
		//log.Println(show.Quality)
		eppisodes := database.ListEpisodes(DB, show.Aid)
		//log.Println(eppisodes)
		for _, eppisode := range eppisodes {
			if eppisode.Status == "download" {
				log.Println(show.Name, eppisode.Epno)
				var results [][]newznabapi.NewznabSearchResult
				for _, prov := range runningConfig.Provider {
					find := prov.GetNewznabResults(show.Name + " " + eppisode.Epno)
					results = append(results, find)
					time.Sleep(1 * time.Second)
				}
				bestresults := newznabapi.FindBestNewznabResult(results, eppisode.Epno, show.Quality)
				//log.Println(newznabapi.NewznabUrlGetNzb(bestresults[0].Guid))
				log.Println(bestresults)
			}

		}
	}

}
