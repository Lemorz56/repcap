package components

import (
	"fmt"
	"log"
	"net"
	"runtime"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/google/gopacket/pcap"
	"github.com/lemorz56/pcapreplay/commons"
)

type InterfacesPane struct {
	Container *fyne.Container
	Label     *widget.Label  //todo: make private
	ComboBox  *widget.Select //todo: make private
	//OnSelect  func(s string)
	// Windows only to get NIC
	DeviceName        *widget.Label
	DeviceDescription *widget.Label
	DeviceAddress     *widget.Label
}

func NewInterfacesPane() *InterfacesPane {
	return &InterfacesPane{}
}

func (ip *InterfacesPane) InitPane() { //onSelectFunc func(s string)

	ip.Label = widget.NewLabelWithStyle("Net Interfaces",
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)

	ip.DeviceName = widget.NewLabelWithStyle("Name",
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)
	ip.DeviceName.Bind(binding.BindString(&commons.DeviceName))

	ip.DeviceDescription = widget.NewLabelWithStyle("Description",
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)
	ip.DeviceDescription.Bind(binding.BindString(&commons.DeviceDescription))

	ip.DeviceAddress = widget.NewLabelWithStyle("Address",
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)
	ip.DeviceAddress.Bind(binding.BindString(&commons.DeviceAddress))

	//ip.ComboBox = widget.NewSelect([]string{}, onSelectFunc)
	ip.ComboBox = widget.NewSelect([]string{}, func(s string) {
		onInterfaceSelect(s, ip)
	})

	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		intfs, _ := net.Interfaces()
		for _, intf := range intfs {
			ip.ComboBox.Options = append(ip.ComboBox.Options, intf.Name)
		}
	} else if runtime.GOOS == "windows" {
		var err error
		devices, err := pcap.FindAllDevs()
		if err != nil {
			log.Fatal(err)
		}

		for _, device := range devices {
			fmt.Println("\nName: ", device.Name)
			ip.ComboBox.Options = append(ip.ComboBox.Options, device.Name)

			fmt.Println("Description: ", device.Description)
			for _, address := range device.Addresses {
				fmt.Println("- IP address: ", address.IP)
			}
		}
	}
}

func (ip *InterfacesPane) CreateAndFillContainer() {
	ip.Container = container.New(
		layout.NewVBoxLayout(),
		ip.Label,
		ip.ComboBox,
		ip.DeviceName,
		ip.DeviceDescription,
		ip.DeviceAddress,
	)
}

func onInterfaceSelect(s string, ip *InterfacesPane) {
	fmt.Println("Chose interface:", s)
	commons.IntfId = s
	fmt.Println("commons.IntfId:", commons.IntfId)

	// todo: create a struct/funcs to handle the interfaces
	// since it differs so much between windows and linux
	if runtime.GOOS == "windows" {
		devices, err := pcap.FindAllDevs()
		if err != nil {
			log.Fatal(err)
		}

		for _, device := range devices {
			if device.Name == s {
				commons.DeviceName = device.Name
				commons.DeviceDescription = device.Description

				stringSlice := []string{}
				for _, address := range device.Addresses {
					stringSlice = append(stringSlice, address.IP.String())
				}
				commons.DeviceAddress = strings.Join(stringSlice, ", ")

				//todo fix:
				ip.DeviceName.SetText(fmt.Sprintf("Name: %s", device.Name))
				ip.DeviceDescription.SetText(fmt.Sprintf("Description: %s", device.Description))
				ip.DeviceAddress.SetText(fmt.Sprintf("Address(es): %s", strings.Join(stringSlice, ", ")))
				break
			}
		}
	}
}
