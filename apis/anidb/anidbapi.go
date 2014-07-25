package anidbapi

import (
	"encoding/xml"
	"github.com/chetbishop/golanganidb/env"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//AnimeTitlesCheck checks to see if the anime-titles.xml file from AniDB has
//been downloaded in the last 24 hours.  AnimeTitlesCheck will download
//anime-titles.xml if the file is older than 24 hours or has not been
//downloaded.
func AnimeTitlesCheck(RunningConfig *env.Config) {
	savelocation := RunningConfig.ProgramConfigPath + "/cache/anime-titles.xml"
	anititles, err := os.Stat(savelocation)
	if err != nil {
		log.Println("anime-titles.dat does not exist ... Downloading")
		AnimeTitlesGet(savelocation)
	} else {
		log.Println("checking to see if 24 hours has passed since last anime list download")
		daypassed := testTime24h(anititles.ModTime())
		if daypassed == true {
			log.Println("Downloading ")
			AnimeTitlesGet(savelocation)
		}
	}

}

//AnimeTitlesGet downloades the anime-titles.xml file from AniDB.
func AnimeTitlesGet(savelocation string) {
	log.Println("downloading anime titles")
	res, err := http.Get("http://anidb.net/api/anime-titles.xml.gz")
	if err != nil {
		log.Println(err)
	}
	animelist, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Println(err)
	}
	os.Remove(savelocation)
	ioutil.WriteFile(savelocation, animelist, 0600)

}

//testTime24h tests to see if 24 hours has passed between two times.
func testTime24h(modtime time.Time) bool {
	timediff := time.Now().Sub(modtime).Hours()
	var result bool
	if timediff > 24 {
		result = true
	} else {
		result = false
	}
	return result
}

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

//AnimeParse takes anime-titles.xml and produces an AnimeTitles struct with
//all revelent information
func AnimeParse(xmlFilestring string) AnimeTitles {
	xmlFile, err := os.Open(xmlFilestring)
	if err != nil {
		log.Println("Error opening file:", err)
	}
	defer xmlFile.Close()
	log.Println("opened file")
	b, _ := ioutil.ReadAll(xmlFile)

	var titles AnimeTitles
	xml.Unmarshal(b, &titles)
	return titles
}

//AnimeSearch will seach an AnimeTitles struct for an anime name and language.
//It will return the aid number(s) and anime name(s) from the AnimeTitles struct.
func AnimeSearch(animeTitlesStruct AnimeTitles, animename string, animelang string) [][]string {
	var searchresults [][]string
	for _, aid := range animeTitlesStruct.AnimeList {
		for x, title := range aid.Title {
			if AnimeTitleCompare(aid.Title[x], animename, animelang) == true {
				searchresults = append(searchresults, []string{strconv.Itoa(aid.Aid), title.Name})
			}
		}
	}
	return searchresults
}

func AnimeTitleCompare(animetitle AnimeTitle, animename string, animelang string) bool {
	structname := strings.ToLower(animetitle.Name)
	structlang := strings.ToLower(animetitle.Lang)
	animename = strings.ToLower(animename)
	animelang = strings.ToLower(animelang)

	if strings.Contains(structname, animename) == true {
		if structlang == animelang {
			return true
		}
	}
	return false
}
func AnimeSearchWrapper(RunningConfig *env.Config, animename string) [][]string {
	AnimeTitlesCheck(RunningConfig)
	animexml := AnimeParse(RunningConfig.ProgramConfigPath + "/cache/anime-titles.xml")
	results := AnimeSearch(animexml, animename, "en")
	return results
}
