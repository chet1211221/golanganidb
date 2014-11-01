package database

import (
	"database/sql"
	"github.com/chetbishop/golanganidb/apis/anidb"
	"log"
	//"strconv"
)

func ListEpisodes(db *sql.DB, aid string) []anidbapi.AnimeTitleSearchResults {
	var (
		epno    string
		name    string
		airdate string
		status  string
		quality string
	)
	var results []anidbapi.AnimeTitleSearchResults
	line := "select episodenumber, name, airdate, status, quality from aid_" + aid
	rows, err := db.Query(line)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var result anidbapi.AnimeTitleSearchResults
		err := rows.Scan(&epno, &name, &airdate, &status, &quality)
		if err != nil {
			log.Println(err)
		}
		result.Epno = epno
		result.EpName = name
		result.Airdate = airdate
		result.Status = status
		result.Quality = quality
		err = rows.Err()
		if err != nil {
			log.Println(err)
		}
		results = append(results, result)
	}
	return results
}
