package database

import (
	"database/sql"
	"log"
)

func AddShow(db *sql.DB, aid string, name string, description string, quality string, lang string) {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}
	stmt, err := tx.Prepare("insert into shows(aid, name, description, quality, lang) values(?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(aid, name, description, quality, lang)
	if err != nil {
		log.Println(err)
	}
	//log.Println(result)
	tx.Commit()
}
