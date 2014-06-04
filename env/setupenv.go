package env

import (
	"log"
)

//SetupEnv controls and/or creates the environment for the program
//Inputs: none at this time
//Outputs: retuns a Config struct with the current configuration.
//Errors: errors will come from env functions.
//Process: Create a new Config struct. Check to see if a configuration file
//already exists.  If no configuration file exists, then complete the Config
//struct with default paramerers and write the configuration file to disk.
//If a configuration file does exist, then read its contents into the Config struct.
func SetupEnv() *Config {
	//set up flags to replace GetHomeDir() and use GetHomeDir() as the default
	RunningConfig := new(Config)
	err := exists(GetHomeDir() + "/.golanganidb/golanganidb.conf")
	if err != nil {
		RunningConfig.ProgramDataPath, RunningConfig.ProgramConfigPath = CreateDir(GetHomeDir())
		initialconfigstring := []string{"client=golanganidb", "clientver=1", "protover=1", "url=http://api.anidb.net", "port=9001", "programdatapath=" + RunningConfig.ProgramDataPath, "programconfigpath=" + RunningConfig.ProgramConfigPath}
		RunningConfig.configtostruct(initialconfigstring)
		WriteConfig(RunningConfig.ProgramConfigPath+"/golanganidb.conf", RunningConfig)
		log.Println("Created new config file at", RunningConfig.ProgramConfigPath+"/golanganidb.conf")
	} else {
		RunningConfig.ReadConfig(GetHomeDir() + "/.golanganidb/golanganidb.conf")
		log.Println("Read existing config file from", GetHomeDir()+"/.golanganidb/golanganidb.conf")
	}
	return RunningConfig
}
