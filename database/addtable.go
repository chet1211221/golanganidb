package database

import (
	"database/sql"
	"log"
)

func AddTable(db *sql.DB, name string) {
	stmt, err := db.Prepare(`
create table ? (aid integer not null primary key, name text);
`)
	defer stmt.Close()
	result, err := stmt.Exec(name)
	if err != nil {
		log.Printf("%q: %s\n", err)
	}
	log.Println(result)

}
