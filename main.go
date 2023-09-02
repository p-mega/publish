package main

import (
	"github.com/p-mega/publish/config"
	"github.com/p-mega/publish/local"
	"github.com/p-mega/publish/remote"
)

func main() {
	config.Init()
	remote.Init()
	local.Exec()
	remote.Exec()
}
