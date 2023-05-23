//go:build windows

package nic

import "github.com/google/gopacket/pcap"

type NicServiceInterface interface {
	InitNics() error
	GetAllNics() ([]NetworkInterfaceCard, error)
	GetNicByName(interfaceId string) (NetworkInterfaceCard, error)
}

// todo: make private struct
type nicService struct {
	networkInterfaceCards []NetworkInterfaceCard
}

type NetworkInterfaceCard struct {
	Id          string
	Description string
	Addresses   []pcap.InterfaceAddress
}

func NewNicService() NicServiceInterface {
	return &nicService{}
}

func (w *nicService) InitNics() error {
	interfaces, err := pcap.FindAllDevs()
	if err != nil {
		return err
	}

	for _, i := range interfaces {
		w.networkInterfaceCards = append(w.networkInterfaceCards, NetworkInterfaceCard{
			Id:          i.Name,
			Description: i.Description,
			Addresses:   i.Addresses,
		})
	}
	return nil
}

func (w *nicService) GetAllNics() ([]NetworkInterfaceCard, error) {
	return w.networkInterfaceCards, nil
}

func (w *nicService) GetNicByName(interfaceId string) (NetworkInterfaceCard, error) {
	for _, nic := range w.networkInterfaceCards {
		if nic.Id == interfaceId {
			return nic, nil
		}
	}

	return NetworkInterfaceCard{}, nil
}
