//webserver will control the web experience for the user
package webserver

import (
	"net/http"
)

func WebServer() {
	http.HandleFunc("/home/", makeHandler(viewHandler))
	http.ListenAndServe(":8080", nil)
}
