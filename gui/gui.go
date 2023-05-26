package gui

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/lemorz56/pcapreplay/commons"
	"github.com/lemorz56/pcapreplay/components"
	"github.com/lemorz56/pcapreplay/nic"
)

// TODO: This should not be in a GUI package and be global?

func Create() {
	// Main Window
	mainWindow := &Window{
		App: app.NewWithID("lemorz56.pcapreplay"),
	}
	mainWindow.Win = mainWindow.App.NewWindow("PCAP Replay")
	mainWindow.Win.CenterOnScreen()
	mainWindow.Win.Resize(fyne.NewSize(800, 800))

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
	// binding
	//replayPane.Stats1.Bind(binding.BindString(&commons.Stats1.Text))
	//replayPane.Stats2.Bind(binding.BindString(&commons.Stats2.Text))
	//replayPane.StatPBar.Bind(binding.BindFloat(&commons.StatPBar.Value))
	mainWindow.ReplayPane = replayPane

	// CONTROLS container
	controlPane := components.NewControlsPane()
	controlPane.InitPane()
	controlPane.CreateAndFillContainer()
	mainWindow.ControlsPane = controlPane

	// TESTING BINDS
	// todo: solve this needing to be commons
	//commons.Stats1 = widget.NewEntry()
	//commons.Stats2 = widget.NewEntry()
	//commons.StatPBar = widget.NewProgressBar()

	// TEST
	//statsPbarBinding := binding.BindFloat(&commons.StatPBar)
	//stats1Binding := binding.BindString(&commons.Stats1)
	//stats2Binding := binding.BindString(&commons.Stats2)

	//replayData := &replayData{
	//	StatPBar:        replayPane.StatPBar,
	//	StatPBarBinding: statsPbarBinding,
	//	Stats1:          replayPane.Stats1,
	//	Stats1Binding:   stats1Binding,
	//	Stats2:          replayPane.Stats2,
	//	Stats2Binding:   stats2Binding,
	//}

	str := binding.NewString()
	_ = str.Set(fmt.Sprintf("Testing at %v", time.Now()))
	text := widget.NewLabelWithData(str)

	// root container that contains all other containers/panes
	motherContainer := container.New(layout.NewVBoxLayout(), text, interfacesPane.Container, replayPane.Container, controlPane.Container)
	// todo: fix so that we use a method to do this
	mainWindow.RootContainer = motherContainer

	mainWindow.Win.SetContent(motherContainer)

	go func() {
		time.Sleep(5 * time.Second)
		_ = str.Set(fmt.Sprintf("Done at %v", time.Now()))
	}()
	mainWindow.Win.ShowAndRun()
}

// type replayData struct {
// 	StatPBar        *widget.ProgressBar
// 	StatPBarBinding binding.Float
// 	Stats1          *widget.Entry
// 	Stats1Binding   binding.String
// 	Stats2          *widget.Entry
// 	Stats2Binding   binding.String
// }
