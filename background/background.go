package background

import (
	//"github.com/chetbishop/golanganidb/apis/anidb"
	//"github.com/chetbishop/golanganidb/apis/newznab"
	//"github.com/chetbishop/golanganidb/database"
	"github.com/chetbishop/golanganidb/env"
	//"html/template"
	//"log"
	//"net/http"
	//"strconv"
	//"strings"
	"database/sql"
	//"time"
)

var runningConfig *env.Config
var DB *sql.DB

func Background(runningConfigImport *env.Config, db *sql.DB) {
	runningConfig = runningConfigImport
	DB = db
	//AnimeTitlesCheck(RunningConfig)
	go UpdateShowsBackground()

}
