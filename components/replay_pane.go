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
	// Stats
	Stats1   *widget.Entry
	Stats2   *widget.Entry
	StatPBar *widget.ProgressBar
	// File
	FileField  *widget.Entry
	FileDialog *dialog.FileDialog
	FileButton *widget.Button
}

func NewReplayPane() *ReplayPane {
	return &ReplayPane{}
}

func (rp *ReplayPane) InitPane(mainWindow fyne.Window) {
	rp.Label = widget.NewLabelWithStyle("Replay",
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)

	rp.Stats1 = widget.NewEntry()
	rp.Stats2 = widget.NewEntry()
	rp.StatPBar = widget.NewProgressBar()
	rp.disableControls()

	rp.FileField = widget.NewEntry()
	rp.FileDialog = dialog.NewFileOpen(rp.onFileDialogClosed, mainWindow) //todo: maybe add window?
	rp.FileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".pcap"}))
	rp.FileDialog.Resize(fyne.NewSize(800, 800))

	rp.FileButton = widget.NewButtonWithIcon("", theme.FolderOpenIcon(), rp.onFileButtonClicked)
}

func (rp *ReplayPane) CreateAndFillContainer() {
	rp.Container = container.New(
		layout.NewVBoxLayout(),
		container.New(layout.NewGridLayout(2), rp.FileField, rp.FileButton),
		rp.Stats1,
		rp.Stats2)
}

func (rp *ReplayPane) disableControls() {
	rp.Stats1.Disable()
	rp.Stats2.Disable()
}

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
