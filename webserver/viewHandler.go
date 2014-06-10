package webserver

import (
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p := &Page{Title: "golandAniDB Home" + title, Body: []byte("This is a body placeholder.")}
	//if err != nil {
	//	http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	//	return
	//}
	renderTemplate(w, "home", p)
}
