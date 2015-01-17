package newznabapi

import (
	"encoding/xml"
)

func NewznabResultsParse(newzNabQueryResults []byte) NewznabSearchResult {
	var details NewznabSearchResult
	xml.Unmarshal(newzNabQueryResults, &details)
	return details
}
