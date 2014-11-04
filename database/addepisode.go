package database

import (
	"database/sql"
	//"github.com/chetbishop/golanganidb/apis/anidb"
	"log"
	"strconv"
)

//episodenumber integer not null primary key, name text, airdate text

func AddEpisode(db *sql.DB, aid int, epno string, name string, airdate string, status string) {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}
	line := "insert into aid_" + strconv.Itoa(aid) + "(episodenumber, name, airdate, status) values(?, ?, ?, ?)"
	stmt, err := tx.Prepare(line)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(epno, name, airdate, status)
	if err != nil {
		log.Println(err)
	}
	//log.Println(result)
	tx.Commit()
}

func UpdateEpisodeName(db *sql.DB, aid int, epno string, name string) {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}
	line := "update aid_" + strconv.Itoa(aid) + " SET name=" + "'" + name + "'" + " WHERE episodenumber=" + epno
	stmt, err := tx.Prepare(line)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Println(err)
	}
	//log.Println(result)
	tx.Commit()
}

func UpdateEpisodeStatus(db *sql.DB, aid int, epno string, status string) {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}
	line := "update aid_" + strconv.Itoa(aid) + " SET status=" + "'" + status + "'" + " WHERE episodenumber=" + epno
	stmt, err := tx.Prepare(line)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Println(err)
	}
	//log.Println(result)
	tx.Commit()
}
