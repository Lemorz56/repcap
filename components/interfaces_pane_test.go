package components

import (
	"testing"

	"github.com/lemorz56/pcapreplay/mocks"
)

var fakeNic = mocks.FakeNic

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
