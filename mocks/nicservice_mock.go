package mocks

import (
	"github.com/google/gopacket/pcap"
	"github.com/lemorz56/pcapreplay/nic"
	"github.com/stretchr/testify/mock"
	"net"
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

type MockNicService struct {
	mock.Mock
	fakeNics []nic.NetworkInterfaceCard
}

func (m *MockNicService) InitNics() error {
	m.fakeNics = fakeNic
	return m.Called().Error(0)
}

func (m *MockNicService) GetAllNics() ([]nic.NetworkInterfaceCard, error) {
	args := m.Called()
	return args.Get(0).([]nic.NetworkInterfaceCard), args.Error(1)
}

func (m *MockNicService) GetNicByName(interfaceId string) (nic.NetworkInterfaceCard, error) {
	args := m.Called(interfaceId)
	return args.Get(0).(nic.NetworkInterfaceCard), args.Error(1)
}
