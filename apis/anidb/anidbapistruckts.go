package anidbapi

import ()

//AnimeTitles is the struct for the top level of anime-titles.xml
type AnimeTitles struct {
	AnimeList []Anime `xml:"anime"` //from anime-titles.xml
}

//Anime is the struct for the anime level of anime-titles.xml
type Anime struct {
	Aid   int          `xml:"aid,attr"` //from anime-titles.xml
	Title []AnimeTitle `xml:"title"`    //from anime-titles.xml
}

//AnimeTitle is the struct for the title lines of anime-titles.xml
type AnimeTitle struct {
	Name      string `xml:",chardata"` //from anime-titles.xml
	AnimeType string `xml:"type,attr"` //from anime-titles.xml
	Lang      string `xml:"lang,attr"` //from anime-titles.xml
}

//AnimeTitleSearchResults is the struct for Anime title search results.
type AnimeTitleSearchResults struct {
	Name    string
	Aid     string
	EpName  string
	Epno    string
	Airdate string
	Status  string
	Lang    string
	Quality string
}
type AnimeDetails struct {
	Aid          int      `xml:"id,attr"`
	Episodes     Episodes `xml:"episodes"`
	EpisodeCount int      `xml:"episodecount"`
	StartDate    string   `xml:"startdate"`
	EndDate      string   `xml:"enddate"`
	Description  string   `xml:"description"`
}
type Episodes struct {
	Episode []Episode `xml:"episode"`
}
type Episode struct {
	Epno    string       `xml:"epno"`
	Airdate string       `xml:"airdate"`
	Title   []AnimeTitle `xml:"title"`
	Status  string
}
