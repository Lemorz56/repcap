package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/lemorz56/pcapreplay/commons"
	"github.com/lemorz56/pcapreplay/components"
)

func CreateGui() {
	// todo: put app in main
	a := app.New()

	// Main Window
	commons.MainWin = a.NewWindow("PCAP Replay - v0")
	commons.MainWin.CenterOnScreen()
	commons.MainWin.Resize(fyne.NewSize(800, 800))

	// INTERFACES container
	// commons.InterfacesLabel = widget.NewLabelWithStyle("Net Interfaces", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	// commons.Interfaces = widget.NewSelect([]string{}, onInterfaceSelect)

	interfacesPane := components.NewInterfacesPane()
	interfacesPane.InitPane()
	interfacesPane.CreateAndFillContainer()

	if commons.IntfId != "" {
		interfacesPane.ComboBox.SetSelected(commons.IntfId)
	}

	// REPLAY container
	// // ctor stuff
	commons.Stats1 = widget.NewEntry()
	commons.Stats2 = widget.NewEntry()
	commons.StatPBar = widget.NewProgressBar()
	// commons.FileField = widget.NewEntry()

	// // IF pcap arg is set at launch
	// commons.FileField.SetText(commons.PcapFile)

	// currentSize := commons.FileField.Size()
	// commons.FileField.Resize(fyne.NewSize(10, currentSize.Height))

	// // FILE Dialog
	// fileDialog := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
	// 	if err != nil {
	// 		dialog.NewConfirm("Error", err.Error(), nil, commons.MainWin)
	// 	}

	// 	commons.PcapFile = uc.URI().Path() //todo: data binding, the UI isnt updating
	// 	//commons.FileField.Text = uc.URI().Name()
	// 	commons.FileField.SetText(uc.URI().Name())
	// 	fmt.Println("path:", uc.URI().Path())
	// 	fmt.Println("name:", uc.URI().Name())
	// }, commons.MainWin)
	// fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".pcap"}))
	// fileDialog.Resize(fyne.NewSize(600, 600))

	// fileButton := widget.NewButtonWithIcon("", theme.FileIcon(), fileDialog.Show)

	// fileRow := container.New(layout.NewHBoxLayout(), commons.FileField, fileButton)
	// fileRow := container.NewWithoutLayout(commons.FileField, fileButton)
	//fileRow := container.New(layout.NewGridLayout(2), commons.FileField, fileButton)

	replayPane := components.NewReplayPane()
	replayPane.InitPane(commons.MainWin)
	replayPane.CreateAndFillContainer()

	replayPane.Stats1.Bind(binding.BindString(&commons.Stats1.Text))
	replayPane.Stats2.Bind(binding.BindString(&commons.Stats2.Text))
	replayPane.StatPBar.Bind(binding.BindFloat(&commons.StatPBar.Value))

	controlPane := components.NewControlsPane()
	controlPane.InitPane()
	controlPane.CreateAndFillContainer()

	//commons.ReplayPane = container.New(layout.NewVBoxLayout(), fileRow, commons.Stats1, commons.Stats2)

	// // CONTROLS container
	// commons.ControlLabel = widget.NewLabelWithStyle("Controls", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	// commons.PlayBtn = widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() { fmt.Println("play") })
	// commons.FastPlayBtn = widget.NewButtonWithIcon("", theme.MediaFastForwardIcon(), func() { fmt.Println("fast foward") })
	// commons.ResetBtn = widget.NewButtonWithIcon("", theme.MediaStopIcon(), func() { fmt.Println("reset") })
	// commons.StepOnePlayBtn = widget.NewButtonWithIcon("StepOne", theme.MediaSkipNextIcon(), func() { fmt.Println("fast rewind") })
	// // -> containers
	// controlButtons := container.New(layout.NewHBoxLayout(), commons.PlayBtn, commons.FastPlayBtn, commons.ResetBtn, commons.StepOnePlayBtn)
	// commons.ControlsPane = container.New(layout.NewVBoxLayout(), commons.ControlLabel, controlButtons)

	// root container that contains all other containers/panes
	motherContainer := container.New(layout.NewVBoxLayout(), interfacesPane.Container, replayPane.Container, controlPane.Container)
	commons.MainWin.SetContent(motherContainer)
	commons.MainWin.ShowAndRun()
}

func onInterfaceSelect(s string) {
	commons.IntfId = s
}
