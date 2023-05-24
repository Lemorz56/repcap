package commons

import (
	"fyne.io/fyne/v2/data/binding"
	"time"

	"fyne.io/fyne/v2"
	"github.com/google/gopacket/pcap"
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

	PcapHandle *pcap.Handle
)

// gui vars
// todo: make structs for vars
var (
	MainWin  fyne.Window
	MainPane fyne.Container

	DeviceName        string //todo: remove
	DeviceDescription string //todo: remove
	DeviceAddress     string //todo: remove

	// CONTAINERS
	//InterfacesPane *fyne.Container
	//ReplayPane     *fyne.Container
	//ControlsPane *fyne.Container

	// STATS

	//Stats1 *widget.Entry //todo: need to create
	//Stats2 *widget.Entry
	Stats1 = binding.NewString() //todo: need to create
	Stats2 = binding.NewString()

	// FileField *widget.Entry

	// NUMBER BOX
	// StepSpinBox *extensions.NumericalEntry

	// PROGRESSBAR

	//StatPBar *widget.ProgressBar
	StatPBar = binding.NewFloat()

	// INTERFACES
	// InterfacesLabel *widget.Label
	// Interfaces      *widget.Select
	// InterfacesList  *[]string

	// CONTROLS
	// ControlLabel   *widget.Label
	// PlayBtn        *widget.Button
	// FastPlayBtn    *widget.Button
	// StepPlayBtn    *widget.Button
	// StepOnePlayBtn *widget.Button
	// ResetBtn       *widget.Button
)
