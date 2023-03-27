package local

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"publish/config"
	"runtime"
)

func Exec() {
	var result string
	for _, item := range config.LocalCmd {
		result = execute(item)
		fmt.Println(result)
	}
}

func execute(arg ...string) (result string) {
	name := "/bin/bash"
	c := "-c"
	if runtime.GOOS == "windows" {
		name = "cmd"
		c = "/C"
	}
	arg = append([]string{c}, arg...)
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Error:can not obtain stdout pipe for command:%s\n", err)
		return
	}
	if err := cmd.Start(); err != nil {
		log.Fatalln("Error:The command is err,", err)
		return

	}
	bytes, err := io.ReadAll(stdout)
	if err != nil {
		log.Fatalln("ReadAll Stdout:", err.Error())
		return
	}
	if err := cmd.Wait(); err != nil {
		log.Fatalln("wait:", err.Error())
		return
	}
	result = string(bytes)
	return
}
