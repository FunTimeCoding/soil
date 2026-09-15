package unit

import (
	"charm.land/bubbletea/v2"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/bubbletea/constant"
	"testing"
)

func TestKeyConstantsMatchLibraryRendering(t *testing.T) {
	assert.String(t, "space", constant.KeySpace)
	assert.String(t, "space", tea.Key{Code: ' ', Text: " "}.String())
	assert.String(t, "enter", constant.KeyEnter)
	assert.String(t, "enter", tea.Key{Code: tea.KeyEnter}.String())
	assert.String(t, "esc", constant.KeyEscape)
	assert.String(t, "esc", tea.Key{Code: tea.KeyEscape}.String())
	assert.String(t, "q", constant.KeyQ)
	assert.String(t, "q", tea.Key{Code: 'q', Text: "q"}.String())
	assert.String(t, "ctrl+c", constant.KeyCtrlC)
	assert.String(t, "ctrl+c", tea.Key{Code: 'c', Mod: tea.ModCtrl}.String())
}
