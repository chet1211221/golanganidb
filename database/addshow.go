package database

import (
	"database/sql"
	"log"
)

func AddShow(db *sql.DB, aid string, name string) {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}
	stmt, err := tx.Prepare("insert into shows(aid, name) values(?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(aid, name)
	if err != nil {
		log.Println(err)
	}
	log.Println(result)
	tx.Commit()
}
