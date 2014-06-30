package webserver

import (
	//"fmt"
	"net/http"
)

//viewHandler creates a handler to return to
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p := &Page{Title: "golandAniDB Home " + title, Body: "This is a body placeholder."}
	//if err != nil {
	//	http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	//	return
	//}
	renderTemplate(w, "view", p)
}
