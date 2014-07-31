//webserver will control the web experience for the user
package webserver

import (
	"database/sql"
	"github.com/chetbishop/golanganidb/apis/anidb"
	"github.com/chetbishop/golanganidb/env"
	"net/http"
)

var runningConfig *env.Config
var DB *sql.DB

func WebServer(runningConfigImport *env.Config, db *sql.DB) {
	runningConfig = runningConfigImport
	DB = db
	mux := http.NewServeMux()
	mux.HandleFunc("/add/search", addSearchHandler)
	mux.HandleFunc("/add/add", addAddHandler)
	mux.HandleFunc("/anime", animeHandler)
	mux.HandleFunc("/", homeHandler)
	fscss := justFilesFilesystem{http.Dir("web/css/")}
	fsjs := justFilesFilesystem{http.Dir("web/js/")}
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(fscss)))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(fsjs)))
	http.ListenAndServe(":8080", mux)
}

type Page struct {
	Title string                             //Title of webpage
	Body  string                             //Body in byte form.
	URL   string                             //URL of the request
	Anime []anidbapi.AnimeTitleSearchResults //various anime information
}
