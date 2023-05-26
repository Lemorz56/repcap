package components

import (
	"fmt"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"github.com/lemorz56/pcapreplay/nic"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/lemorz56/pcapreplay/commons"
)

type InterfacesPane struct {
	NicService        nic.NicServiceInterface
	Container         *fyne.Container
	Label             *widget.Label
	ComboBox          *widget.Select
	SelectedInterface *selectedInterface
}

type selectedInterface struct {
	Name        *widget.Label
	Description *widget.Label
	Address     *widget.Label
}

func NewInterfacesPane(ns nic.NicServiceInterface) *InterfacesPane {
	ip := &InterfacesPane{
		NicService: ns,
		Container:  nil,
		Label:      widget.NewLabelWithStyle("Net Interfaces", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		ComboBox:   nil,
		SelectedInterface: &selectedInterface{
			Name:        widget.NewLabel("Name"),
			Description: widget.NewLabel("Description"),
			Address:     widget.NewLabel("Address"),
		},
	}

	ip.SelectedInterface.Address.Bind(binding.BindString(&commons.DeviceName))
	ip.SelectedInterface.Description.Bind(binding.BindString(&commons.DeviceDescription))
	ip.SelectedInterface.Address.Bind(binding.BindString(&commons.DeviceAddress))

	ip.initCombobox()

	return ip
}

func (ip *InterfacesPane) InitPane() {
	err := ip.NicService.InitNics()
	if err != nil {
		log.Fatal(err) //todo: make popup?
	}

	devices, err := ip.NicService.GetAllNics()
	if err != nil {
		log.Fatal(err) //todo: make popup?
	}

	for _, device := range devices {
		ip.ComboBox.Options = append(ip.ComboBox.Options, device.Id)
		//todo: add debug logs
		fmt.Println("appended device.Id:", device.Id)
		if commons.IntfId != "" {
			if device.Id == commons.IntfId {
				ip.ComboBox.SetSelected(device.Id)
			}
		}
	}
}

func (ip *InterfacesPane) CreateAndFillContainer() {
	ip.Container = container.New(
		layout.NewVBoxLayout(),
		ip.Label,
		ip.ComboBox,
		ip.SelectedInterface.Name,
		ip.SelectedInterface.Description,
		ip.SelectedInterface.Address,
	)
}

func (ip *InterfacesPane) initCombobox() {
	ip.ComboBox = widget.NewSelect([]string{}, func(s string) {
		onInterfaceSelect(s, ip)
	})
}

func onInterfaceSelect(s string, ip *InterfacesPane) {
	commons.IntfId = s
	device, err := ip.NicService.GetNicByName(s)
	if err != nil {
		log.Fatal(err)
	}

	var stringSlice []string
	stringSlice = append(stringSlice, device.Addresses...)

	commons.DeviceName = fmt.Sprintf("Name: %s", device.Id)
	commons.DeviceDescription = fmt.Sprintf("Description: %s", device.Description)
	commons.DeviceAddress = fmt.Sprintf("Address(es): %s", strings.Join(stringSlice, ", "))

	ip.SelectedInterface.Name.SetText(commons.DeviceName)
	ip.SelectedInterface.Description.SetText(commons.DeviceDescription)
	ip.SelectedInterface.Address.SetText(commons.DeviceAddress)
	//ip.SelectedInterface.Name.Refresh()
	//ip.SelectedInterface.Description.Refresh()
	//ip.SelectedInterface.Address.Refresh()
}
