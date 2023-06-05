package gui

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/lemorz56/repcap/commons"
	"github.com/lemorz56/repcap/components"
	"github.com/lemorz56/repcap/nic"
)

func Create() {
	// Main Window
	mainWindow := &Window{
		App: app.NewWithID("lemorz56.repcap"),
	}
	mainWindow.Win = mainWindow.App.NewWindow("PCAP Replay")
	mainWindow.Win.CenterOnScreen()
	mainWindow.Win.Resize(fyne.NewSize(800, 800))

	ctx := context.Background()

	// INTERFACES container
	nicService := nic.NewNicService()

	interfacesPane := components.NewInterfacesPane(nicService)
	interfacesPane.InitPane()
	interfacesPane.CreateAndFillContainer()
	if commons.IntfId != "" {
		interfacesPane.ComboBox.SetSelected(commons.IntfId)
	}
	mainWindow.InterfacesPane = interfacesPane

	// REPLAY container
	replayPane := components.NewReplayPane()
	replayPane.MainWindow = &mainWindow.Win
	replayPane.InitPane()
	replayPane.CreateAndFillContainer()

	// CONTROLS container
	controlPane := components.NewControlsPane()
	controlPane.InitPane(ctx)
	controlPane.CreateAndFillContainer()

	mainWindow.InterfacesPane = interfacesPane
	mainWindow.ReplayPane = replayPane
	mainWindow.ControlsPane = controlPane

	// root container that contains all other containers/panes
	// todo: fix so that we use a method to do this, for now since the positions might be changing we can just do it like this
	motherContainer := container.New(layout.NewVBoxLayout(), interfacesPane.Container, replayPane.Container, layout.NewSpacer(), controlPane.Container)
	mainWindow.RootContainer = motherContainer
	mainWindow.Win.SetContent(motherContainer)
	mainWindow.Win.ShowAndRun()

	//todo: implement cache using fyne.storage
	//uri, err := storage.
}
