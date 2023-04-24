// https://github.com/fyne-io/fyne#getting-started
package main

import (
	"os"

	"github.com/lemorz56/pcapreplay/commons"
	"github.com/lemorz56/pcapreplay/pcap"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "PCAP Replay"
	app.Version = "0.0.1"
	app.Usage = "pcapreplay"
	app.UsageText = "pcapreplay --interface <interface> [--realtime] [--gui] --pcap <pcap file>"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "interface",
			Usage:       "system interface id",
			Destination: &commons.IntfId,
			Aliases:     []string{"i"},
		},
		&cli.BoolFlag{
			Name:        "realtime",
			Usage:       "replay without the real time between each packets",
			Value:       false,
			Destination: &commons.ReplayFast,
			Aliases:     []string{"r"},
		},
		&cli.BoolFlag{
			Name:        "gui",
			Usage:       "start the helper gui",
			Required:    false,
			Value:       false,
			Destination: &commons.WithGui,
		},
		&cli.StringFlag{
			Name:        "pcap",
			Usage:       "pcap file to replay",
			Required:    true,
			Destination: &commons.PcapFile,
		},
	}

	app.Action = func(c *cli.Context) error {
		if commons.WithGui {
			// todo: gui
		} else {
			pcap.Replay()
		}

		return nil
	}

	app.Run(os.Args)
}
