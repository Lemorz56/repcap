package commons

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/google/gopacket/pcap"
	"github.com/lemorz56/pcapreplay/extensions"
)

// options vars
var (
	IntfId   = ""
	PcapFile = ""

	ReplayFast = false
	WithGui    = false
)

// replay vars
var (
	Start     time.Time
	BytesSent int

	LastTS   time.Time
	LastSend time.Time

	Pkt     int
	TsStart time.Time
	TsEnd   time.Time
	Packets int
	Size    int

	PcapHndl *pcap.Handle
)

// gui vars
// todo: make structs for vars
var (
	MainWin  fyne.Window
	MainPane fyne.Container

	// CONTAINERS
	InterfacesPane *fyne.Container
	ReplayPane     *fyne.Container
	ControlsPane   *fyne.Container

	// STATS
	Stats1    *widget.Entry
	Stats2    *widget.Entry
	FileField *widget.Entry

	// NUMBER BOX
	StepSpinBox *extensions.NumericalEntry

	// PROGRESSBAR
	StatPBar *widget.ProgressBar

	// INTERFACES
	InterfacesLabel *widget.Label
	Interfaces      *widget.Select
	// InterfacesList  *[]string

	// CONTROLS
	ControlLabel   *widget.Label
	PlayBtn        *widget.Button
	FastPlayBtn    *widget.Button
	StepPlayBtn    *widget.Button
	StepOnePlayBtn *widget.Button
	ResetBtn       *widget.Button
)
