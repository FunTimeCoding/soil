package unit

import (
	"charm.land/bubbles/v2/table"
	"charm.land/bubbletea/v2"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/bubbletea/constant"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/tick"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/toast"
	"github.com/funtimecoding/soil/pkg/bubbletea/style"
	"github.com/funtimecoding/soil/pkg/bubbletea/table/item"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestStyle(t *testing.T) {
	style.Table(&table.Model{})
}

func TestConstructors(t *testing.T) {
	assert.NotNil(t, fetch.Command())
	assert.NotNil(t, monitor.New(false, memory.New()))
	assert.NotNil(t, tick.Command())
	assert.NotNil(t, toast.New(0, stringsConstant.UpperAlfa))
	assert.NotNil(t, item.New(true))
}

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
