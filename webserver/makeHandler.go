package webserver

import (
	"net/http"
	"regexp"
)

//makeHandler is a function literal that checks for a valid url and returns a
//function of type http.HandlerFunc and the title string.  To be used with
//viewHandler().
//Input:
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

//validPath is a regular expression that must compile in order to check for
//valid url paths
var validPath = regexp.MustCompile("^/(home|show)/([a-zA-Z0-9]+)$")
