package env

import (
	"fmt"
	"os"
)

//CreateDir will create the folder sturcture for the AniBD program to run
//properly.
//Input: pathroot is a string that describes where the root folder for the
//program should be.
//Output: programpath is a string that will be the base folder used to store
//anime eppisodes and cache AniDB data.
//Output: configpath is a string that will be the base folder used to store
//configuration data.
//Errors: Any errors created will be related to the filesystem and whether or
//not directories were created.
//Process: CreateDir takes the pathroot input and creates the necessary
//directores for golanganidb.  If the directories already exist, they are not
//recreated.
func CreateDir(pathroot string) (string, string) {
	programpath := pathroot + "/golanganidb"
	configpath := pathroot + "/.golanganidb"
	cachepath := programpath + "/cache"
	makedir(programpath)
	makedir(configpath)
	makedir(cachepath)
	return programpath, configpath
}

//exists returns whether the given file or directory exists or not in the form
//of a bool
//Input: path is a string that points to the file or directory to check
//Output: bool for if the file or directory exists.
//Output: error is outputted if the file does not exist or the operating sytem
//had an error accessing the path or file.
//Error: error message return is related to the file or directory not existing
//and possibly other details if the operating system is blocking access to the
//file or directory.
//Process: check the path with os.stat.  if no error is returned the file exists
//and return true.  If an error is returned then check to see if the error is
//known to report that a file or directory does not exits and return the error.
//If the error is not known then return the error.
func exists(path string) (bool, error) {
	//This needs to be simplified - jak
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) == false {
		return false, err
	}
	return false, err
}

//makedir creates a directory with execute permissions for owner and group.
//Input: path is a string that points to the directory to create
//Output: Text is printed to the console indicating if the creation was
//successfull.
//Error: The script panics when it is unable to create a directory.
//Process: take the input and create a directory.
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
