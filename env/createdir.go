//env will control the filesystem environment
package env

import (
	"log"
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
	cachepath := configpath + "/cache"
	makedir(programpath)
	makedir(configpath)
	makedir(cachepath)
	return programpath, configpath
}

//exists returns whether the given file or directory exists or not in the form
//of an error message.
//Input: path is a string that points to the file or directory to check
//Output: error message from os.Stat
//Process: check the path with os.stat.  if no error is returned the file exists
//and retur nil.  If an error is returned then return the error.
func exists(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	} else {
		return nil
	}
}

//makedir creates a directory with execute permissions for owner and group.
//Input: path is a string that points to the directory to create
//Output: Text is printed to the console indicating if the creation was
//successfull.
//Error: The script panics when it is unable to create a directory.
//Process: take the input and create a directory.
func makedir(path string) {
	err := exists(path)
	if err != nil {
		patherr := os.Mkdir(path, 0770)
		if patherr != nil {
			log.Fatal("Unable to create directory for program.")
		} else {
			log.Println(path, " created successfully")
		}
	} else {
		log.Println(path, " already exists")
	}
}
