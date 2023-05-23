package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"net"
	"os"
	"runtime"

	gpcap "github.com/google/gopacket/pcap"
	"github.com/lemorz56/pcapreplay/commons"
	"github.com/lemorz56/pcapreplay/pcap"
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

	app.Commands = []*cli.Command{{
		Name: "list",
		Aliases: []string{
			"ls",
		},
		Description: "list all interfaces",
		Action:      ListAllInterfaces,
	},
	}

	app.Action = func(c *cli.Context) error {
		if commons.WithGui {
			CreateGui()
		} else {
			pcap.Replay()
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// todo: https://github.com/esnet/gdg/tree/master/.github

func ListAllInterfaces(c *cli.Context) error {
	if runtime.GOOS == "windows" {
		interfaces, err := gpcap.FindAllDevs()
		if err != nil {
			fmt.Println("Could not list all interfaces: ", err)
			return err
		}

		for _, i := range interfaces {
			fmt.Println("Name:", i.Name)
			fmt.Println("Desc:", i.Description)
			fmt.Println("Addrs:", i.Addresses)
		}
		return nil
	} else {
		interfaces, err := net.Interfaces()
		if err != nil {
			fmt.Println("Could not list all interfaces: ", err)
			return err
		}

		for _, i := range interfaces {
			fmt.Println("Name:", i.Name)
		}
		return nil
	}
}
