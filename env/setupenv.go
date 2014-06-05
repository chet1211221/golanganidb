package env

import (
	"flag"
	"log"
)

//SetupEnv controls and/or creates the environment for the program
//Input: programRootPath is a string stating the root path for the program.
//Defaults to user's home directory
//Input: configurationFilePtr is a string stating the path to the configuration
//file.  Defaults to user's home directory .golanganidb/golanganidb.conf.
//Outputs: retuns a Config struct with the current configuration.
//Errors: errors will come from env functions.
//Process: Create a new Config struct. Check to see if a configuration file
//already exists.  If no configuration file exists, then complete the Config
//struct with default paramerers and write the configuration file to disk.
//If a configuration file does exist, then read its contents into the Config struct.
func SetupEnv() *Config {

	programRootPath := flag.String("root-path", GetHomeDir(), "Root path for all program data")
	flag.Parse()
	configurationFilePtr := flag.String("configuration-file", *programRootPath+"/.golanganidb/golanganidb.conf", "Location of custom configuration file")
	flag.Parse()
	RunningConfig := new(Config)
	err := exists(*configurationFilePtr)
	if err != nil {
		RunningConfig.ProgramDataPath, RunningConfig.ProgramConfigPath = CreateDir(*programRootPath)
		initialconfigstring := []string{"client=golanganidb", "clientver=1", "protover=1", "url=http://api.anidb.net", "port=9001", "programdatapath=" + RunningConfig.ProgramDataPath, "programconfigpath=" + RunningConfig.ProgramConfigPath}
		RunningConfig.configtostruct(initialconfigstring)
		WriteConfig(RunningConfig.ProgramConfigPath+"/golanganidb.conf", RunningConfig)
		log.Println("Created new config file at", RunningConfig.ProgramConfigPath+"/golanganidb.conf")
	} else {
		RunningConfig.ReadConfig(*configurationFilePtr)
		log.Println("Read existing config file from", GetHomeDir()+"/.golanganidb/golanganidb.conf")
	}
	return RunningConfig
}
