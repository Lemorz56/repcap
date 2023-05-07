package components

import (
	"testing"

	"fyne.io/fyne/v2/app"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndFillContainer(t *testing.T) {
	a := app.New()
	win := a.NewWindow("test")

	rp := NewReplayPane()
	rp.InitPane(win)
	rp.CreateAndFillContainer()

	assert.NotNil(t, rp.Container)
	assert.Contains(t, rp.Container.Objects, rp.Stats1, rp.Stats2)
}

func TestInitPane(t *testing.T) {
	a := app.New()
	win := a.NewWindow("test")

	rp := NewReplayPane()
	rp.InitPane(win)

	assert.NotNil(t, rp.Label)
	assert.NotNil(t, rp.Stats1)
	assert.NotNil(t, rp.Stats2)
	assert.NotNil(t, rp.FileField)
	assert.NotNil(t, rp.FileDialog)
	assert.NotNil(t, rp.FileButton)
}
