package config

import (
	"log"

	"github.com/spf13/viper"
)

var PrivateKey string

func Init() {
	config := viper.New()
	config.AddConfigPath("./")
	config.SetConfigName("publish")
	config.SetConfigType("yaml")
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln("not found config file: publish.yaml")
		} else {
			log.Fatalln("this configuration file has syntax errors")
		}
	}
	Server.Host = config.GetString("server.host")
	Server.Port = config.GetInt("server.port")
	Server.User = config.GetString("server.user")
	Server.Password = config.GetString("server.password")
	FileRoute.Local = config.GetString("file.local")
	FileRoute.Remote = config.GetString("file.remote")
	LocalCmd = config.GetStringSlice("localCmd")
	RemoteBeforeCmd = config.GetStringSlice("RemoteBeforeCmd")
	RemoteAfterCmd = config.GetStringSlice("RemoteAfterCmd")
}
