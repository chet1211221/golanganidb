package env

import (
	"log"
	"os"
	"strconv"
)

//CreateConfig will take a Config struct and write it to disk.
//Input: configfile is a string that points to the location on disk for the
//configuration file.
//Input: initalconfig points to a Config struct to write to disk.
//Output: Configuration file is written to disk
//Error: If the file already exists then output is written to the console and
// the file on disk is not modified.
//Error: If the config file is unable to be created then then panic is created.
//Process: Checks to see if a configuration file already exists.  If true, then
//nothing is done.  If no configuration file is found, then the file is created
//and the Config struct is written to disk.  If the configuration file is unable
//to be created, then the panic is generated.
func CreateConfig(configfile string, initialconfig *Config) {
	//need to update this to be a general write configuration file rather than
	//just writing the inital configuration.
	err := exists(configfile)
	if err != nil {
		configfilecreated, err := os.Create(configfile)
		if err != nil {
			log.Fatal(err, "Unable to create configuration file for program.")
		}
		defer configfilecreated.Close()
		configfilecreated.WriteString("client=" + initialconfig.Client + "\r\n")
		configfilecreated.WriteString("clientver=" + strconv.Itoa(initialconfig.Clientver) + "\r\n")
		configfilecreated.WriteString("protover=" + strconv.Itoa(initialconfig.Protover) + "\r\n")
		configfilecreated.WriteString("url=" + initialconfig.Url + "\r\n")
		configfilecreated.WriteString("port=" + strconv.Itoa(initialconfig.Port) + "\r\n")
	} else {
		log.Println("Config file already exists")
	}

}
