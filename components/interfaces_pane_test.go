package components

import (
	"github.com/google/gopacket/pcap"
	"github.com/lemorz56/pcapreplay/mocks"
	"github.com/lemorz56/pcapreplay/nic"
	"net"
	"testing"
)

var fakeNic = []nic.NetworkInterfaceCard{
	{
		Id:          "fakeName",
		Description: "fakeDescription",
		Addresses: []pcap.InterfaceAddress{
			{
				IP:        net.IP{127, 0, 0, 1},
				Netmask:   net.IPMask{255, 0, 0, 0},
				Broadaddr: net.IP{127, 0, 0, 1},
				P2P:       net.IP{127, 0, 0, 1},
			},
		},
	},
}

func TestNewInterfacesPane(t *testing.T) {
	//todo
}

func TestInterfacesPane_InitPane(t *testing.T) {
	// Arrange

	mockNicService := new(mocks.MockNicService)
	mockNicService.On("InitNics").Return(nil)
	mockNicService.On("GetAllNics").Return(fakeNic, nil)
	//mockNicService.On("GetNicByName", "fakeName").Return(fakeNic[0], nil)
	ip := NewInterfacesPane(mockNicService)

	// Act
	ip.InitPane()

	// Assert
	mockNicService.AssertExpectations(t)
}
