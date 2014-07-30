package anidbapi

import (
	"encoding/xml"
	"github.com/chetbishop/golanganidb/env"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	//"strings"
	//"time"
)

//AnimeTitlesGet downloades the anime-titles.xml file from AniDB.
func AnimeDetailsGet(aid string, runningconfig *env.Config) {
	url := "http://api.anidb.net:" + strconv.Itoa(runningconfig.Port) +
		"/httpapi?request=anime&client=" + runningconfig.Client +
		"&clientver=" + strconv.Itoa(runningconfig.Clientver) + "&protover=" +
		strconv.Itoa(runningconfig.Protover) + "&aid=" + aid
	log.Println("downloading anime details")
	log.Println(url)
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	AnimeDetails, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Println(err)
	}
	savelocation := runningconfig.ProgramConfigPath + "/cache/" + aid + ".xml"
	os.Remove(savelocation)
	ioutil.WriteFile(savelocation, AnimeDetails, 0600)

}

//AnimeDetailsCheck checks to see if the xml file for each AID from AniDB has
//been downloaded in the last 24 hours.  AnimeTitlesCheck will download
//the xml if the file is older than 24 hours or has not been
//downloaded.
func AnimeDetailsCheck(aid string, runningConfig *env.Config) {
	savelocation := runningConfig.ProgramConfigPath + "/cache/" + aid + ".xml"
	animedetails, err := os.Stat(savelocation)
	if err != nil {
		log.Println("xml for anime does not exist ... Downloading")
		AnimeDetailsGet(aid, runningConfig)
	} else {
		log.Println("checking to see if 24 hours has passed since last anime list download")
		daypassed := testTime24h(animedetails.ModTime())
		if daypassed == true {
			log.Println("Downloading ")
			AnimeDetailsGet(aid, runningConfig)
		}
	}

}

//AnimeParse takes anime-titles.xml and produces an AnimeTitles struct with
//all revelent information
func AnimeDetailsParse(xmlFilestring string) AnimeDetails {
	xmlFile, err := os.Open(xmlFilestring)
	if err != nil {
		log.Println("Error opening file:", err)
	}
	defer xmlFile.Close()
	log.Println("opened file")
	b, _ := ioutil.ReadAll(xmlFile)

	var details AnimeDetails
	xml.Unmarshal(b, &details)
	return details
}
