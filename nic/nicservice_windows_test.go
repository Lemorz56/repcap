//go:build windows

package nic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://stackoverflow.com/questions/57708978/mock-functions-in-golang-for-testing
// https://stackoverflow.com/a/70080876
//todo: fix this code

// TestWindowsNicService_InitNics
func TestWindowsNicService_InitNics(t *testing.T) {
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

	nic, err := nicService.GetNicByName("Loopback") //todo: this will fail

	// Assert
	assert.Nil(t, err)
	assert.NotEmpty(t, nic)
}
