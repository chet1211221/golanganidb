package webserver

import (
	"html/template"
	"net/http"
)

//renderTemplate using html/template to create a templates from a saved .html
//file.
//Inputs: w is a http.ResponseWriter that is an interface used by an HTTP
//handler to construct an HTTP response.
//Input: tmpl is a string that points to the .html file for the template.
//Adding the .html is not required and has the potenial to mess up the function.
//Input: p is a Page struct.
//Output: tmpl the template with the info from the Page struct is written to
//w the http.ResponseWriter.
//Error: If the template is not found, then the http code 500 is written to
//w the http.ResponseWriter.
//Process: html/templates ExecuteTemple is called using the approperiate
//ResponseWriter, the tmpl string plus the .html file extention, and the
//approperiate Page struct. If ExecuteTemplate reruns any errors, then a http
//code 500 is written to the http.ResponseWriter.
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var templates = template.Must(template.ParseFiles("index.html", "view.html"))
