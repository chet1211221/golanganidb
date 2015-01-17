package newznabapi

import (
	"encoding/xml"
	//"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	//"net/url"
	//"os"
	"strings"
)

func (provider *Newznab) GetNewznabResults(searchTerm string) []NewznabSearchResult {
	searchUrl := provider.NewznabUrlQuery(searchTerm)
	queryResultsRaw, err := http.Get(searchUrl)
	if err != nil {
		log.Println(err)
	}

	decoder := xml.NewDecoder(queryResultsRaw.Body)
	var p []NewznabSearchResult
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		// Inspect the type of the token just read.
		switch se := t.(type) {
		case xml.StartElement:
			//log.Println(se.Name.Local)
			// If we just read a StartElement token
			// ...and its name is "page"
			var temp NewznabSearchResult
			if se.Name.Local == "item" {

				// decode a whole chunk of following XML into the
				// variable p which is a Page (se above)
				decoder.DecodeElement(&temp, &se)
				temp.Guid = strings.Split(temp.Guid, "/")[4]
				temp.Url = provider.NewznabUrlGetNzb(temp.Guid)
				//log.Println(len(strings.Split(temp.Guid, ".")))

				p = append(p, temp)
			}

		}

	}
	return p

}
