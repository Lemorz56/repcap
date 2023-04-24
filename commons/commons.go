package commons

import (
	"time"

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

	PcapHndl *pcap.Handle
)
