package database

import (
	"database/sql"
	"log"
)

func AddShowTable(db *sql.DB, name string) {
	line := "create table IF NOT EXISTS aid_" + name + "(episodenumber integer not null primary key, name text, airdate text, status text, quality text)"
	stmt, err := db.Prepare(line)
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Printf("%q: %s\n", err)
	}
	//log.Println(result)
}
