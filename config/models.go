package config

var Server struct {
	Host     string
	Port     int
	User     string
	Password string
}

var FileRoute struct {
	Local  string
	Remote string
}

var LocalCmd []string

var RemoteBeforeCmd []string

var RemoteAfterCmd []string
