package newznabapi

import (
	//"encoding/xml"
	//"fmt"
	//"io/ioutil"
	//"log"
	//"net/http"
	//"net/url"
	//"os"
	"strings"
)

func FindBestNewznabResult(results [][]NewznabSearchResult, search string, quality string) []NewznabSearchResult {
	var combinedresults []NewznabSearchResult
	var returnresults []NewznabSearchResult
	for _, result := range results {
		for _, res := range result {
			combinedresults = append(combinedresults, res)
		}
	}
	for _, newresult := range combinedresults {
		if strings.Contains(newresult.Title, search) == true {
			qualitysearch := strings.TrimSuffix(quality, "p")
			if strings.Contains(newresult.Title, qualitysearch) == true {
				returnresults = append(returnresults, newresult)
			}
		}

	}
	return returnresults
}
