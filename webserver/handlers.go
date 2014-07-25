package webserver

import (
	"github.com/chetbishop/golanganidb/apis/anidb"
	"html/template"
	//"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	p.Title = "Home"
	t, _ := template.ParseFiles("web/home.html", "web/header.html", "web/footer.html")
	t.ExecuteTemplate(w, "home", p)
}

func addSearchHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	p.Title = "Add Anime"
	if r.Method == "GET" {
		t, _ := template.ParseFiles("web/addAnimeSearch.html", "web/header.html", "web/footer.html")
		t.ExecuteTemplate(w, "addAnimeSearch", p)
	} else if r.Method == "POST" {
		r.ParseForm()
		aname := r.FormValue("animename")
		anidbapi.AnimeSearchWrapper(runningConfig, aname)
		t, _ := template.ParseFiles("web/addAnimeSearch.html", "web/header.html", "web/footer.html")
		t.ExecuteTemplate(w, "header", p)
	}
}
