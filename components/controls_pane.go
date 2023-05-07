package components

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ControlsPane struct {
	Container   *fyne.Container
	Label       *widget.Label //todo: make private
	PlayBtn     *widget.Button
	FastPlayBtn *widget.Button
	//StepPlayBtn    *widget.Button
	StepOnePlayBtn *widget.Button
	ResetBtn       *widget.Button
	//StepSpinBox    *extensions.NumericalEntry
}

func NewControlsPane() *ControlsPane {
	return &ControlsPane{}
}

func (cp *ControlsPane) InitPane() {

	cp.Label = widget.NewLabelWithStyle("Controls",
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)
	cp.initButtons()

	//cp.StepSpinBox = extensions.NewNumericalEntry()
}

func (cp *ControlsPane) CreateAndFillContainer() {
	cp.Container = container.New(
		layout.NewVBoxLayout(),
		cp.Label,
		container.New(
			layout.NewHBoxLayout(),
			cp.PlayBtn,
			cp.FastPlayBtn,
			cp.StepOnePlayBtn,
			cp.ResetBtn),
	)
}

// initialize all buttons with icons and callbacks
func (cp *ControlsPane) initButtons() {
	cp.PlayBtn = widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() { fmt.Println("play") })
	cp.StepOnePlayBtn = widget.NewButtonWithIcon("StepOne", theme.MediaSkipNextIcon(), func() { fmt.Println("fast rewind") })
	//cp.StepPlayBtn = widget.NewButtonWithIcon("", theme.Media, func() { fmt.Println("step") })
	cp.FastPlayBtn = widget.NewButtonWithIcon("", theme.MediaFastForwardIcon(), func() { fmt.Println("fast foward") })
	cp.ResetBtn = widget.NewButtonWithIcon("", theme.MediaStopIcon(), func() { fmt.Println("reset") })
}
