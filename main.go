package main

import (
	"github.com/lemorz56/pcapreplay/cli"
	"log"
	"os"
)

var version = "0.0.1"

func main() {
	app := cli.CreateCliApp(version)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
