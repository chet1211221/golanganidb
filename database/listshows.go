package database

import (
	"database/sql"
	"github.com/chetbishop/golanganidb/apis/anidb"
	"log"
	"strconv"
)

func GetShowNameDescription(db *sql.DB, aid string) (string, string) {
	var (
		dbaid             int
		name              string
		description       string
		resultname        string
		resultdescription string
	)
	rows, err := db.Query("select aid, name, description from shows")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err := rows.Scan(&dbaid, &name, &description)
		if err != nil {
			log.Println(err)
		}
		if strconv.Itoa(dbaid) == aid {
			resultname = name
			resultdescription = description
		}
		err = rows.Err()
		if err != nil {
			log.Println(err)
		}
	}
	return resultname, resultdescription
}
func ListShows(db *sql.DB) []anidbapi.AnimeTitleSearchResults {
	var (
		aid     int
		name    string
		quality string
	)
	var results []anidbapi.AnimeTitleSearchResults
	rows, err := db.Query("select aid, name, quality from shows")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var result anidbapi.AnimeTitleSearchResults
		err := rows.Scan(&aid, &name, &quality)
		if err != nil {
			log.Println(err)
		}
		result.Aid = strconv.Itoa(aid)
		result.Name = name
		result.Quality = quality
		err = rows.Err()
		if err != nil {
			log.Println(err)
		}
		results = append(results, result)
	}
	return results
}
