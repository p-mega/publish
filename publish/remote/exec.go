package remote

import (
	"log"
	"publish/config"
)

func Exec() {
	session, err := GetSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	for _, item := range config.RemoteBeforeCmd {
		buf, err := session.CombinedOutput(item)
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(string(buf))
		session.Stdout = nil
		session.Stderr = nil
	}
	Upload()
	for _, item := range config.RemoteAfterCmd {
		buf, err := session.CombinedOutput(item)
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(string(buf))
		session.Stdout = nil
		session.Stderr = nil
	}
}
