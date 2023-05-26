package mocks

import (
	"net"

	"github.com/lemorz56/pcapreplay/nic"
	"github.com/stretchr/testify/mock"
)

var addr1 = net.IPAddr{
	IP:   net.IP{127, 0, 0, 1},
	Zone: "",
}

var FakeNic = []nic.NetworkInterfaceCard{
	{
		Id:          "fakeName",
		Description: "fakeDescription",
		Addresses: []string{
			addr1.String(),
		},
	},
}

type MockNicService struct {
	mock.Mock
	FakeNics []nic.NetworkInterfaceCard
}

func (m *MockNicService) InitNics() error {
	m.FakeNics = FakeNic
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
