package newznabapi

type NewznabSearchResult struct {
	Title string `xml:"title"`
	Guid  string `xml:"guid"`
	Url   string
	//Link  string `xml:"link"`
}
type Newznab struct {
	Name    string
	BaseUrl string
	ApiKey  string
}
