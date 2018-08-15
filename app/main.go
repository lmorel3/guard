package main

import (
	"github.com/lmorel3/guard-go/app/config"
	"github.com/lmorel3/guard-go/app/server"
)

func main() {
	config.Init()
	server.Init()
}
