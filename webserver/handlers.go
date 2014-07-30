package webserver

import (
	"github.com/chetbishop/golanganidb/apis/anidb"
	"github.com/chetbishop/golanganidb/database"
	"html/template"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	p.Title = "Home"
	p.Anime = database.ListShows(DB)
	t, _ := template.ParseFiles("web/home.html", "web/header.html", "web/footer.html")
	t.ExecuteTemplate(w, "header", p)
	t.ExecuteTemplate(w, "home", p)
	t.ExecuteTemplate(w, "footer", p)

}

func addSearchHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	p.Title = "Add Anime"
	if r.Method == "GET" {
		t, _ := template.ParseFiles("web/addAnimeSearch.html", "web/header.html", "web/footer.html")
		t.ExecuteTemplate(w, "header", p)
		t.ExecuteTemplate(w, "addAnimeSearch", p)
		t.ExecuteTemplate(w, "footer", p)
	} else if r.Method == "POST" {
		r.ParseForm()
		aname := r.FormValue("animename")
		p.Anime = anidbapi.AnimeSearchWrapper(runningConfig, aname)
		t, _ := template.ParseFiles("web/addAnimeResults.html", "web/header.html", "web/footer.html")
		t.ExecuteTemplate(w, "header", p)
		t.ExecuteTemplate(w, "addAnimeResults", p)
		t.ExecuteTemplate(w, "footer", p)

	}
}
func addAddHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for x := range r.Form {
		database.AddShow(DB, r.FormValue(x), x)
		anidbapi.AnimeDetailsGet(r.FormValue(x), runningConfig)
		result := anidbapi.AnimeDetailsParse(runningConfig.ProgramConfigPath + "/cache/" + r.FormValue(x) + ".xml")
		log.Println(result)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
