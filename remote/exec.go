package remote

import (
	"bytes"
	"fmt"
	"log"
	"publish/config"
)

func Exec() {
	var stdOut, stdErr bytes.Buffer
	session, err := getSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	session.Stdout = &stdOut
	session.Stderr = &stdErr
	for _, item := range config.RemoteBeforeCmd {
		session.Run(item)
		fmt.Println(stdOut.String())
		fmt.Println(stdErr.String())
	}
	upload()
	for _, item := range config.RemoteAfterCmd {
		session.Run(item)
		fmt.Println(stdOut.String())
		fmt.Println(stdErr.String())
	}
}
