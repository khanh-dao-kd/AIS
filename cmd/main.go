package main

import (
	"ais_service/internal/wiring"
)

const (
	flagConfigFilePath = ""
)

func main() {
	server, cleanup, err := wiring.InitializeServer(flagConfigFilePath)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	err = server.Start()
	if err != nil {
		panic(err)
	}
}
