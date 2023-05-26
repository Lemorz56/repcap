package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/lemorz56/pcapreplay/commons"
)

type ReplayPane struct {
	// Main
	Container *fyne.Container
	Label     *widget.Label
	// todo: sneaky reference to the main window maybe?
	MainWindow *fyne.Window
	// Stats
	//Stats1   *widget.Entry
	//Stats2   *widget.Entry
	Stats1   *widget.Label
	Stats2   *widget.Label
	StatPBar *widget.ProgressBar
	// File
	FileField  *widget.Entry
	FileDialog *dialog.FileDialog
	FileButton *widget.Button
}

func NewReplayPane() *ReplayPane {
	return &ReplayPane{}
}

func (rp *ReplayPane) InitPane() { //todo: pass window to initPane or NewReplayPane
	rp.Label = widget.NewLabelWithStyle("Replay",
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)
	//rp.Stats1 = widget.NewEntry()
	//rp.Stats2 = widget.NewEntry()
	rp.Stats1 = widget.NewLabelWithData(commons.Stats1)
	rp.Stats2 = widget.NewLabelWithData(commons.Stats2)
	rp.StatPBar = widget.NewProgressBarWithData(commons.StatPBar)
	//rp.disableControls()

	rp.FileField = widget.NewEntry()
	if commons.PcapFile != "" {
		uri := storage.NewFileURI(commons.PcapFile)
		rp.FileField.SetText(uri.Name())

	}
	rp.FileDialog = dialog.NewFileOpen(rp.onFileDialogClosed, *rp.MainWindow)
	rp.FileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".pcap"}))
	rp.FileDialog.Resize(fyne.NewSize(800, 800))

	rp.FileButton = widget.NewButtonWithIcon("", theme.FolderOpenIcon(), rp.onFileButtonClicked)
}

func (rp *ReplayPane) CreateAndFillContainer() {
	rp.Container = container.New(
		layout.NewVBoxLayout(),
		container.New(layout.NewGridLayout(2), rp.FileField, rp.FileButton),
		rp.Stats1,
		rp.Stats2,
		rp.StatPBar)
}

// func (rp *ReplayPane) disableControls() {
// 	//rp.Stats1.Disable()
// 	//rp.Stats2.Disable()
// }

func (rp *ReplayPane) onFileDialogClosed(uc fyne.URIReadCloser, err error) {
	if err != nil {
		diag := dialog.NewConfirm("Unexpected Error", err.Error(), nil, nil)
		diag.Show()
		return
	}

	if uc != nil {
		if uc.URI() != nil {
			if uc.URI().Path() != "" {
				commons.PcapFile = uc.URI().Path()
				rp.FileField.SetText(uc.URI().Name())
			}
		}
	} else {
		rp.FileField.SetText("No Valid File Found")
	}
}

func (rp *ReplayPane) onFileButtonClicked() {
	rp.FileDialog.Show()
}
