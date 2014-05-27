package env

import (
	"fmt"
	"os"
)

//CreateDir will create the folder sturcture for the AniBD program to run properly.
//Inputs: pathroot is a string that describes where the root folder for the program should be.
//Output: programpath is a string that will be the base folder used to store anime eppisodes and cache AniDB data.
//Output: configpath is a string that will be the base folder used to store configuration data.
//Errors: Any errors created will be related to the filesystem and whether or not directories were created.
//Process: CreateDir takes the pathroot input and creates the necessary directores for golanganidb.  If the directories already exist, they are not recreated.
func CreateDir(pathroot string) (string, string) {
	programpath := pathroot + "/golanganidb"
	configpath := pathroot + "/.golanganidb"
	cachepath := programpath + "/cache"
	makedir(programpath)
	makedir(configpath)
	makedir(cachepath)
	return programpath, configpath
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	//This needs to be simplified - jak
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) == false {
		return false, nil
	}
	return false, err
}

func makedir(path string) {
	pathexists, err := exists(path)
	if pathexists == false {
		patherr := os.Mkdir(path, 0770)
		if patherr == nil {
			fmt.Println(path, " created successfully")
		} else {
			fmt.Println(patherr)
			panic(fmt.Sprintf("%v\n", "Unable to create directory for program."))
		}
	} else {
		fmt.Println(path, " already exists")
	}
}
