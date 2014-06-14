package webserver

//Page is a struct for information to be displayed on the webpages
//Additional fields will be added as I figure out which elements are needed
type Page struct {
	Title string //Title of webpage
	Body  []byte //Body in byte form.
}
