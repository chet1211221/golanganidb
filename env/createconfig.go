package env

import (
	"log"
	"os"
	"strconv"
)

//CreateConfig will take a Config struct and write it to disk.
//Input: configfile is a string that points to the location on disk for the
//configuration file.
//Input: configstruct points to a Config struct to write to disk.
//Output: Configuration file is written to disk
//Error: Errors returned will be related to creating the file on the OS. If the
//configuration file cannot be written, changes made while the program is
//running will not be saved
//Process: Created the configuration file and write the Config struct to disk.
func WriteConfig(configfile string, configstruct *Config) {
	configfilecreated, err := os.OpenFile(configfile, os.O_RDWR|os.O_CREATE, 0550)
	if err != nil {
		log.Println(err, "Unable to create configuration file for program.")
	}
	defer configfilecreated.Close()
	configfilecreated.WriteString("client=" + configstruct.Client + "\r\n")
	configfilecreated.WriteString("clientver=" + strconv.Itoa(configstruct.Clientver) + "\r\n")
	configfilecreated.WriteString("protover=" + strconv.Itoa(configstruct.Protover) + "\r\n")
	configfilecreated.WriteString("url=" + configstruct.Url + "\r\n")
	configfilecreated.WriteString("port=" + strconv.Itoa(configstruct.Port) + "\r\n")
}
