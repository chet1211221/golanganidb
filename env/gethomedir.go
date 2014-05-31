package env

import (
	"fmt"
	"os/user"
)

//GetHomeDir will return the home directory for the user executing the program.
//Output: The full path of the user's home directory.
//Errors: Any errors generated while accessing the user information.
//Process: Get user information and return the home directory
func GetHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	homedir := usr.HomeDir
	return homedir
}
