package main

import (
	"fmt"
	"net"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/lemorz56/pcapreplay/commons"
)

func DisableControls() {
	// Entries
	commons.Stats1.Disable()
	commons.Stats2.Disable()
	commons.FileField.Disable()

	commons.StepSpinBox.Disable()
	commons.Interfaces.Disable()

	// Buttons
	commons.PlayBtn.Disable()
	commons.FastPlayBtn.Disable()
	commons.StepPlayBtn.Disable()
	commons.StepOnePlayBtn.Disable()
	commons.ResetBtn.Disable()
}

func EnableControls() {
	// Entries
	commons.Stats1.Enable()
	commons.Stats2.Enable()
	commons.FileField.Enable()

	commons.StepSpinBox.Enable()
	commons.Interfaces.Enable()

	// Buttons
	commons.PlayBtn.Enable()
	commons.FastPlayBtn.Enable()
	commons.StepPlayBtn.Enable()
	commons.StepOnePlayBtn.Enable()
	commons.ResetBtn.Enable()
}

func CreateGui() {
	// todo: put app in main
	a := app.New()

	// Main Window
	commons.MainWin = a.NewWindow("PCAP Replay - v0")
	commons.MainWin.CenterOnScreen()
	commons.MainWin.Resize(fyne.NewSize(800, 800))

	// INTERFACES container
	commons.InterfacesLabel = widget.NewLabelWithStyle("Net Interfaces", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	commons.Interfaces = widget.NewSelect([]string{}, onInterfaceSelect)

	intfs, _ := net.Interfaces()
	for _, intf := range intfs {
		commons.Interfaces.Options = append(commons.Interfaces.Options, intf.Name)
	}
	commons.InterfacesPane = container.New(layout.NewVBoxLayout(), commons.InterfacesLabel, commons.Interfaces)

	// REPLAY container
	// ctor stuff
	commons.Stats1 = widget.NewEntry()
	commons.Stats2 = widget.NewEntry()
	commons.FileField = widget.NewEntry()

	//todo: add fileDialog button next to field

	fileDialog := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.NewConfirm("Error", err.Error(), nil, commons.MainWin)
		}

		commons.PcapFile = uc.URI().Path()
		commons.FileField.Text = uc.URI().Name()
	}, commons.MainWin)

	commons.ReplayPane = container.New(layout.NewVBoxLayout())

	// CONTROLS container
	commons.ControlLabel = widget.NewLabelWithStyle("Controls", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	commons.PlayBtn = widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() { fmt.Println("play") })
	commons.FastPlayBtn = widget.NewButtonWithIcon("", theme.MediaFastForwardIcon(), func() { fmt.Println("fast foward") })
	commons.ResetBtn = widget.NewButtonWithIcon("", theme.MediaStopIcon(), func() { fmt.Println("reset") })
	commons.StepOnePlayBtn = widget.NewButtonWithIcon("StepOne", theme.MediaSkipNextIcon(), func() { fmt.Println("fast rewind") })
	// -> containers
	controlButtons := container.New(layout.NewHBoxLayout(), commons.PlayBtn, commons.FastPlayBtn, commons.ResetBtn, commons.StepOnePlayBtn)
	commons.ControlsPane = container.New(layout.NewVBoxLayout(), commons.ControlLabel, controlButtons)

	//
	motherContainer := container.New(layout.NewVBoxLayout(), commons.InterfacesPane, commons.ControlsPane)
	commons.MainWin.SetContent(motherContainer)
	commons.MainWin.ShowAndRun()
}

func onInterfaceSelect(s string) {
	interfaces, _ := net.Interfaces()

	commons.IntfId = interfaces[commons.Interfaces.SelectedIndex()].Name
}
