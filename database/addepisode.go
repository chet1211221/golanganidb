package database

import (
	"database/sql"
	//"github.com/chetbishop/golanganidb/apis/anidb"
	"log"
	"strconv"
)

//episodenumber integer not null primary key, name text, airdate text

func AddEpisode(db *sql.DB, aid int, epno string, name string, airdate string) {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}
	line := "insert into aid_" + strconv.Itoa(aid) + "(episodenumber, name, airdate) values(?, ?, ?)"
	stmt, err := tx.Prepare(line)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(epno, name, airdate)
	if err != nil {
		log.Println(err)
	}
	log.Println(result)
	tx.Commit()
}
