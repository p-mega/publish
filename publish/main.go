package main

import (
	"publish/config"
	"publish/local"
	"publish/remote"
)

func main() {
	config.Init()
	remote.Init()
	local.Exec()
	remote.Exec()
}
