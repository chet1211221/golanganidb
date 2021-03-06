package env

import (
	//"fmt"
	"github.com/chetbishop/golanganidb/apis/newznab"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

//ReadConfig will turn an existing configuration file into a Config struct.
//Input: configfile is a sting that points to an existing configuration file.
//Ouptput: A struct of type Config
//Error: An error is returned if the existing configuration file cannot be read.
//Process: Read the existing configuration file as a byte.  Create a new Config
//struct. Split the byte file into lines and pass to the configtostruct function.
func (configstruct *Config) ReadConfig(configfile string) *Config {
	configfilebyte, err := ioutil.ReadFile(configfile)
	if err != nil {
		log.Println(err)
	}
	configfilestring := strings.SplitAfter(string(configfilebyte[:]), "\r\n")
	configstruct.configtostruct(configfilestring)
	return configstruct
}

//configtostruct takes config file string and creates a config struct.
//Input: configstring is a []string containing the configuration parameters.
//Output: A struct of type Config is returned
//Errors: No errors are returned, but a blank struct could be returned if the
//configstring is invalid.
//Process: Use the stringsearch function to match the lines of the configuration
//file with the members of the struct.
func (newconfig *Config) configtostruct(configstring []string) *Config {
	newconfig.Client = stringsearch(configstring, "client")
	newconfig.Clientver, _ = strconv.Atoi(stringsearch(configstring, "clientver"))
	newconfig.Protover, _ = strconv.Atoi(stringsearch(configstring, "protover"))
	newconfig.Url = stringsearch(configstring, "url")
	newconfig.Port, _ = strconv.Atoi(stringsearch(configstring, "port"))
	newconfig.ProgramConfigPath = stringsearch(configstring, "programconfigpath")
	newconfig.ProgramDataPath = stringsearch(configstring, "programdatapath")
	providerscount, err := strconv.Atoi(stringsearch(configstring, "providers"))
	if err != nil {
		log.Println(err)
	}
	for i := 0; i <= providerscount-1; i++ {
		var x newznabapi.Newznab
		x.Name = stringsearch(configstring, "provider"+strconv.Itoa(i)+"Name")
		x.BaseUrl = stringsearch(configstring, "provider"+strconv.Itoa(i)+"Url")
		x.ApiKey = stringsearch(configstring, "provider"+strconv.Itoa(i)+"ApiKey")
		newconfig.Provider = append(newconfig.Provider, x)
	}
	return newconfig
}

//stringsearch will search a configuration file []string for a specific line and
//return the information after the equals sign.
//Input: searchstring is the []string of the configuration file.
//Input: substring is the line in the configuration file to seach for in form of
//string.
//Output: The information after the equal sign in the configuration file in the
//form of a string.
//Errors: If the line of the configuration file is not found, then nil is returned.
func stringsearch(searchstrings []string, substring string) string {
	for searchstringline := range searchstrings {
		result := strings.Contains(searchstrings[searchstringline], substring)
		if result == true {
			configlinesplit := strings.SplitAfter(searchstrings[searchstringline], "=")
			return strings.TrimSpace(configlinesplit[1])
		}
	}
	//see http://golang.org/ref/spec#The_zero_value "" is the correct zero value
	//for a string
	return ""
}
