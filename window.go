package main

import (
	"fyne.io/fyne/v2"
	"github.com/lemorz56/pcapreplay/components"
)

type Windows struct {
	// Main
	App fyne.App
	Win fyne.Window
	// RootContainer
	RootContainer *fyne.Container

	// Main Panes
	InterfacesPane *components.InterfacesPane
	ReplayPane     *components.ReplayPane
	ControlsPane   *components.ControlsPane
}

func CreateNewWindow() {

}
