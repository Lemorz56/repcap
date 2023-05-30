//go:build windows

package nic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestWindowsNicService_InitNics
func TestWindowsNicService_InitNics(t *testing.T) {
	//todo: create a mock for pcap.FindAllDevs()
	// Arrange
	var nicService = NewNicService()
	// Act
	err := nicService.InitNics()
	// Assert
	assert.Nil(t, err)
}

// TestWindowsNicService_GetAllNics
func TestWindowsNicService_GetAllNics(t *testing.T) {
	// Arrange
	var nicService = NewNicService()

	// Act
	err := nicService.InitNics()
	if err != nil {
		t.FailNow()
	}

	nics, err := nicService.GetAllNics()

	// Assert
	assert.Nil(t, err)
	assert.NotEmpty(t, nics)
}

// TestWindowsNicService_GetNicByName
func TestWindowsNicService_GetNicByName(t *testing.T) {
	// Arrange
	var nicService = NewNicService()

	// Act
	err := nicService.InitNics()
	if err != nil {
		t.FailNow()
	}

	//todo: improve test to not use GetAllNics() cheat
	nics, err := nicService.GetAllNics()
	if err != nil {
		t.FailNow()
	}
	nic, err := nicService.GetNicByName(nics[0].Id)

	// Assert
	assert.Nil(t, err)
	assert.NotEmpty(t, nic)
	assert.Equal(t, nics[0].Id, nic.Id)
}
