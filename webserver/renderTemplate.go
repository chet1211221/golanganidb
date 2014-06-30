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

//templates is a variable containing *Template structs.
//Input: The templates to use with the webserver.
//Error: If the template files referenced do not parse correctly, a panic is
//generated.
//Process: The template files are parsed by template.ParseFiles and then checked
//with template.Must.
var templates = template.Must(template.ParseFiles("web/view.html"))
