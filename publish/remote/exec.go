package remote

import (
	"log"
	"publish/config"
)

func Exec() {
	for _, item := range config.RemoteBeforeCmd {
		run(item)
	}
	Upload()
	for _, item := range config.RemoteAfterCmd {
		run(item)
	}
}

func run(shell string) {
	session, err := GetSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	buf, err := session.CombinedOutput(shell)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(string(buf))
}
