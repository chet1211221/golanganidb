package env

import (
	"fmt"
)

func SetupEnv() {
	//need to change process so that config file is checked first for existance
	//then read or created.
	//add programdatapath and programconfig path to Config struct
	programdatapath, programconfigpath := CreateDir(GetHomeDir())
	fmt.Println(programdatapath, programconfigpath)
	//runningconfig := Config{"golanganidb", 1, 1, "http://api.anidb.net", 9001}
	//WriteConfig(programconfigpath+"/golanganidb.conf", &runningconfig)
	//runningconfig = ReadConfig(programconfigpath + "/golanganidb.conf")
	//fmt.Println(runningconfig)

}
