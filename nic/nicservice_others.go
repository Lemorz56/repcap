//go:build !windows

package nic

import "net"

type NicServiceInterface interface {
	InitNics() error
	GetAllNics() ([]NetworkInterfaceCard, error)
	GetNicInfoByName(interfaceId string) (NetworkInterfaceCard, error)
}

// todo: make private struct
type nicService struct {
	networkInterfaceCards []NetworkInterfaceCard
}

type NetworkInterfaceCard struct {
	Id          string
	Description string
	Addresses   []net.Addr
}

func NewNicService() NicServiceInterface {
	return &nicService{}
}

func (w *nicService) InitNics() error {
	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			return err
		}
		w.networkInterfaceCards = append(w.networkInterfaceCards, NetworkInterfaceCard{
			Id:          i.Name,
			Description: i.Flags.String(),
			Addresses:   addrs,
		})
	}
	return nil
}

func (w *nicService) GetAllNics() ([]NetworkInterfaceCard, error) {
	return w.networkInterfaceCards, nil
}

func (w *nicService) GetNicInfoByName(interfaceId string) (NetworkInterfaceCard, error) {
	for _, nic := range w.networkInterfaceCards {
		if nic.Id == interfaceId {
			return nic, nil
		}
	}

	return NetworkInterfaceCard{}, nil
}
