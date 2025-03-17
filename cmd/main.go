package main

import (
	"ais_service/internal/wiring"
	"os"
)

const (
	flagConfigFilePath = ""
)

func main() {
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")

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
