package database

import (
	"database/sql"
	"github.com/chetbishop/golanganidb/apis/anidb"
	"log"
	"strconv"
)

func ListShows(db *sql.DB) []anidbapi.AnimeTitleSearchResults {
	var (
		aid  int
		name string
	)
	var results []anidbapi.AnimeTitleSearchResults
	rows, err := db.Query("select aid, name from shows")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var result anidbapi.AnimeTitleSearchResults
		err := rows.Scan(&aid, &name)
		if err != nil {
			log.Println(err)
		}
		result.Aid = strconv.Itoa(aid)
		result.Name = name
		err = rows.Err()
		if err != nil {
			log.Println(err)
		}
		results = append(results, result)
	}
	return results
}
