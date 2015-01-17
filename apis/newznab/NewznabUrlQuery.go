package newznabapi

import (
	//"encoding/xml"
	//"fmt"
	//"io/ioutil"
	"log"
	//"net/http"
	"net/url"
	//"os"
)

func (provider *Newznab) NewznabUrlQuery(searchTerm string) string {
	var queryUrl *url.URL
	queryUrl, err := url.Parse(provider.BaseUrl + "/api")
	if err != nil {
		log.Println(err)
	}
	searchParameters := url.Values{}
	searchParameters.Add("t", "search")
	searchParameters.Add("apikey", provider.ApiKey)
	searchParameters.Add("q", searchTerm)
	searchParameters.Add("o", "xml")
	queryUrl.RawQuery = searchParameters.Encode()
	return queryUrl.String()
}
func (provider *Newznab) NewznabUrlGetNzb(guid string) string {
	var queryUrl *url.URL
	queryUrl, err := url.Parse(provider.BaseUrl + "/api")
	if err != nil {
		log.Println(err)
	}
	searchParameters := url.Values{}
	searchParameters.Add("t", "get")
	searchParameters.Add("apikey", provider.ApiKey)
	searchParameters.Add("id", guid)
	queryUrl.RawQuery = searchParameters.Encode()
	return queryUrl.String()

}
