package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/bubbletea/constant"
	"testing"
)

func TestKeyConstant(t *testing.T) {
	assert.String(t, "ctrl+c", constant.KeyCtrlC)
}
