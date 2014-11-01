package env

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func CreateDatabase(RunningConfig *Config) *sql.DB {
	dbname := "database"
	dbloc := RunningConfig.ProgramConfigPath + "/" + dbname
	err := os.MkdirAll(dbloc, 0777)
	if err != nil {
		log.Fatal(err)
	}
	DB, err := sql.Open("sqlite3", dbloc+"/"+dbname+".db")
	if err != nil {
		log.Fatal(err)
	}
	sql := `
create table shows (aid integer not null primary key, name text, description text, lang text, quality text);
`
	_, err = DB.Exec(sql)
	if err != nil {
		log.Printf("%q: %s\n", err, sql)
	}
	return DB
}
