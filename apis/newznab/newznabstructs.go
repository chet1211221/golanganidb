package newznabapi

type NewznabSearchResult struct {
	Title string `xml:"title"`
	Guid  string `xml:"guid"`
}
type Newznab struct {
	BaseUrl string
	ApiKey  string
}
