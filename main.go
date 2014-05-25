package main

import (
	"fmt"
	"github.com/chet1211221/golanganidb/env"
)

func main() {
	programdatapath, programconfigpath := env.CreateDir(env.GetHomeDir())
	fmt.Println(programdatapath)
	fmt.Println(programconfigpath)
	initialconfig := env.Config{"golanganidb", 1, 1, "http://api.anidb.net", 9001}
	env.CreateConfig(programconfigpath+"/golanganidb.conf", &initialconfig)
	fmt.Printf("%v\n", initialconfig)
	//runningconfig := new(setupenv.Config)
	env.ReadConfig(programconfigpath + "/golanganidb.conf")

}
