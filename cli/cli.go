package cli

import (
	"fmt"

	"github.com/lemorz56/repcap/commons"
	"github.com/lemorz56/repcap/gui"
	"github.com/lemorz56/repcap/nic"
	"github.com/lemorz56/repcap/pcap"
	"github.com/urfave/cli/v2"
)

func CreateCliApp(version string) *cli.App {
	app := cli.NewApp()
	app.Name = "repcap"
	app.Version = version
	app.Usage = "repcap"
	app.UsageText = "repcap --interface <interface> [--fast]  --pcap <pcap file>"
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Printf("Command not found: %v\n", command)
		fmt.Println("Use --help or -h to see available commands")
	}

	app.Commands = []*cli.Command{
		{
			Name: "replay",
			Aliases: []string{
				"rp",
			},
			Description: "Replay a pcap file",
			Action: func(c *cli.Context) error {
				pcap.Replay()
				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "interface",
					Usage:       "Network Interface to use",
					Destination: &commons.IntfId,
					Required:    true,
					Aliases:     []string{"i"},
				},
				&cli.BoolFlag{
					Name:        "fast",
					Usage:       "Replay fast without the recorded delay between each packet",
					Value:       true,
					Destination: &commons.ReplayFast,
					Aliases:     []string{"f"},
				},
				&cli.StringFlag{
					Name:        "pcap",
					Usage:       "Pcap file to replay",
					Required:    true,
					Destination: &commons.PcapFile,
					Aliases:     []string{"p"},
				},
			},
		},
		{
			Name: "list",
			Aliases: []string{
				"ls",
			},
			Description: "list all interfaces",
			Action:      ListAllInterfaces,
		},
		{
			Name: "gui",
			Aliases: []string{
				"g",
			},
			Description: "Start the gui",
			Action: func(c *cli.Context) error {
				commons.WithGui = true
				gui.Create()
				return nil
			},
		},
	}

	return app
}

func ListAllInterfaces(c *cli.Context) error {
	nicSvc := nic.NewNicService()
	err := nicSvc.InitNics()
	if err != nil {
		fmt.Println("Could not initialize NicService: ", err)
		return err
	}
	nics, err := nicSvc.GetAllNics()
	if err != nil {
		fmt.Println("Could not list all interfaces: ", err)
		return err
	}

	for _, i := range nics {
		fmt.Println("Name:", i.Id)
		fmt.Println("Description:", i.Description)
		fmt.Println("Addresses:", i.Addresses)
		fmt.Println("")
	}
	return nil
}
