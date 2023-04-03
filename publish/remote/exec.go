package remote

import (
	"fmt"
	"log"
	"publish/config"
	"strconv"
)

func Exec() {
	for _, item := range config.RemoteBeforeCmd {
		if item[:9] == "port-kill" {
			port, err := strconv.Atoi(item[10:])
			if err != nil {
				log.Fatalln("kill the port is illegal")
			}
			item = fmt.Sprintf("PID=$(lsof -i :%d | grep -m 1 \"LISTEN\" | awk '{print $2}')\nif [ -z \"$PID\" ]; then\n\techo \"PID is empty\"\nelse\n\tkill $PID\nfi", port)
		}
		if item[:2] == "rm" {
			fp := item[3:]
			item = fmt.Sprintf("if [ -f %s ]; then\n\trm %s\nfi", fp, fp)
		}
		run(item)
	}
	Upload()
	for _, item := range config.RemoteAfterCmd {
		if item[:9] == "port-kill" {
			port, err := strconv.Atoi(item[10:])
			if err != nil {
				log.Fatalln("kill the port is illegal")
			}
			item = fmt.Sprintf("PID=$(lsof -i :%d | grep -m 1 \"LISTEN\" | awk '{print $2}')\nif [ -z \"$PID\" ]; then\n\techo \"PID is empty\"\nelse\n\tkill $PID\nfi", port)
		}
		if item[:2] == "rm" {
			fp := item[3:]
			item = fmt.Sprintf("if [ -f %s ]; then\n\trm %s\nfi", fp, fp)
		}
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
