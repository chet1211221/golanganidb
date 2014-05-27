package env

import (
	//"os"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

//Read the preexisting config file into a Config struct to be used later on in the program. Need to work on adding the read config into the struck
func ReadConfig(configfile string) *Config {
	configfilebyte, err := ioutil.ReadFile(configfile)
	if err != nil {
		fmt.Println(err)
	}
	configfilestring := strings.SplitAfter(string(configfilebyte[:]), "\r\n")
	configstruct := new(Config)
	configstruct.configtostruct(configfilestring)
	//fmt.Println(configfilestring)
	//configparams := []string{"client", "clientver", "protover", "url", "port"}
	//for a := range configparams {
	//b := stringsearch(configfilestring, configparams[a])
	//fmt.Println(b)

	//}
	return configstruct
}

//what happens when you run out of letters? You can only have 26 variables bro -- jak
func (h *Config) configtostruct(configstring []string) *Config {
	h.Client = stringsearch(configstring, "client")
	h.Clientver, _ = strconv.Atoi(stringsearch(configstring, "clientver"))
	h.Protover, _ = strconv.Atoi(stringsearch(configstring, "protover"))
	h.Url = stringsearch(configstring, "url")
	h.Port, _ = strconv.Atoi(stringsearch(configstring, "port"))
	return h
}

func stringsearch(searchstrings []string, substring string) string {
	//See comment about 'h' above in configtostruct() - only 24 remaining! -- jak
	for g := range searchstrings {
		result := strings.Contains(searchstrings[g], substring)
		if result == true {
			configlinesplit := strings.SplitAfter(searchstrings[g], "=")
			//configlinesplit = strings.TrimSpace(configlinesplit)
			//fmt.Println(strings.TrimSpace(configlinesplit[1]))
			return strings.TrimSpace(configlinesplit[1])
		}
	}
	//if I didn't like you so much I would slap the taste from your mouth -- jak
	return nil
}
