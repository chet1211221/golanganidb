//webserver will control the web experience for the user
package webserver

import (
	"net/http"
)

func WebServer() {
	fscss := justFilesFilesystem{http.Dir("web/css/")}
	fsjs := justFilesFilesystem{http.Dir("web/js/")}
	http.HandleFunc("/home/", makeHandler(viewHandler))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(fscss)))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(fsjs)))
	http.ListenAndServe(":8080", nil)
}
