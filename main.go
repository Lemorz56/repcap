package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lemorz56/pcapreplay/cli"
)

var (
	version = "v0.0.0"
	commit  = "local"
	date    = "01-01-01"
)

func main() {
	app := cli.CreateCliApp(fmt.Sprintf("%s (%s) %s", version, commit, date))
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
