package pcap

import (
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/lemorz56/repcap/commons"
)

func WritePacketDelayed(handle *pcap.Handle, buf []byte, ci gopacket.CaptureInfo) {
	if ci.CaptureLength != ci.Length {
		return
	}

	intervalInCapture := ci.Timestamp.Sub(commons.LastTS)
	elapsedTime := time.Since(commons.LastSend)

	if (intervalInCapture > elapsedTime) && !commons.LastSend.IsZero() {
		time.Sleep(intervalInCapture - elapsedTime)
	}

	commons.LastSend = time.Now()
	//todo: add counter to count the number of packets NOT sent
	// if WritePacket returns an error, increment the counter?
	_ = WritePacket(handle, buf)
	commons.LastTS = ci.Timestamp
}

func WritePacket(handle *pcap.Handle, buf []byte) error {
	if err := handle.WritePacketData(buf); err != nil {
		return errors.New("failed to send packet: " + err.Error())
	}
	return nil
}

func Infos(filename string) (start, end time.Time, packets, size int) {
	handleRead, err := pcap.OpenOffline(filename)
	if err != nil {
		log.Fatalf("PCAP OpenOffline error (handle to read packet): %v", err)
	}
	defer handleRead.Close()

	var previousTS time.Time
	var deltaTotal time.Duration

	for {
		data, ci, err := handleRead.ReadPacketData()
		if err != nil && err != io.EOF {
			log.Fatal(err)
		} else if err == io.EOF {
			break
		}

		if start.IsZero() {
			start = ci.Timestamp
		}
		end = ci.Timestamp
		packets++
		size += len(data)

		if previousTS.IsZero() {
			previousTS = ci.Timestamp
		} else {
			deltaTotal += ci.Timestamp.Sub(previousTS)
			previousTS = ci.Timestamp
		}
	}

	sec := int(deltaTotal.Seconds())
	if sec == 0 {
		sec = 1
	}

	log.Printf("Avg packet rate %d/s\n", packets/sec)
	if commons.WithGui {
		s := fmt.Sprintf("%d", packets/sec)
		err = commons.Stats1.Set(s)
		if err != nil {
			log.Printf("Error: %v", err)
		}
	}

	return start, end, packets, size
}

func LoadPcap(filename string) {
	var err error
	commons.PcapHandle, err = pcap.OpenOffline(filename)
	if err != nil {
		log.Fatalf("PCAP OpenOffline error (handle to read packet): %v", err)
	}

	commons.Start = time.Now()
	commons.Pkt = 0
	commons.TsStart, commons.TsEnd, commons.Packets, commons.Size = Infos(commons.PcapFile)
}

func OpenDest(netIntf string) *pcap.Handle {
	inactive, err := pcap.NewInactiveHandle(netIntf)
	if err != nil {
		log.Fatalf("Error creating inactive handle: %v", err)
	}
	_ = inactive.SetPromisc(true)
	defer inactive.CleanUp()

	handleWrite, _ := inactive.Activate()

	return handleWrite
}

func EndReplay() {
	if commons.PcapHandle != nil {
		commons.PcapHandle.Close()
		commons.PcapHandle = nil
	}

	if commons.WithGui {
		err := commons.StatPBar.Set(-1)
		if err != nil {
			log.Printf("Error setting progress bar: %v", err)
		}
	}
}

func InternalReplay(handleWrite *pcap.Handle) bool {
	data, ci, err := commons.PcapHandle.ReadPacketData()

	switch {
	case err == io.EOF:
		log.Printf("\nFinished in %s", time.Since(commons.Start))
		EndReplay()

		if commons.WithGui {
			s := fmt.Sprintf("Finished in %s", time.Since(commons.Start))
			err = commons.Stats2.Set(s)
			if err != nil {
				fmt.Println("Error setting stats2: ", err)
			}
			commons.PcapHandle = nil
		}
		return true

	case err != nil:
		log.Printf("Failed to read packet %d: %s\n", commons.Pkt, err)

	default:
		if commons.ReplayFast {
			_ = WritePacket(handleWrite, data)
		} else {
			WritePacketDelayed(handleWrite, data, ci)
		}

		commons.BytesSent += len(data)
		duration := time.Since(commons.Start)
		commons.Pkt++

		if duration > time.Second {
			rate := commons.BytesSent / int(duration.Seconds())
			remainingTime := commons.TsEnd.Sub(commons.TsStart) - duration
			log.Printf("\rrate %d kB/sec - sent %d/%d kB - %d/%d packets - remaining time %s", rate/1000, commons.BytesSent/1000, commons.Size/1000, commons.Pkt, commons.Packets, remainingTime)

			if commons.WithGui && !commons.ReplayFast {
				val := float64(commons.Pkt*100/commons.Packets) / 100

				err = commons.StatPBar.Set(val)
				if err != nil {
					log.Printf("Error setting progress bar: %v", err)
				}

				s := fmt.Sprintf("rate %d kB/sec - sent %d/%d kB - %d/%d packets - remaining time %s", rate/1000, commons.BytesSent/1000, commons.Size/1000, commons.Pkt, commons.Packets, remainingTime)

				err = commons.Stats2.Set(s)
				if err != nil {
					log.Printf("Error setting stats2: %v", err)
				}
			}
		}
	}
	return false
}

func Replay() {
	if commons.PcapHandle != nil {
		EndReplay()
	}
	LoadPcap(commons.PcapFile)
	handleWrite := OpenDest(commons.IntfId)

	for {
		if InternalReplay(handleWrite) {
			break
		}
	}
}

func ReplayStep(step int) {
	if commons.PcapHandle == nil {
		LoadPcap(commons.PcapFile)
	}
	handleWrite := OpenDest(commons.IntfId)

	for i := 0; i < step; i++ {
		if InternalReplay(handleWrite) {
			break
		}
	}
}
