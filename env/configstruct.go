package env

//Config is a struct that contains the configuration information.
type Config struct {
	Client            string //The registered name of the client on the AniDB API.
	Clientver         int    //The registered version of the client on the AniDB API.
	Protover          int    //The version of the AniDB API to use.
	Url               string //The URL for the AniDB API.
	Port              int    //Port number used to connect to the AniDB API.
	ProgramDataPath   string //Path where Anime and Episodes will be stored
	ProgramConfigPath string //Path for all other data
}
