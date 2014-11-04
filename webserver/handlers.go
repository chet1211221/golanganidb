package webserver

import (
	"github.com/chetbishop/golanganidb/apis/anidb"
	"github.com/chetbishop/golanganidb/apis/newznab"
	"github.com/chetbishop/golanganidb/database"
	"github.com/chetbishop/golanganidb/env"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
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
		animelang := r.FormValue("animelang")
		p.Lang = animelang
		p.Anime = anidbapi.AnimeSearchWrapper(runningConfig, aname, animelang)
		t, _ := template.ParseFiles("web/addAnimeResults.html", "web/header.html", "web/footer.html")
		t.ExecuteTemplate(w, "header", p)
		t.ExecuteTemplate(w, "addAnimeResults", p)
		t.ExecuteTemplate(w, "footer", p)

	}
}
func addAddHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	animelang := r.FormValue("animelang")
	animequality := r.FormValue("animequality")
	for x := range r.Form["titles"] {
		s := strings.Split(r.Form["titles"][x], ",")
		//log.Println("aid: ", s[0], "title: ", s[1])
		anidbapi.AnimeDetailsCheck(s[0], runningConfig)
		result := anidbapi.AnimeDetailsParse(runningConfig.ProgramConfigPath + "/cache/" + s[0] + ".xml")
		database.AddShow(DB, s[0], s[1], result.Description, animequality, animelang)
		database.AddShowTable(DB, s[0])
		database.PopulateShowWithEpisode(DB, result, animelang)

	}
	http.Redirect(w, r, "/", http.StatusFound)
}
func animeHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	aid := r.URL.Path[len("/anime/") : len(r.URL.Path)-1]
	p.URL = "/anime/" + aid
	p.Title, p.Body = database.GetShowNameDescription(DB, aid)
	p.Anime = database.ListEpisodes(DB, aid)
	//result := anidbapi.AnimeDetailsParse(runningConfig.ProgramConfigPath + "/cache/" + aid + ".xml")
	//log.Println(result.Description)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("web/animeAid.html", "web/header.html", "web/footer.html")
		t.ExecuteTemplate(w, "header", p)
		t.ExecuteTemplate(w, "animeAid", p)
		t.ExecuteTemplate(w, "footer", p)
	} else if r.Method == "POST" {
		r.ParseForm()
		status := r.FormValue("status")
		for x := range r.Form["update"] {
			aidint, err := strconv.Atoi(aid)
			if err != nil {
				log.Println(err)
			}
			database.UpdateEpisodeStatus(DB, aidint, r.Form["update"][x], status)
		}
		http.Redirect(w, r, "/anime/"+aid+"/", http.StatusFound)
	}

}

func providersHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	p.Title = "NewzNab Providers"
	p.RunningConfig = runningConfig
	if r.Method == "GET" {
		t, _ := template.ParseFiles("web/settingsProviders.html", "web/header.html", "web/footer.html")
		t.ExecuteTemplate(w, "header", p)
		t.ExecuteTemplate(w, "settingsProviders", p)
		t.ExecuteTemplate(w, "footer", p)
	} else if r.Method == "POST" {
		r.ParseForm()
		var x newznabapi.Newznab
		x.Name = r.FormValue("name")
		x.BaseUrl = r.FormValue("baseurl")
		x.ApiKey = r.FormValue("apikey")
		p.RunningConfig.Provider = append(p.RunningConfig.Provider, x)
		t, _ := template.ParseFiles("web/settingsProviders.html", "web/header.html", "web/footer.html")
		t.ExecuteTemplate(w, "header", p)
		t.ExecuteTemplate(w, "settingsProviders", p)
		t.ExecuteTemplate(w, "footer", p)
		env.WriteConfig(runningConfig.ProgramConfigPath+"/golanganidb.conf", runningConfig)
	}

}
